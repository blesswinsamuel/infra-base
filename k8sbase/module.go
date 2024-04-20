package k8sbase

import "github.com/blesswinsamuel/infra-base/k8sapp"

func RegisterModules() {
	newModules := map[string]k8sapp.ModuleWithMeta{
		"postgres": &k8sapp.ModuleCommons[*Postgres]{},
		"mariadb":  &k8sapp.ModuleCommons[*MariaDB]{},
		"redis":    &k8sapp.ModuleCommons[*Redis]{},

		"traefik-forward-auth": &k8sapp.ModuleCommons[*TraefikForwardAuth]{},
		"lldap":                &k8sapp.ModuleCommons[*LLDAP]{},
		"authelia":             &k8sapp.ModuleCommons[*Authelia]{},

		"cert-issuer":  &k8sapp.ModuleCommons[*CertIssuerProps]{},
		"cert-manager": &k8sapp.ModuleCommons[*CertManagerProps]{},

		"traefik": &k8sapp.ModuleCommons[*TraefikProps]{},

		"alerting-rules":           &k8sapp.ModuleCommons[*AlertingRulesProps]{},
		"alertmanager":             &k8sapp.ModuleCommons[*AlertmanagerProps]{},
		"crowdsec-traefik-bouncer": &k8sapp.ModuleCommons[*CrowdsecTraefikBouncerProps]{},
		"crowdsec":                 &k8sapp.ModuleCommons[*CrowdsecProps]{},
		"grafana-dashboards":       &k8sapp.ModuleCommons[*GrafanaDashboardsProps]{},
		"grafana":                  &k8sapp.ModuleCommons[*Grafana]{},
		"kube-state-metrics":       &k8sapp.ModuleCommons[*KubeStateMetricsProps]{},
		"loki":                     &k8sapp.ModuleCommons[*LokiProps]{},
		"node-exporter":            &k8sapp.ModuleCommons[*NodeExporterProps]{},
		"vector":                   &k8sapp.ModuleCommons[*VectorProps]{},
		"victoria-metrics":         &k8sapp.ModuleCommons[*VictoriaMetricsProps]{},
		"vmagent":                  &k8sapp.ModuleCommons[*VmagentProps]{},
		"vmalert":                  &k8sapp.ModuleCommons[*VmalertProps]{},

		"cluster-secret-store": &k8sapp.ModuleCommons[*ClusterSecretStoreProps]{},
		"external-secrets":     &k8sapp.ModuleCommons[*ExternalSecretsProps]{},

		"backup-job":           &k8sapp.ModuleCommons[*BackupJobProps]{},
		"kopia":                &k8sapp.ModuleCommons[*KopiaProps]{},
		"kube-gitops":          &k8sapp.ModuleCommons[*KubeGitOpsProps]{},
		"kubernetes-dashboard": &k8sapp.ModuleCommons[*KubernetesDashboardProps]{},
		"reloader":             &k8sapp.ModuleCommons[*ReloaderProps]{},
		"pg-backuper":          &k8sapp.ModuleCommons[*PgBackuper]{},

		"docker-creds": &k8sapp.ModuleCommons[*UtilsDockerCreds]{},
	}
	k8sapp.RegisterModules(newModules)
}
