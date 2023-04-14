package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type MariaDBProps struct {
	Enabled       bool      `yaml:"enabled"`
	HelmChartInfo ChartInfo `yaml:"helm"`
	Database      string    `yaml:"database"`
	Username      string    `yaml:"username"`
}

func NewMariaDB(scope constructs.Construct, props MariaDBProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: GetNamespace(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("mariadb"), &cprops)

	NewHelmCached(chart, jsii.String("helm"), &HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("mariadb"),
		Namespace:   chart.Namespace(),
		Values: &map[string]interface{}{
			"nameOverride": "mariadb",
			"auth": map[string]interface{}{
				"database":       props.Database,
				"username":       props.Username,
				"existingSecret": "mariadb-passwords",
			},
			"metrics": map[string]interface{}{"enabled": true},
		},
	})

	NewExternalSecret(chart, jsii.String("external-secret"), &ExternalSecretProps{
		Name:            jsii.String("mariadb-passwords"),
		RefreshInterval: jsii.String("2m"),
		Secrets: map[string]string{
			"mariadb-password":             "MARIADB_PASSWORD",
			"mariadb-root-password":        "MARIADB_ROOT_PASSWORD",
			"mariadb-replication-password": "MARIADB_REPLICATION_PASSWORD",
		},
	})

	return chart
}