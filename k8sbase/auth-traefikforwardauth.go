package k8sbase

import (
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"

	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type TraefikForwardAuth struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
	ImageInfo     k8sapp.ImageInfo `json:"image"`
	Ingress       struct {
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
	AllowUsers []string `json:"allowUsers"`
	Args       []string `json:"args"`
	Provider   string   `json:"provider"`
}

// https://github.com/k8s-at-home/charts/tree/master/charts/stable/traefik-forward-auth
// https://github.com/k8s-at-home/library-charts/tree/main/charts/stable/common
// https://github.com/thomseddon/traefik-forward-auth
func (props *TraefikForwardAuth) Render(scope kubegogen.Scope) {
	if props.Provider == "" {
		props.Provider = "google"
	}
	env := map[string]interface{}{
		"WHITELIST":     strings.Join(props.AllowUsers, ","),
		"LOG_FORMAT":    "json",
		"LOG_LEVEL":     "info",
		"AUTH_HOST":     props.Ingress.SubDomain + "." + GetDomain(scope),
		"COOKIE_DOMAIN": GetDomain(scope),
		"SECRET":        valueFromSecretKeyRef("traefik-forward-auth", "SECRET"),
		// "PROVIDERS_GOOGLE_PROMPT": "select_account",
	}
	externalSecretRemoteRefs := map[string]string{
		"SECRET": "FORWARD_AUTH_COOKIE_SECRET",
	}
	// https://github.com/thomseddon/traefik-forward-auth/wiki/Provider-Setup#github
	switch props.Provider {
	case "google":
		env["PROVIDERS_GOOGLE_CLIENT_ID"] = valueFromSecretKeyRef("traefik-forward-auth", "PROVIDERS_GOOGLE_CLIENT_ID")
		env["PROVIDERS_GOOGLE_CLIENT_SECRET"] = valueFromSecretKeyRef("traefik-forward-auth", "PROVIDERS_GOOGLE_CLIENT_SECRET")
		env["DEFAULT_PROVIDER"] = "google"
		externalSecretRemoteRefs["PROVIDERS_GOOGLE_CLIENT_SECRET"] = "FORWARD_AUTH_GOOGLE_CLIENT_SECRET"
		externalSecretRemoteRefs["PROVIDERS_GOOGLE_CLIENT_ID"] = "FORWARD_AUTH_GOOGLE_CLIENT_ID"
	case "github":
		env["PROVIDERS_GENERIC_OAUTH_CLIENT_ID"] = valueFromSecretKeyRef("traefik-forward-auth", "PROVIDERS_GENERIC_OAUTH_CLIENT_ID")
		env["PROVIDERS_GENERIC_OAUTH_CLIENT_SECRET"] = valueFromSecretKeyRef("traefik-forward-auth", "PROVIDERS_GENERIC_OAUTH_CLIENT_SECRET")
		env["PROVIDERS_GENERIC_OAUTH_AUTH_URL"] = "https://github.com/login/oauth/authorize"
		env["PROVIDERS_GENERIC_OAUTH_TOKEN_URL"] = "https://github.com/login/oauth/access_token"
		// env["PROVIDERS_GENERIC_OAUTH_USER_URL"] = "https://api.github.com/user"
		env["PROVIDERS_GENERIC_OAUTH_USER_URL"] = "https://api.github.com/user/emails"
		env["PROVIDERS_GENERIC_OAUTH_SCOPE"] = "user:email"
		env["DEFAULT_PROVIDER"] = "generic-oauth"
		externalSecretRemoteRefs["PROVIDERS_GENERIC_OAUTH_CLIENT_ID"] = "FORWARD_AUTH_GITHUB_CLIENT_ID"
		externalSecretRemoteRefs["PROVIDERS_GENERIC_OAUTH_CLIENT_SECRET"] = "FORWARD_AUTH_GITHUB_CLIENT_SECRET"
	}
	// TODO: remove helm dependency
	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "traefik-forward-auth",
		Namespace:   scope.Namespace(),
		Values: map[string]interface{}{
			"image": map[string]interface{}{
				"repository": props.ImageInfo.Repository,
				"tag":        props.ImageInfo.Tag,
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
			"env":  env,
			"middleware": map[string]interface{}{
				"nameOverride": "traefik-forward-auth",
			},
		},
	})

	k8sapp.NewExternalSecret(scope, &k8sapp.ExternalSecretProps{
		Name:       "traefik-forward-auth",
		RemoteRefs: externalSecretRemoteRefs,
	})
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
