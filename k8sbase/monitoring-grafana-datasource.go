package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/k8s"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type GrafanaDatasourceProps struct {
}

// https://github.com/prometheus-community/helm-charts/blob/main/charts/alertmanager
// https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-alert

func NewGrafanaDatasource(scope constructs.Construct, props GrafanaDatasourceProps) cdk8s.Chart {
	cprops := cdk8s.ChartProps{
		Namespace: k8sapp.GetNamespaceContextPtr(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("grafana-datasource"), &cprops)

	k8s.NewKubeConfigMap(chart, jsii.String("grafana-datasource-victoriametrics"), &k8s.KubeConfigMapProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("grafana-datasource-victoriametrics"),
			Labels: &map[string]*string{
				"grafana_datasource": jsii.String("1"),
			},
		},
		Data: &map[string]*string{
			"victoriametrics.yaml": jsii.String(infrahelpers.ToYamlString(map[string]interface{}{
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
			})),
		},
	})

	k8s.NewKubeConfigMap(chart, jsii.String("grafana-datasource-loki"), &k8s.KubeConfigMapProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("grafana-datasource-loki"),
			Labels: &map[string]*string{
				"grafana_datasource": jsii.String("1"),
			},
		},
		Data: &map[string]*string{
			"loki.yaml": jsii.String(infrahelpers.ToYamlString(map[string]interface{}{
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
			})),
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
