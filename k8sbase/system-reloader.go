package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type ReloaderProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/stakater/Reloader/blob/master/deployments/kubernetes/chart/reloader

func (props *ReloaderProps) Chart(scope kubegogen.Construct) kubegogen.Construct {
	chart := k8sapp.NewHelmChart(scope, "reloader", &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "reloader",
		Values: map[string]interface{}{
			"service": map[string]interface{}{
				"port": 9090,
				"annotations": map[string]string{
					"prometheus.io/port":   "9090",
					"prometheus.io/scrape": "true",
				},
			},
		},
	})

	return chart
}
