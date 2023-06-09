package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type KubeStateMetricsProps struct {
	Enabled       bool             `yaml:"enabled"`
	HelmChartInfo k8sapp.ChartInfo `yaml:"helm"`
}

func NewKubeStateMetrics(scope constructs.Construct, props KubeStateMetricsProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: k8sapp.GetNamespaceContextPtr(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("kube-state-metrics"), &cprops)

	k8sapp.NewHelmCached(chart, jsii.String("helm"), &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("kube-state-metrics"),
		Namespace:   chart.Namespace(),
		Values: &map[string]interface{}{
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
