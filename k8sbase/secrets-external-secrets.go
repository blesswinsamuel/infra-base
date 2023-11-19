package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type ExternalSecretsProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/external-secrets/external-secrets/tree/main/deploy/charts/external-secrets
func (props *ExternalSecretsProps) Chart(scope kubegogen.Construct) kubegogen.Construct {
	cprops := kubegogen.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("external-secrets", cprops)

	k8sapp.NewHelm(chart, "helm", &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "external-secrets",
		Namespace:   chart.Namespace(),
		Values: map[string]interface{}{
			"installCRDs": "true",
			"prometheus": map[string]interface{}{
				"enabled": "true",
			},
			"webhook": map[string]interface{}{
				"prometheus": map[string]interface{}{
					"enabled": "true",
				},
			},
			"certController": map[string]interface{}{
				"prometheus": map[string]interface{}{
					"enabled": "true",
				},
			},
		},
	})

	return chart
}
