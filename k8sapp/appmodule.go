package k8sapp

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/packager"

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
		log.Fatalf("NodeToValue: %v", err)
	}
	valuesMerged := map[string]interface{}{}
	for _, valuesFile := range valuesFiles {
		fmt.Println("Loading values from", valuesFile)
		valuesFileBytes, err := os.ReadFile(valuesFile)
		if err != nil {
			log.Fatalf("ReadFile: %v", err)
		}
		if templateMap != nil {
			tpl := template.New("tpl")
			tpl.Delims("[{@", "@}]")
			out, err := tpl.Parse(string(valuesFileBytes))
			if err != nil {
				log.Fatalf("Parse: %v", err)
			}
			w := bytes.NewBuffer([]byte{})
			out.Execute(w, templateMap)
			valuesFileBytes = w.Bytes()
		}
		fileValues := map[string]interface{}{}
		if err := yaml.UnmarshalWithOptions(valuesFileBytes, &fileValues, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
		valuesMerged = mergeMaps(valuesMerged, fileValues)
	}
	valuesNode, err := yaml.ValueToNode(valuesMerged)
	if err != nil {
		log.Fatalf("ValueToNode: %v", err)
	}
	if err := yaml.NodeToValue(valuesNode, values, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

type Module interface {
	Chart(scope packager.Construct) packager.Construct
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

var registeredModules map[string]Module = map[string]Module{}

func RegisterModules(modules map[string]Module) {
	for k, v := range modules {
		RegisterModule(k, v)
	}
}

func RegisterModule(name string, module Module) {
	registeredModules[name] = module
}

func logModuleTiming(moduleName string, level int) func() {
	prefix := ""
	for i := 0; i < level; i++ {
		prefix += "  "
	}

	startTime := time.Now()
	log.Printf("%sStarting %q...", prefix, moduleName)
	return func() {
		log.Printf("%s └── Done %q in %s", prefix, moduleName, time.Since(startTime))
	}
}

func Render(scope packager.Construct, values ValuesProps) {
	for _, key := range values.Services.keyOrder {
		namespace, services := key, values.Services.Map[key]
		t := logModuleTiming(namespace, 0)
		namespaceChart := NewNamespaceChart(scope, namespace)
		for _, key := range services.keyOrder {
			serviceName, service := key, services.Map[key]
			moduleName := serviceName
			if service != nil {
				servicePart := struct {
					Module string `json:"_module"`
				}{}
				if err := yaml.NodeToValue(service, &servicePart); err != nil {
					log.Fatalf("NodeToValue: %v", err)
				}
				if servicePart.Module != "" {
					moduleName = servicePart.Module
				}
			}
			t := logModuleTiming(serviceName, 1)
			module := registeredModules[moduleName]
			if module == nil {
				log.Fatalf("module %q is not registered.", moduleName)
			}
			// fmt.Println(namespace, serviceName, service, reflect.TypeOf(module))
			if defaultValues, ok := DefaultValues[moduleName]; ok {
				if defaultValues == nil {
					log.Fatalf("defaultValues for %q is nil.", moduleName)
				}
				err := yaml.NodeToValue(defaultValues, module, yaml.Strict(), yaml.UseJSONUnmarshaler())
				if err != nil {
					log.Fatalf("NodeToValue: %v", err)
				}
			}
			if service != nil {
				moduleMap := map[string]interface{}{}
				if err := yaml.NodeToValue(service, &moduleMap, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
					log.Fatalf("NodeToValue(map): %v", err)
				}
				if _, ok := moduleMap["_module"]; ok {
					delete(moduleMap, "_module")
					service, err := yaml.ValueToNode(moduleMap)
					if err != nil {
						log.Fatalf("ValueToNode: %v", err)
					}
					if err := yaml.NodeToValue(service, module, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
						log.Fatalf("NodeToValue(Module): %v", err)
					}
					// TODO: fix this - sometimes, extra quotes are getting added
				} else {
					if err := yaml.NodeToValue(service, module, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
						log.Fatalf("NodeToValue(Module): %v", err)
					}
				}
			}
			// unmarshal(module, service)
			namespaceChart.SetContext("name", serviceName)
			module.Chart(namespaceChart)
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
// 			log.Fatalf("NodeToValue: %v", err)
// 		}
// 	}
// }
