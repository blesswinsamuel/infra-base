package k8sbase

import (
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"

	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type TraefikForwardAuthProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
	ImageInfo     k8sapp.ImageInfo `json:"image"`
	Ingress       struct {
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
	WhiteList      string   `json:"whilelist"`
	ImageTagSuffix string   `json:"imageTagSuffix"`
	Args           []string `json:"args"`
}

// https://github.com/k8s-at-home/charts/tree/master/charts/stable/traefik-forward-auth
// https://github.com/k8s-at-home/library-charts/tree/main/charts/stable/common
// https://github.com/thomseddon/traefik-forward-auth
func (props *TraefikForwardAuthProps) Chart(scope kubegogen.Construct) kubegogen.Construct {
	cprops := kubegogen.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("traefik-forward-auth", cprops)

	k8sapp.NewHelm(chart, "helm", &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "traefik-forward-auth",
		Namespace:   chart.Namespace(),
		Values: map[string]interface{}{
			"image": map[string]interface{}{
				"repository": props.ImageInfo.Repository,
				"tag":        strings.ReplaceAll(props.ImageInfo.Tag, "-arm", props.ImageTagSuffix),
			},
			"controller": map[string]interface{}{
				"annotations": map[string]interface{}{
					"secret.reloader.stakater.com/reload": "traefik-forward-auth",
				},
			},
			"ingress": map[string]interface{}{
				"main": map[string]interface{}{
					"enabled": true,
					"annotations": infrahelpers.MergeAnnotations(
						GetCertIssuerAnnotation(scope),
						map[string]string{"traefik.ingress.kubernetes.io/router.middlewares": "auth-traefik-forward-auth@kubernetescrd"},
					),
					"hosts": []map[string]interface{}{
						{
							"host": props.Ingress.SubDomain + "." + GetDomain(scope),
							"paths": []map[string]interface{}{
								{"path": "/"},
							},
						},
					},
					"tls": []map[string]interface{}{
						{
							"hosts": []string{
								props.Ingress.SubDomain + "." + GetDomain(scope),
							},
							"secretName": "traefik-forward-auth-tls",
						},
					},
				},
			},
			"args": props.Args,
			"env": map[string]interface{}{
				"WHITELIST":                      props.WhiteList,
				"LOG_FORMAT":                     "json",
				"LOG_LEVEL":                      "info",
				"AUTH_HOST":                      props.Ingress.SubDomain + "." + GetDomain(scope),
				"COOKIE_DOMAIN":                  GetDomain(scope),
				"PROVIDERS_GOOGLE_CLIENT_ID":     valueFromSecretKeyRef("traefik-forward-auth", "PROVIDERS_GOOGLE_CLIENT_ID"),
				"PROVIDERS_GOOGLE_CLIENT_SECRET": valueFromSecretKeyRef("traefik-forward-auth", "PROVIDERS_GOOGLE_CLIENT_SECRET"),
				"SECRET":                         valueFromSecretKeyRef("traefik-forward-auth", "SECRET"),
				// "PROVIDERS_GOOGLE_PROMPT": "select_account",
			},
			"middleware": map[string]interface{}{
				"nameOverride": "traefik-forward-auth",
			},
		},
	})

	k8sapp.NewExternalSecret(chart, "external-secret", &k8sapp.ExternalSecretProps{
		Name: "traefik-forward-auth",
		RemoteRefs: map[string]string{
			"PROVIDERS_GOOGLE_CLIENT_SECRET": "AUTH_GOOGLE_CLIENT_SECRET",
			"PROVIDERS_GOOGLE_CLIENT_ID":     "AUTH_GOOGLE_CLIENT_ID",
			"SECRET":                         "AUTH_COOKIE_SECRET",
		},
	})

	return chart
}

func valueFromSecretKeyRef(name string, key string) map[string]interface{} {
	return map[string]interface{}{
		"valueFrom": map[string]interface{}{
			"secretKeyRef": map[string]interface{}{
				"name": name,
				"key":  key,
			},
		},
	}
}
