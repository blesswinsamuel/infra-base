package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type PostgresProps struct {
	Enabled       bool      `yaml:"enabled"`
	HelmChartInfo ChartInfo `yaml:"helm"`
	ImageInfo     ImageInfo `yaml:"image"`
	Database      string    `yaml:"database"`
	Username      string    `yaml:"username"`
}

func NewPostgres(scope constructs.Construct, props PostgresProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: GetNamespace(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("postgres"), &cprops)

	NewHelmCached(chart, jsii.String("helm"), &HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("postgres"),
		Namespace:   chart.Namespace(),
		Values: &map[string]interface{}{
			"nameOverride": "postgres",
			"image":        props.ImageInfo.ToMap(),
			"auth": map[string]interface{}{
				"database":       props.Database,
				"username":       props.Username,
				"existingSecret": "postgres-passwords",
			},
			"metrics": map[string]interface{}{"enabled": true},
		},
	})

	NewExternalSecret(chart, jsii.String("external-secret"), &ExternalSecretProps{
		Name:            jsii.String("postgres-passwords"),
		RefreshInterval: jsii.String("2m"),
		Secrets: map[string]string{
			"postgres-password":    "POSTGRES_ADMIN_PASSWORD",
			"password":             "POSTGRES_USER_PASSWORD",
			"replication-password": "POSTGRES_REPLICATION_PASSWORD",
		},
	})

	return chart
}
