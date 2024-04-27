package k8sapp

import (
	"context"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
	"github.com/rs/zerolog/log"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"

	_ "embed"
)

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

type Module interface {
	Render(scope kubegogen.Scope)
}

type ModuleWithMeta interface {
	Module
	GetModuleName() string
	GetGrafanaDashboards() map[string]GrafanaDashboard
	GetAlertingRules() map[string]AlertingRules
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

type ModuleCommons[T Module] struct {
	Module            string                                              `json:"_module"`
	GrafanaDashboards infrahelpers.MergeableMap[string, GrafanaDashboard] `json:"_grafana_dashboards"`
	AlertingRules     infrahelpers.MergeableMap[string, AlertingRules]    `json:"_alerting_rules"`

	Rest T `json:",inline"`
}

func (m ModuleCommons[T]) Render(scope kubegogen.Scope) {
	m.Rest.Render(scope)
}

func (m ModuleCommons[T]) GetModuleName() string {
	return m.Module
}

func (m ModuleCommons[T]) GetGrafanaDashboards() map[string]GrafanaDashboard {
	return m.GrafanaDashboards
}

func (m ModuleCommons[T]) GetAlertingRules() map[string]AlertingRules {
	return m.AlertingRules
}

var registeredModules map[string]ModuleWithMeta = map[string]ModuleWithMeta{}
var registeredDefaultValues map[string]ast.Node = map[string]ast.Node{}

func RegisterModules(modules map[string]ModuleWithMeta, defaultValues map[string]ast.Node) {
	for k, v := range modules {
		RegisterModule(k, v, defaultValues[k])
	}
}

func RegisterModule(name string, module ModuleWithMeta, defaultValues ast.Node) {
	registeredModules[name] = module
	if defaultValues != nil {
		registeredDefaultValues[name] = defaultValues
	}
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
