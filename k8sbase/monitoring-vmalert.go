package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type VmalertProps struct {
	Enabled       bool       `yaml:"enabled"`
	HelmChartInfo ChartInfo  `yaml:"helm"`
	Resources     *Resources `yaml:"resources"`
	Ingress       struct {
		Enabled   bool   `yaml:"enabled"`
		SubDomain string `yaml:"subDomain"`
	} `yaml:"ingress"`
}

// https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-agent
func NewVmalert(scope constructs.Construct, props VmalertProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: GetNamespace(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("vmalert"), &cprops)

	NewHelmCached(chart, jsii.String("helm"), &HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("vmalert"),
		Namespace:   chart.Namespace(),
		Values: &map[string]any{
			"alertmanager": map[string]any{
				"enabled": false,
			},
			"server": map[string]any{
				"service": map[string]any{
					"annotations": map[string]any{
						"prometheus.io/scrape": "true",
						"prometheus.io/port":   "8880",
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
							"secretName": "vmalert-tls",
						},
					},
					"pathType": "Prefix",
				},
				"configMap": "alerting-rules",
				"datasource": map[string]any{
					"url": "http://victoriametrics-victoria-metrics-single-server:8428",
				},
				"notifier": map[string]any{
					"alertmanager": map[string]any{
						"url": "http://alertmanager:9093",
					},
				},
				"remote": map[string]any{
					"write": map[string]any{
						"url": "http://victoriametrics-victoria-metrics-single-server:8428",
					},
					"read": map[string]any{
						"url": "http://victoriametrics-victoria-metrics-single-server:8428",
					},
				},
				"extraArgs": map[string]any{
					"rule":         "/config/*.yaml",
					"external.url": "https://" + "grafana" + "." + GetDomain(scope),
					// # external.alert.source: explore?orgId=1&left=["now-1h","now","VictoriaMetrics",{"expr":""},{"mode":"Metrics"},{"ui":[true,true,true,"none"]}]
					// # external.alert.source: {{ `explore?orgId=1&left=["now-1h","now","VictoriaMetrics",{"expr":"{{$expr|quotesEscape|pathEscape}}"}]` }}
					// https://github.com/VictoriaMetrics/VictoriaMetrics/blob/8edb390e215cbffe9bb267eea8337dbf1df1c76f/deployment/docker/docker-compose.yml#L75
					"external.alert.source": `explore?orgId=1&left={"datasource":"VictoriaMetrics","queries":[{"expr":"{{$expr|quotesEscape|crlfEscape|queryEscape}}","refId":"A"}],"range":{"from":"now-1h","to":"now"}}`,
					// # - "-external.label=env=${ENV_NAME}"
					// # - "-evaluationInterval=30s"
					// # - "-rule=/config/alert_rules.yml"
				},
				"resources": props.Resources,
			},
		},
	})

	return chart
}
