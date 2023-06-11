package k8sbase

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
)

type GrafanaDatasourceProps struct {
}

// https://github.com/prometheus-community/helm-charts/blob/main/charts/alertmanager
// https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-alert

func NewGrafanaDatasource(scope packager.Construct, props GrafanaDatasourceProps) packager.Chart {
	cprops := &packager.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := packager.NewChart(scope, "grafana-datasource", cprops)

	k8sapp.NewConfigMap(chart, jsii.String("grafana-datasource-victoriametrics"), &k8sapp.ConfigmapProps{
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
						// # jsonData:
						// #   alertmanagerUid: alertmanager
					},
				},
			}),
		},
	})

	k8sapp.NewConfigMap(chart, jsii.String("grafana-datasource-loki"), &k8sapp.ConfigmapProps{
		Name: "grafana-datasource-loki",
		Labels: map[string]string{
			"grafana_datasource": "1",
		},
		Data: map[string]string{
			"loki.yaml": infrahelpers.ToYamlString(map[string]interface{}{
				"apiVersion": 1,
				"deleteDatasources": []map[string]interface{}{
					{
						"name":  "Loki",
						"orgId": 1,
					},
				},
				"datasources": []map[string]interface{}{
					{
						"name":   "Loki",
						"type":   "loki",
						"access": "proxy",
						"orgId":  1,
						"uid":    "loki",
						"url":    "http://loki:3100",
						"jsonData": map[string]interface{}{
							"maxLines": 1000,
							// # alertmanagerUid: alertmanager
						},
					},
				},
			}),
		},
	})

	return chart
}

// {{ define "cluster-base.monitoring.datasource.alertmanager" }}
// ---
// # apiVersion: v1
// # kind: ConfigMap
// # metadata:
// #   name: grafana-datasource-alertmanager
// #   namespace: monitoring
// #   labels:
// #     grafana_datasource: "1"
// # data:
// #   loki.yaml: |-
// #     apiVersion: 1

// #     deleteDatasources:
// #       - name: Alertmanager
// #         orgId: 1

// #     datasources:
// #       - name: Alertmanager
// #         type: alertmanager
// #         access: proxy
// #         orgId: 1
// #         uid: alertmanager
// #         url: http://vmalert-alertmanager:9093
// #         jsonData:
// #           implementation: prometheus
// {{- end }}
