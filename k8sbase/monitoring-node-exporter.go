package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type NodeExporterProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

func (props *NodeExporterProps) Chart(scope kubegogen.Construct) kubegogen.Construct {
	cprops := kubegogen.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("node-exporter", cprops)

	k8sapp.NewHelm(chart, "helm", &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "node-exporter",
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
