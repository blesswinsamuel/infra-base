package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type KubeStateMetricsProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

func (props *KubeStateMetricsProps) Chart(scope kubegogen.Scope) kubegogen.Scope {
	cprops := kubegogen.ScopeProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.CreateScope("kube-state-metrics", cprops)

	k8sapp.NewHelm(chart, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "kube-state-metrics",
		Namespace:   chart.Namespace(),
		Values: map[string]interface{}{
			"fullnameOverride": "kube-state-metrics",
			"service": map[string]interface{}{
				"annotations": map[string]string{
					"prometheus.io/port": "8080",
				},
			},
		},
	})

	return chart
}
