package k8sbase

import (
	"bytes"
	"context"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"

	"github.com/blesswinsamuel/infra-base/k8sapp"
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

func LoadValues[T any](values *T, valuesFiles []string, templateMap map[string]any) {
	for _, valuesFile := range valuesFiles {
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
		// TODO: maybe go back to gopkg.in/yaml.v3 library
		// "github.com/goccy/go-yaml" library overwrites the map (example: grafana dashboards)
		// mergo approach doesn't override boolean values because it considers them as zero values
		// validation
		// var values T
		customUnmarshaller := yaml.CustomUnmarshaler(func(dst *map[string]GrafanaDashboardsConfigProps, b []byte) error {
			// workaround to allow merging of maps
			v := map[string]GrafanaDashboardsConfigProps{}
			if err := yaml.UnmarshalWithOptions(b, &v, yaml.Strict()); err != nil {
				return err
			}
			for k, v := range v {
				(*dst)[k] = v
			}
			return nil
		})
		if err := yaml.UnmarshalWithOptions(valuesFileBytes, values, yaml.Strict(), yaml.UseJSONUnmarshaler(), customUnmarshaller); err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}
}

type Module interface {
	Chart(scope packager.Construct) packager.Construct
}

type OrderedMapItem[K comparable, V any] struct {
	Key   K
	Value V
}

type OrderedMap[K comparable, V any] []OrderedMapItem[K, V]

func (m *OrderedMap[K, V]) UnmarshalYAML(ctx context.Context, data []byte) error {
	var goMap map[K]V
	if err := yaml.UnmarshalContext(ctx, data, &goMap); err != nil {
		return err
	}
	var orderedMapItems yaml.MapSlice
	if err := yaml.UnmarshalContext(ctx, data, &orderedMapItems); err != nil {
		return err
	}
	for _, item := range orderedMapItems {
		k := item.Key.(K)
		v := goMap[k]
		*m = append(*m, OrderedMapItem[K, V]{k, v})
	}
	return nil
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

func init() {
	newModules := map[string]Module{
		"postgres": &PostgresProps{},
		"mariadb":  &MariaDBProps{},
		"redis":    &RedisProps{},

		"traefik-forward-auth": &TraefikForwardAuthProps{},
		"lldap":                &LLDAPProps{},
		"authelia":             &AutheliaProps{},

		"cert-issuer":  &CertIssuerProps{},
		"cert-manager": &CertManagerProps{},

		"traefik": &TraefikProps{},

		"alerting-rules":           &AlertingRulesProps{},
		"alertmanager":             &AlertmanagerProps{},
		"crowdsec-traefik-bouncer": &CrowdsecTraefikBouncerProps{},
		"crowdsec":                 &CrowdsecProps{},
		"grafana-dashboards":       &GrafanaDashboardsProps{},
		"grafana-datasource":       &GrafanaDatasourceProps{},
		"grafana":                  &GrafanaProps{},
		"kube-state-metrics":       &KubeStateMetricsProps{},
		"loki":                     &LokiProps{},
		"node-exporter":            &NodeExporterProps{},
		"vector":                   &VectorProps{},
		"victoria-metrics":         &VictoriaMetricsProps{},
		"vmagent":                  &VmagentProps{},
		"vmalert":                  &VmalertProps{},

		"cluster-secret-store": &ClusterSecretStoreProps{},
		"external-secrets":     &ExternalSecretsProps{},

		"backup-job":           &BackupJobProps{},
		"kopia":                &KopiaProps{},
		"kube-gitops":          &KubeGitOpsProps{},
		"kubernetes-dashboard": &KubernetesDashboardProps{},
		"reloader":             &ReloaderProps{},

		"docker-creds": &UtilsDockerCreds{},
	}
	RegisterModules(newModules)
}

func GetRegisteredModule(name string) Module {
	return registeredModules[name]
}

func logModuleTiming(moduleName string, level int) func() {
	prefix := ""
	for i := 0; i < level; i++ {
		prefix += "  "
	}

	startTime := time.Now()
	log.Printf("%sStarting %q..", prefix, moduleName)
	return func() {
		log.Printf("%s └── Done %q in %s", prefix, moduleName, time.Since(startTime))
	}
}

func Render(scope packager.Construct, values ValuesProps) {
	customUnmarshaller := yaml.CustomUnmarshaler(func(dst *map[string]GrafanaDashboardsConfigProps, b []byte) error {
		// workaround to allow merging of maps
		v := map[string]GrafanaDashboardsConfigProps{}
		if err := yaml.UnmarshalWithOptions(b, &v, yaml.Strict()); err != nil {
			return err
		}
		if *dst == nil {
			*dst = map[string]GrafanaDashboardsConfigProps{}
		}
		for k, v := range v {
			(*dst)[k] = v
		}
		return nil
	})
	for _, v := range values.Services {
		namespace, services := v.Key, v.Value
		t := logModuleTiming(namespace, 0)
		namespaceChart := k8sapp.NewNamespaceChart(scope, namespace)
		for _, v := range services {
			serviceName, service := v.Key, v.Value
			t := logModuleTiming(serviceName, 1)
			// k8sbase.NewService(app, k8sbase.ServiceProps{
			// 	Namespace:    namespace,
			// 	ServiceName:  serviceName,
			// 	ServiceProps: service,
			// })
			module := GetRegisteredModule(serviceName)
			if module == nil {
				log.Fatalf("module %q is not registered.", serviceName)
			}
			// fmt.Println(namespace, serviceName, service, reflect.TypeOf(module))
			if defaultValues, ok := DefaultValues[serviceName]; ok {
				if defaultValues == nil {
					log.Fatalf("defaultValues for %q is nil.", serviceName)
				}
				err := yaml.NodeToValue(defaultValues, module, yaml.Strict(), yaml.UseJSONUnmarshaler(), customUnmarshaller)
				if err != nil {
					log.Fatalf("NodeToValue: %v", err)
				}
			}
			if service != nil {
				err := yaml.NodeToValue(service, module, yaml.Strict(), yaml.UseJSONUnmarshaler(), customUnmarshaller)
				if err != nil {
					log.Fatalf("NodeToValue: %v", err)
				}
			}
			module.Chart(namespaceChart)
			t()
		}
		t()
	}
}
