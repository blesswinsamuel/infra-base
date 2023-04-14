package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type GrafanaProps struct {
	Enabled              bool      `yaml:"enabled"`
	HelmChartInfo        ChartInfo `yaml:"helm"`
	AnonymousAuthEnabled bool      `yaml:"anonymousAuthEnabled"`
	AuthProxyEnabled     bool      `yaml:"authProxyEnabled"`
	Namespaced           bool      `yaml:"namespaced"`
	DatasourceLabel      *string   `yaml:"datasourceLabel"`
	DatasourceLabelValue *string   `yaml:"datasourceLabelValue"`
	DashboardLabel       *string   `yaml:"dashboardLabel"`
	DashboardLabelValue  *string   `yaml:"dashboardLabelValue"`
	Ingress              struct {
		SubDomain string `yaml:"subDomain"`
	} `yaml:"ingress"`
}

// https://github.com/grafana/helm-charts
func NewGrafana(scope constructs.Construct, props GrafanaProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: GetNamespace(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("grafana"), &cprops)

	NewHelmCached(chart, jsii.String("helm"), &HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("grafana"),
		Namespace:   chart.Namespace(),
		Values: &map[string]interface{}{
			"env": MergeMaps(
				map[string]string{
					"GF_SERVER_ENABLE_GZIP":                      "true",
					"GF_SECURITY_DISABLE_INITIAL_ADMIN_CREATION": "true",
				},
				Ternary(props.AnonymousAuthEnabled, map[string]string{
					"GF_AUTH_ANONYMOUS_HIDE_VERSION": "true",
					"GF_AUTH_ANONYMOUS_ENABLED":      "true",
					"GF_AUTH_ANONYMOUS_ORG_NAME":     "Main Org.",
					"GF_AUTH_ANONYMOUS_ORG_ROLE":     "Admin",
					"GF_AUTH_DISABLE_LOGIN_FORM":     "true",
				}, nil),
				Ternary(props.AuthProxyEnabled, map[string]string{
					"GF_AUTH_PROXY_ENABLED":            "true",
					"GF_AUTH_PROXY_HEADER_NAME":        "Remote-User",
					"GF_AUTH_PROXY_HEADER_PROPERTY":    "username",
					"GF_AUTH_PROXY_AUTO_SIGN_UP":       "true",
					"GF_AUTH_PROXY_HEADERS":            "Groups:Remote-Group",
					"GF_AUTH_PROXY_ENABLE_LOGIN_TOKEN": "false",
				}, nil),
			),
			"podAnnotations": map[string]string{
				"prometheus.io/scrape": "true",
				"prometheus.io/port":   "3000",
			},
			"ingress": map[string]interface{}{
				"enabled": true,
				"hosts": []string{
					props.Ingress.SubDomain + "." + GetDomain(scope),
				},
				"annotations": GetCertIssuerAnnotation(scope),
				"tls": []map[string]interface{}{
					{
						"hosts": []string{
							props.Ingress.SubDomain + "." + GetDomain(scope),
						},
						"secretName": "grafana-tls",
					},
				},
			},
			"sidecar": map[string]interface{}{
				"datasources": map[string]interface{}{
					"enabled":    true,
					"label":      props.DatasourceLabel,
					"labelValue": props.DatasourceLabelValue,
					"resource":   "configmap",
				},
				"dashboards": map[string]interface{}{
					"enabled":          true,
					"label":            props.DashboardLabel,
					"labelValue":       props.DashboardLabelValue,
					"resource":         "configmap",
					"folderAnnotation": "grafana_folder",
					"provider": map[string]interface{}{
						"foldersFromFilesStructure": true,
					},
				},
			},
			"rbac": map[string]interface{}{"namespaced": props.Namespaced, "pspEnabled": false},
			"testFramework": map[string]interface{}{
				"enabled": false,
			},
		},
	})

	return chart
}
