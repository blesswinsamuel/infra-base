package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/k8sbase/helpers"
)

type CertManagerProps struct {
	Enabled       bool              `yaml:"enabled"`
	HelmChartInfo helpers.ChartInfo `yaml:"helm"`
}

// https://github.com/cert-manager/cert-manager/tree/master/deploy/charts/cert-manager
// https://artifacthub.io/packages/helm/cert-manager/cert-manager
func NewCertManager(scope constructs.Construct, props CertManagerProps) constructs.Construct {
	if !props.Enabled {
		return nil
	}
	chart := k8sapp.NewNamespaceChart(scope, "cert-manager")

	helpers.NewHelmCached(chart, jsii.String("helm"), &helpers.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("cert-manager"),
		Namespace:   chart.Namespace(),
		Values: &map[string]interface{}{
			"installCRDs": "true",
		},
	})

	return chart
}
