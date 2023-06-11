package k8sbase

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
)

type NodeExporterProps struct {
	Enabled       bool             `json:"enabled"`
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

func NewNodeExporter(scope packager.Construct, props NodeExporterProps) packager.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := &packager.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := packager.NewChart(scope, "node-exporter", cprops)

	k8sapp.NewHelm(chart, jsii.String("helm"), &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("node-exporter"),
		Namespace:   chart.Namespace(),
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

	return chart
}
