package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type MariaDBProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
	Database      string           `json:"database"`
	Username      string           `json:"username"`
}

func (props *MariaDBProps) Chart(scope kubegogen.Construct) kubegogen.Construct {
	cprops := kubegogen.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("mariadb", cprops)

	k8sapp.NewHelm(chart, "helm", &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "mariadb",
		Namespace:   chart.Namespace(),
		Values: map[string]interface{}{
			"nameOverride": "mariadb",
			"auth": map[string]interface{}{
				"database":       props.Database,
				"username":       props.Username,
				"existingSecret": "mariadb-passwords",
			},
			"metrics": map[string]interface{}{"enabled": true},
		},
	})

	k8sapp.NewExternalSecret(chart, "external-secret", &k8sapp.ExternalSecretProps{
		Name: "mariadb-passwords",
		RemoteRefs: map[string]string{
			"mariadb-password":             "MARIADB_PASSWORD",
			"mariadb-root-password":        "MARIADB_ROOT_PASSWORD",
			"mariadb-replication-password": "MARIADB_REPLICATION_PASSWORD",
		},
	})

	return chart
}
