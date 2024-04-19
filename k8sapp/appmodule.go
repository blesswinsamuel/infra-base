package k8sapp

import (
	"bytes"
	"context"
	"html/template"
	"os"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
	"github.com/rs/zerolog/log"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"

	_ "embed"
)

//go:embed values-default.yaml
var defaultValues []byte

var DefaultValues map[string]ast.Node

func init() {
	if err := yaml.UnmarshalWithOptions(defaultValues, &DefaultValues, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
		panic(err)
	}
}

func mergeMaps(a, b map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(a))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		if v, ok := v.(map[string]interface{}); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[string]interface{}); ok {
					out[k] = mergeMaps(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
}

func LoadValues(values *ValuesProps, valuesFiles []string, templateMap map[string]any) {
	err := yaml.NodeToValue(DefaultValues["global"], &values.Global, yaml.Strict(), yaml.UseJSONUnmarshaler())
	if err != nil {
		log.Panic().Err(err).Msg("NodeToValue")
	}
	valuesMerged := map[string]interface{}{}
	for _, valuesFile := range valuesFiles {
		log.Info().Msgf("Loading values from %s", valuesFile)
		valuesFileBytes, err := os.ReadFile(valuesFile)
		if err != nil {
			log.Panic().Err(err).Msg("ReadFile")
		}
		if templateMap != nil {
			tpl := template.New("tpl")
			tpl.Delims("[{@", "@}]")
			out, err := tpl.Parse(string(valuesFileBytes))
			if err != nil {
				log.Panic().Err(err).Msg("Parse")
			}
			w := bytes.NewBuffer([]byte{})
			out.Execute(w, templateMap)
			valuesFileBytes = w.Bytes()
		}
		fileValues := map[string]interface{}{}
		if err := yaml.UnmarshalWithOptions(valuesFileBytes, &fileValues, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
			log.Panic().Err(err).Msg("Unmarshal")
		}
		valuesMerged = mergeMaps(valuesMerged, fileValues)
	}
	valuesNode, err := yaml.ValueToNode(valuesMerged)
	if err != nil {
		log.Panic().Err(err).Msg("ValueToNode")
	}
	if err := yaml.NodeToValue(valuesNode, values, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
		log.Panic().Err(err).Msg("Unmarshal")
	}
}

type Module interface {
	Chart(scope kubegogen.Construct) kubegogen.Construct
}

type ModuleWithMeta interface {
	Module
	GetModuleName() string
	GetDashboards() map[string]GrafanaDashboardProps
}

type OrderedMap[K comparable, V any] struct {
	keyOrder []K
	Map      map[K]V
}

func (m *OrderedMap[K, V]) UnmarshalYAML(ctx context.Context, data []byte) error {
	newMap := infrahelpers.MergeableMap[K, V](m.Map)
	if err := yaml.UnmarshalContext(ctx, data, &newMap); err != nil {
		return err
	}
	var orderedMapItems yaml.MapSlice
	if err := yaml.UnmarshalContext(ctx, data, &orderedMapItems); err != nil {
		return err
	}
	if m.Map == nil {
		m.Map = make(map[K]V)
	}
	// fmt.Println(m.keyOrder, m.Map)
	for _, item := range orderedMapItems {
		k := item.Key.(K)
		if _, ok := m.Map[k]; !ok {
			m.keyOrder = append(m.keyOrder, k)
			m.Map[k] = newMap[k]
		}
		// else {
		// 	var v ast.Node
		// 	// m.Map[k] = infrahelpers.Merge(m.Map[k], newMap[k])
		// }
	}
	// fmt.Println(m.keyOrder, m.Map)
	return nil
}

type GlobalProps struct {
	Domain                         string `json:"domain"`
	CertIssuer                     string `json:"clusterCertIssuerName"`
	ClusterExternalSecretStoreName string `json:"clusterExternalSecretStoreName"`
}

type ValuesProps struct {
	Global   GlobalProps                                      `json:"global"`
	Services OrderedMap[string, OrderedMap[string, ast.Node]] `json:"services"`
}

type ModuleCommons[T Module] struct {
	Module     string                                                   `json:"_module"`
	Dashboards infrahelpers.MergeableMap[string, GrafanaDashboardProps] `json:"_dashboards"`

	Rest T `json:",inline"`
}

func (m ModuleCommons[T]) Chart(scope kubegogen.Construct) kubegogen.Construct {
	return m.Rest.Chart(scope)
}

func (m ModuleCommons[T]) GetModuleName() string {
	return m.Module
}

func (m ModuleCommons[T]) GetDashboards() map[string]GrafanaDashboardProps {
	return m.Dashboards
}

var registeredModules map[string]ModuleWithMeta = map[string]ModuleWithMeta{}

func RegisterModules(modules map[string]ModuleWithMeta) {
	for k, v := range modules {
		RegisterModule(k, v)
	}
}

func RegisterModule(name string, module ModuleWithMeta) {
	registeredModules[name] = module
}

func logModuleTiming(moduleName string, level int) func() {
	prefix := ""
	for i := 0; i < level; i++ {
		prefix += "  "
	}

	startTime := time.Now()
	log.Debug().Msgf("%sStarting %q...", prefix, moduleName)
	return func() {
		log.Debug().Msgf("%s └── Done %q in %s", prefix, moduleName, time.Since(startTime))
	}
}

func Render(scope kubegogen.Construct, values ValuesProps) {
	for _, key := range values.Services.keyOrder {
		namespace, services := key, values.Services.Map[key]
		t := logModuleTiming(namespace, 0)
		namespaceChart := NewNamespaceChart(scope, namespace)
		for _, key := range services.keyOrder {
			serviceName, serviceProps := key, services.Map[key]
			moduleName := serviceName
			if serviceProps != nil {
				// module name is different from service name
				serviceCommons := struct {
					Module string `json:"_module"`
				}{}
				if err := yaml.NodeToValue(serviceProps, &serviceCommons, yaml.UseJSONUnmarshaler()); err != nil {
					log.Panic().Err(err).Msg("NodeToValue (module)")
				}
				if serviceCommons.Module != "" {
					moduleName = serviceCommons.Module
				}
			}
			t := logModuleTiming(serviceName, 1)
			module := registeredModules[moduleName]
			if module == nil {
				log.Panic().Msgf("module %q is not registered.", moduleName)
			}
			// fmt.Println(namespace, serviceName, service, reflect.TypeOf(module))
			if defaultValues, ok := DefaultValues[moduleName]; ok {
				if defaultValues == nil {
					log.Panic().Msgf("defaultValues for %q is nil.", moduleName)
				}
				if err := yaml.NodeToValue(defaultValues, module, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
					log.Panic().Err(err).Msg("NodeToValue(defaults)")
				}
			}
			if serviceProps != nil {
				if err := yaml.NodeToValue(serviceProps, module, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
					log.Panic().Err(err).Msg("NodeToValue(Module)")
				}
			}
			// unmarshal(module, service)
			namespaceChart.SetContext("name", serviceName)
			chart := module.Chart(namespaceChart)
			NewGrafanaDashboards(chart, module.GetDashboards())
			t()
		}
		t()
	}
}

// func unmarshal[T Module](module T, val ast.Node) {
// 	moduleWithMeta := &struct {
// 		ModuleName string `json:"_module"`
// 		Module     T      `json:",inline"`
// 	}{Module: module}
// 	if val != nil {
// 		err := yaml.NodeToValue(val, moduleWithMeta, yaml.Strict(), yaml.UseJSONUnmarshaler())
// 		if err != nil {
// 			log.Panic().Err(err).Msg("NodeToValue")
// 		}
// 	}
// }
