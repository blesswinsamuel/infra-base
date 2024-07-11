package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
)

func RegisterModules() {
	newModules := map[string]k8sapp.ModuleWithMeta{
		"postgres": &k8sapp.ModuleCommons[*Postgres]{},
		"redis":    &k8sapp.ModuleCommons[*Redis]{},

		"traefik-forward-auth": &k8sapp.ModuleCommons[*TraefikForwardAuth]{},
		"lldap":                &k8sapp.ModuleCommons[*LLDAP]{},
		"authelia":             &k8sapp.ModuleCommons[*Authelia]{},

		"cert-issuer":  &k8sapp.ModuleCommons[*CertIssuer]{},
		"cert-manager": &k8sapp.ModuleCommons[*CertManagerProps]{},

		"traefik":               &k8sapp.ModuleCommons[*TraefikProps]{},
		"default-wildcard-cert": &k8sapp.ModuleCommons[*DefaultWildcardCertificateProps]{},

		"alerting-rules":            &k8sapp.ModuleCommons[*AlertingRulesProps]{},
		"alertmanager":              &k8sapp.ModuleCommons[*AlertmanagerProps]{},
		"crowdsec":                  &k8sapp.ModuleCommons[*Crowdsec]{},
		"crowdsec-firewall-bouncer": &k8sapp.ModuleCommons[*CrowdsecFirewallBouncer]{},
		"grafana-dashboards":        &k8sapp.ModuleCommons[*GrafanaDashboardsProps]{},
		"grafana":                   &k8sapp.ModuleCommons[*Grafana]{},
		"loki":                      &k8sapp.ModuleCommons[*LokiProps]{},
		"node-exporter":             &k8sapp.ModuleCommons[*NodeExporterProps]{},
		"vector":                    &k8sapp.ModuleCommons[*VectorProps]{},
		"victoriametrics":           &k8sapp.ModuleCommons[*VictoriaMetrics]{},
		"vmagent":                   &k8sapp.ModuleCommons[*VmagentProps]{},
		"vmalert":                   &k8sapp.ModuleCommons[*VmalertProps]{},

		"external-secrets-store": &k8sapp.ModuleCommons[*ExternalSecretsStore]{},
		"external-secrets":       &k8sapp.ModuleCommons[*ExternalSecretsProps]{},

		"kube-gitops": &k8sapp.ModuleCommons[*KubeGitOpsProps]{},
		"pg-backuper": &k8sapp.ModuleCommons[*PgBackuper]{},

		"docker-creds": &k8sapp.ModuleCommons[*UtilsDockerCreds]{},
	}
	k8sapp.RegisterModules(newModules, defaultValues)
}

func RegisterModule[T k8sapp.Module](name string, module T) {
	k8sapp.RegisterModule(name, &k8sapp.ModuleCommons[T]{}, defaultValues[name])
}
