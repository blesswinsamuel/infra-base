package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
)

type RedisProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

func (props *RedisProps) Chart(scope packager.Construct) packager.Construct {
	cprops := packager.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("redis", cprops)

	k8sapp.NewHelm(chart, "helm", &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "redis",
		Namespace:   chart.Namespace(),
		Values: map[string]interface{}{
			"architecture": "standalone",
			"auth": map[string]interface{}{
				"enabled": false,
			},
			"metrics": map[string]interface{}{
				"enabled": true,
			},
		},
	})

	return chart
}
