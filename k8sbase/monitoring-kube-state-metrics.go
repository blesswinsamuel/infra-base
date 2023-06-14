package k8sbase

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
)

type KubeStateMetricsProps struct {
	Enabled       bool             `json:"enabled"`
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

func NewKubeStateMetrics(scope packager.Construct, props KubeStateMetricsProps) packager.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := packager.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("kube-state-metrics", cprops)

	k8sapp.NewHelm(chart, jsii.String("helm"), &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("kube-state-metrics"),
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
