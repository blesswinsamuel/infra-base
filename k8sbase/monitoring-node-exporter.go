package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type NodeExporterProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
	Disable       bool             `json:"disable"`
}

// https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus-node-exporter/values.yaml
func (props *NodeExporterProps) Render(scope kubegogen.Scope) {
	if props.Disable {
		return
	}
	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "node-exporter",
		Namespace:   scope.Namespace(),
		Values: map[string]interface{}{
			"fullnameOverride": "node-exporter",
			"service": map[string]interface{}{
				"annotations": map[string]string{
					"prometheus.io/scrape": "true",
					"prometheus.io/port":   "9100",
				},
			},
		},
	})
}
