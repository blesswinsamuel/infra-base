package k8sbase

import (
	_ "embed"

	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
	"github.com/rs/zerolog/log"
)

//go:embed values-default.yaml
var defaultValuesBytes []byte

func RegisterModules() {
	var defaultValues map[string]ast.Node
	if err := yaml.UnmarshalWithOptions(defaultValuesBytes, &defaultValues, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
		log.Panic().Err(err).Msg("Unmarshal default values")
	}
	newModules := map[string]k8sapp.ModuleWithMeta{
		"postgres": &k8sapp.ModuleCommons[*Postgres]{},
		"redis":    &k8sapp.ModuleCommons[*Redis]{},

		"traefik-forward-auth": &k8sapp.ModuleCommons[*TraefikForwardAuth]{},
		"lldap":                &k8sapp.ModuleCommons[*LLDAP]{},
		"authelia":             &k8sapp.ModuleCommons[*Authelia]{},

		"cert-issuer":  &k8sapp.ModuleCommons[*CertIssuer]{},
		"cert-manager": &k8sapp.ModuleCommons[*CertManagerProps]{},

		"traefik": &k8sapp.ModuleCommons[*TraefikProps]{},

		"alerting-rules":     &k8sapp.ModuleCommons[*AlertingRulesProps]{},
		"alertmanager":       &k8sapp.ModuleCommons[*AlertmanagerProps]{},
		"crowdsec":           &k8sapp.ModuleCommons[*Crowdsec]{},
		"grafana-dashboards": &k8sapp.ModuleCommons[*GrafanaDashboardsProps]{},
		"grafana":            &k8sapp.ModuleCommons[*Grafana]{},
		"kube-state-metrics": &k8sapp.ModuleCommons[*KubeStateMetricsProps]{},
		"loki":               &k8sapp.ModuleCommons[*LokiProps]{},
		"node-exporter":      &k8sapp.ModuleCommons[*NodeExporterProps]{},
		"vector":             &k8sapp.ModuleCommons[*VectorProps]{},
		"victoriametrics":    &k8sapp.ModuleCommons[*VictoriaMetrics]{},
		"vmagent":            &k8sapp.ModuleCommons[*VmagentProps]{},
		"vmalert":            &k8sapp.ModuleCommons[*VmalertProps]{},

		"external-secrets-store": &k8sapp.ModuleCommons[*ExternalSecretsStore]{},
		"external-secrets":       &k8sapp.ModuleCommons[*ExternalSecretsProps]{},

		"kube-gitops": &k8sapp.ModuleCommons[*KubeGitOpsProps]{},
		"reloader":    &k8sapp.ModuleCommons[*ReloaderProps]{},
		"pg-backuper": &k8sapp.ModuleCommons[*PgBackuper]{},

		"docker-creds": &k8sapp.ModuleCommons[*UtilsDockerCreds]{},
	}
	k8sapp.RegisterModules(newModules, defaultValues)
}
