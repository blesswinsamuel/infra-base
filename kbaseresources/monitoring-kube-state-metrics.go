package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
)

type KubeStateMetricsProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

func (props *KubeStateMetricsProps) Render(scope kgen.Scope) {
	// TODO: remove helm dependency
	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "kube-state-metrics",
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
