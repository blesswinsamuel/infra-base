package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type KubeStateMetricsProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

func (props *KubeStateMetricsProps) Render(scope kubegogen.Scope) {
	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "kube-state-metrics",
		Namespace:   scope.Namespace(),
		Values: map[string]interface{}{
			"fullnameOverride": "kube-state-metrics",
			"service": map[string]interface{}{
				"annotations": map[string]string{
					"prometheus.io/port": "8080",
				},
			},
		},
	})
}
