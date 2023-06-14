package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
)

type CertManagerProps struct {
	Enabled       bool             `json:"enabled"`
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/cert-manager/cert-manager/tree/master/deploy/charts/cert-manager
// https://artifacthub.io/packages/helm/cert-manager/cert-manager
func NewCertManager(scope packager.Construct, props CertManagerProps) packager.Construct {
	if !props.Enabled {
		return nil
	}

	chart := k8sapp.NewHelmChart(scope, "helm", &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "cert-manager",
		Values: map[string]interface{}{
			"installCRDs": "true",
		},
	})

	return chart
}
