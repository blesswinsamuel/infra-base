package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type ReloaderProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/stakater/Reloader/blob/master/deployments/kubernetes/chart/reloader

func (props *ReloaderProps) Render(scope kubegogen.Scope) {
	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
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
}
