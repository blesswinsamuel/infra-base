package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
)

type ExternalSecretsProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/external-secrets/external-secrets/tree/main/deploy/charts/external-secrets
func (props *ExternalSecretsProps) Render(scope kgen.Scope) {
	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "external-secrets",
		Values: map[string]interface{}{
			"installCRDs": "true",
			"metrics": map[string]interface{}{
				"service": map[string]any{
					"enabled": true,
					"annotations": map[string]any{
						"prometheus.io/path":   "/metrics",
						"prometheus.io/port":   "8080",
						"prometheus.io/scrape": "true",
					},
				},
			},
			"webhook": map[string]interface{}{
				"metrics": map[string]interface{}{
					"service": map[string]any{
						"enabled": true,
						"annotations": map[string]any{
							"prometheus.io/path":   "/metrics",
							"prometheus.io/port":   "8080",
							"prometheus.io/scrape": "true",
						},
					},
				},
			},
			"certController": map[string]interface{}{
				"metrics": map[string]interface{}{
					"service": map[string]any{
						"enabled": true,
						"annotations": map[string]any{
							"prometheus.io/path":   "/metrics",
							"prometheus.io/port":   "8080",
							"prometheus.io/scrape": "true",
						},
					},
				},
			},
		},
	})
}
