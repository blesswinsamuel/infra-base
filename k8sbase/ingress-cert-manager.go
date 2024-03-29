package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type CertManagerProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/cert-manager/cert-manager/tree/master/deploy/charts/cert-manager
// https://artifacthub.io/packages/helm/cert-manager/cert-manager
func (props *CertManagerProps) Chart(scope kubegogen.Construct) kubegogen.Construct {
	chart := k8sapp.NewHelmChart(scope, "helm", &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "cert-manager",
		Values: map[string]interface{}{
			"installCRDs": "true",
		},
	})

	return chart
}
