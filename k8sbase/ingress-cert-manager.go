package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type CertManagerProps struct {
	Enabled       bool             `json:"enabled"`
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/cert-manager/cert-manager/tree/master/deploy/charts/cert-manager
// https://artifacthub.io/packages/helm/cert-manager/cert-manager
func NewCertManager(scope constructs.Construct, props CertManagerProps) constructs.Construct {
	if !props.Enabled {
		return nil
	}

	chart := k8sapp.NewHelmChart(scope, jsii.String("helm"), &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("cert-manager"),
		Values: map[string]interface{}{
			"installCRDs": "true",
		},
	})

	return chart
}
