package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	corev1 "k8s.io/api/core/v1"
)

type VictoriametricsProps struct {
	Enabled       bool                        `json:"enabled"`
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
func NewVictoriaMetrics(scope constructs.Construct, props VictoriametricsProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: k8sapp.GetNamespaceContextPtr(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("victoriametrics"), &cprops)

	k8sapp.NewHelm(chart, jsii.String("helm"), &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("victoriametrics"),
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
					"vmalert.proxyURL": `http://vmalert-victoria-metrics-alert-server:8880`,
				},
				"resources":        props.Resources,
				"persistentVolume": props.PersistentVolume,
			},
		},
	})

	return chart
}
