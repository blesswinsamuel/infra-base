package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type GrafanaProps struct {
	HelmChartInfo        k8sapp.ChartInfo `json:"helm"`
	AnonymousAuthEnabled bool             `json:"anonymousAuthEnabled"`
	AuthProxyEnabled     bool             `json:"authProxyEnabled"`
	Namespaced           bool             `json:"namespaced"`
	DatasourceLabel      *string          `json:"datasourceLabel"`
	DatasourceLabelValue *string          `json:"datasourceLabelValue"`
	DashboardLabel       *string          `json:"dashboardLabel"`
	DashboardLabelValue  *string          `json:"dashboardLabelValue"`
	DisableSanitizeHTML  bool             `json:"disableSanitizeHTML"`
	Ingress              struct {
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
}

// https://github.com/grafana/helm-charts/tree/main/charts/grafana
func (props *GrafanaProps) Chart(scope kubegogen.Construct) kubegogen.Construct {
	cprops := kubegogen.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("grafana", cprops)

	k8sapp.NewHelm(chart, "helm", &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "grafana",
		Namespace:   chart.Namespace(),
		Values: map[string]interface{}{
			"env": infrahelpers.MergeMaps(
				map[string]string{
					"GF_SERVER_ENABLE_GZIP":                      "true",
					"GF_SECURITY_DISABLE_INITIAL_ADMIN_CREATION": "true",
					"GF_PANELS_DISABLE_SANITIZE_HTML":            infrahelpers.Ternary(props.DisableSanitizeHTML, "true", "false"),
					"GF_ANALYTICS_REPORTING_ENABLED":             "false",
					"GF_ANALYTICS_CHECK_FOR_UPDATES":             "false",
					"GF_ANALYTICS_CHECK_FOR_PLUGIN_UPDATES":      "false",
				},
				infrahelpers.Ternary(props.AnonymousAuthEnabled, map[string]string{
					"GF_AUTH_ANONYMOUS_HIDE_VERSION": "true",
					"GF_AUTH_ANONYMOUS_ENABLED":      "true",
					"GF_AUTH_ANONYMOUS_ORG_NAME":     "Main Org.",
					"GF_AUTH_ANONYMOUS_ORG_ROLE":     "Admin",
					"GF_AUTH_DISABLE_LOGIN_FORM":     "true",
				}, nil),
				infrahelpers.Ternary(props.AuthProxyEnabled, map[string]string{
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
					"enabled":         true,
					"label":           props.DatasourceLabel,
					"labelValue":      props.DatasourceLabelValue,
					"resource":        "both",
					"skipReload":      true,
					"initDatasources": true,
					"searchNamespace": "ALL",
				},
				"dashboards": map[string]interface{}{
					"enabled":          true,
					"label":            props.DashboardLabel,
					"labelValue":       props.DashboardLabelValue,
					"resource":         "both",
					"folderAnnotation": "grafana_folder",
					"provider": map[string]interface{}{
						"foldersFromFilesStructure": true,
					},
					"searchNamespace": "ALL",
					// "skipReload":     true,
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
