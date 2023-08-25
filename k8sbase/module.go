package k8sbase

import "github.com/blesswinsamuel/infra-base/k8sapp"

func RegisterModules() {
	newModules := map[string]k8sapp.Module{
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
	k8sapp.RegisterModules(newModules)
}
