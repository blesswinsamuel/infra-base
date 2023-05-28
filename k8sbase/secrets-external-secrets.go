package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/helpers"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type ExternalSecretsProps struct {
	Enabled       bool              `yaml:"enabled"`
	HelmChartInfo helpers.ChartInfo `yaml:"helm"`
}

// https://github.com/external-secrets/external-secrets/tree/main/deploy/charts/external-secrets
func NewExternalSecrets(scope constructs.Construct, props ExternalSecretsProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: helpers.GetNamespace(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("external-secrets"), &cprops)

	helpers.NewHelmCached(chart, jsii.String("helm"), &helpers.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("external-secrets"),
		Namespace:   chart.Namespace(),
		Values: &map[string]interface{}{
			"installCRDs": "true",
		},
	})

	return chart
}
