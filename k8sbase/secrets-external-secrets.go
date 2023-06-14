package k8sbase

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
)

type ExternalSecretsProps struct {
	Enabled       bool             `json:"enabled"`
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/external-secrets/external-secrets/tree/main/deploy/charts/external-secrets
func NewExternalSecrets(scope packager.Construct, props ExternalSecretsProps) packager.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := packager.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("external-secrets", cprops)

	k8sapp.NewHelm(chart, jsii.String("helm"), &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("external-secrets"),
		Namespace:   chart.Namespace(),
		Values: map[string]interface{}{
			"installCRDs": "true",
		},
	})

	return chart
}
