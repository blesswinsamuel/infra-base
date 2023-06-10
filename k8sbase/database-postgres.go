package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type PostgresProps struct {
	Enabled       bool             `json:"enabled"`
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
	ImageInfo     k8sapp.ImageInfo `json:"image"`
	Database      string           `json:"database"`
	Username      string           `json:"username"`
}

func NewPostgres(scope constructs.Construct, props PostgresProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: k8sapp.GetNamespaceContextPtr(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("postgres"), &cprops)

	k8sapp.NewHelm(chart, jsii.String("helm"), &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("postgres"),
		Namespace:   chart.Namespace(),
		Values: map[string]interface{}{
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

	k8sapp.NewExternalSecret(chart, jsii.String("external-secret"), &k8sapp.ExternalSecretProps{
		Name: "postgres-passwords",
		RemoteRefs: map[string]string{
			"postgres-password":    "POSTGRES_ADMIN_PASSWORD",
			"password":             "POSTGRES_USER_PASSWORD",
			"replication-password": "POSTGRES_REPLICATION_PASSWORD",
		},
	})

	return chart
}
