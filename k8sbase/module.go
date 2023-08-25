package k8sbase

import (
	"github.com/goccy/go-yaml/ast"

	"github.com/blesswinsamuel/infra-base/packager"
)

type Module interface {
	Chart(scope packager.Construct) packager.Construct
}

// type RawMessage struct {
// 	unmarshal func(interface{}) error
// }

// func (msg *RawMessage) UnmarshalYAML(unmarshal func(interface{}) error) error {
// 	msg.unmarshal = unmarshal
// 	return nil
// }

// func (msg *RawMessage) Unmarshal(v interface{}) error {
// 	return msg.unmarshal(v)
// }

type ValuesProps struct {
	Global   GlobalProps                    `json:"global"`
	Services map[string]map[string]ast.Node `json:"services"`
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
		"docker-creds":         &SecretsDockerCredsProps{},
		"external-secrets":     &ExternalSecretsProps{},

		"backup-job":           &BackupJobProps{},
		"kopia":                &KopiaProps{},
		"kube-gitops":          &KubeGitOpsProps{},
		"kubernetes-dashboard": &KubernetesDashboardProps{},
		"reloader":             &ReloaderProps{},
	}
	RegisterModules(newModules)
}

func GetRegisteredModule(name string) Module {
	return registeredModules[name]
}

// func (props *DatabaseProps) Chart(scope packager.Construct) packager.Construct {
// 	if !props.Enabled {
// 		return nil
// 	}
// 	defer logModuleTiming("database")()

// 	chart := k8sapp.NewNamespaceChart(scope, "database")

// 	NewMariaDB(chart, props.MariaDB)
// 	NewPostgres(chart, props.Postgres)
// 	NewRedis(chart, props.Redis)

// 	return chart
// }
