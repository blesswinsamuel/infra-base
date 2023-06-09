package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type ReloaderProps struct {
	Enabled       bool             `yaml:"enabled"`
	HelmChartInfo k8sapp.ChartInfo `yaml:"helm"`
}

// https://github.com/stakater/Reloader/blob/master/deployments/kubernetes/chart/reloader

func NewReloader(scope constructs.Construct, props ReloaderProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}

	chart := k8sapp.NewHelmChart(scope, jsii.String("reloader"), &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("reloader"),
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
