package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
	corev1 "k8s.io/api/core/v1"
)

type VictoriaMetricsProps struct {
	HelmChartInfo k8sapp.ChartInfo            `json:"helm"`
	Resources     corev1.ResourceRequirements `json:"resources"`
	Ingress       struct {
		Enabled   bool   `json:"enabled"`
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
	RetentionPeriod  string         `json:"retentionPeriod"`
	PersistentVolume map[string]any `json:"persistentVolume"`
}

// https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-single
func (props *VictoriaMetricsProps) Chart(scope packager.Construct) packager.Construct {
	cprops := packager.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("victoriametrics", cprops)

	if props.PersistentVolume == nil {
		props.PersistentVolume = map[string]any{}
	}
	k8sapp.NewHelm(chart, "helm", &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "victoriametrics",
		Namespace:   chart.Namespace(),
		Values: map[string]any{
			"server": map[string]any{
				"retentionPeriod": props.RetentionPeriod,
				"statefulSet": map[string]any{
					"service": map[string]any{
						"annotations": map[string]any{
							"prometheus.io/scrape": "true",
							"prometheus.io/port":   "8428",
						},
					},
				},
				"ingress": map[string]any{
					"enabled":     props.Ingress.Enabled,
					"annotations": GetCertIssuerAnnotation(scope),
					"hosts": []map[string]any{
						{
							"name": props.Ingress.SubDomain + "." + GetDomain(scope),
							"path": "/",
							"port": "http",
						},
					},
					"tls": []map[string]any{
						{
							"hosts": []string{
								props.Ingress.SubDomain + "." + GetDomain(scope),
							},
							"secretName": "victoriametrics-tls",
						},
					},
					"pathType": "Prefix",
				},
				"extraArgs": map[string]any{
					"vmalert.proxyURL": `http://vmalert:8880`,
				},
				"resources":        props.Resources,
				"persistentVolume": props.PersistentVolume,
			},
		},
	})

	k8sapp.NewConfigMap(chart, "grafana-datasource-victoriametrics", &k8sapp.ConfigmapProps{
		Name: "grafana-datasource-victoriametrics",
		Labels: map[string]string{
			"grafana_datasource": "1",
		},
		Data: map[string]string{
			"victoriametrics.yaml": infrahelpers.ToYamlString(map[string]interface{}{
				"apiVersion": 1,
				"deleteDatasources": []map[string]interface{}{
					{
						"name":  "VictoriaMetrics",
						"orgId": 1,
					},
				},
				"datasources": []map[string]interface{}{
					{
						"name":      "VictoriaMetrics",
						"type":      "prometheus",
						"access":    "proxy",
						"orgId":     1,
						"uid":       "victoriametrics",
						"url":       "http://victoriametrics-victoria-metrics-single-server:8428",
						"isDefault": true,
						"version":   1,
						"editable":  false,
						"jsonData": map[string]interface{}{
							"timeInterval": "60s",
						},
						// # jsonData:
						// #   alertmanagerUid: alertmanager
					},
				},
			}),
		},
	})

	return chart
}
