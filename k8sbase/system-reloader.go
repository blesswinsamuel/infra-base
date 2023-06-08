package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/k8sbase/helpers"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type ReloaderProps struct {
	Enabled       bool              `yaml:"enabled"`
	HelmChartInfo helpers.ChartInfo `yaml:"helm"`
}

// https://github.com/stakater/Reloader/blob/master/deployments/kubernetes/chart/reloader

func NewReloader(scope constructs.Construct, props ReloaderProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: k8sapp.GetNamespaceContextPtr(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("reloader"), &cprops)

	helpers.NewHelmCached(chart, jsii.String("helm"), &helpers.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("reloader"),
		Namespace:   chart.Namespace(),
		Values: &map[string]interface{}{
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
