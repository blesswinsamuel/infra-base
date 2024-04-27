package k8sbase

import (
	"fmt"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	traefikv1alpha1 "github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

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
	env := map[string]string{
		"WHITELIST":     strings.Join(props.AllowUsers, ","),
		"TZ":            "UTC",
		"LOG_FORMAT":    "json",
		"LOG_LEVEL":     "info",
		"AUTH_HOST":     props.Ingress.SubDomain + "." + GetDomain(scope),
		"COOKIE_DOMAIN": GetDomain(scope),
		// "PROVIDERS_GOOGLE_PROMPT": "select_account",
	}
	externalSecretRemoteRefs := map[string]string{
		"SECRET": "FORWARD_AUTH_COOKIE_SECRET",
	}
	// https://github.com/thomseddon/traefik-forward-auth/wiki/Provider-Setup#github
	switch props.Provider {
	case "google":
		env["DEFAULT_PROVIDER"] = "google"
		externalSecretRemoteRefs["PROVIDERS_GOOGLE_CLIENT_SECRET"] = "FORWARD_AUTH_GOOGLE_CLIENT_SECRET"
		externalSecretRemoteRefs["PROVIDERS_GOOGLE_CLIENT_ID"] = "FORWARD_AUTH_GOOGLE_CLIENT_ID"
	case "github":
		env["PROVIDERS_GENERIC_OAUTH_AUTH_URL"] = "https://github.com/login/oauth/authorize"
		env["PROVIDERS_GENERIC_OAUTH_TOKEN_URL"] = "https://github.com/login/oauth/access_token"
		// env["PROVIDERS_GENERIC_OAUTH_USER_URL"] = "https://api.github.com/user"
		env["PROVIDERS_GENERIC_OAUTH_USER_URL"] = "https://api.github.com/user/emails"
		env["PROVIDERS_GENERIC_OAUTH_SCOPE"] = "user:email"
		env["DEFAULT_PROVIDER"] = "generic-oauth"
		externalSecretRemoteRefs["PROVIDERS_GENERIC_OAUTH_CLIENT_ID"] = "FORWARD_AUTH_GITHUB_CLIENT_ID"
		externalSecretRemoteRefs["PROVIDERS_GENERIC_OAUTH_CLIENT_SECRET"] = "FORWARD_AUTH_GITHUB_CLIENT_SECRET"
	}
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Name:               "traefik-forward-auth",
		DNSPolicy:          corev1.DNSClusterFirst,
		EnableServiceLinks: infrahelpers.Ptr(true),
		Containers: []k8sapp.ApplicationContainer{
			{
				Name:             "traefik-forward-auth",
				Image:            props.ImageInfo,
				Ports:            []k8sapp.ContainerPort{{Name: "http", Port: 4181, Ingress: &k8sapp.ApplicationIngress{Host: props.Ingress.SubDomain + "." + GetDomain(scope)}}},
				Args:             props.Args,
				Env:              env,
				EnvFromSecretRef: []string{"traefik-forward-auth"},
				LivenessProbe:    &corev1.Probe{InitialDelaySeconds: 0, FailureThreshold: 3, PeriodSeconds: 10, TimeoutSeconds: 1, ProbeHandler: corev1.ProbeHandler{TCPSocket: &corev1.TCPSocketAction{Port: intstr.FromString("http")}}},
				ReadinessProbe:   &corev1.Probe{InitialDelaySeconds: 0, FailureThreshold: 3, PeriodSeconds: 10, TimeoutSeconds: 1, ProbeHandler: corev1.ProbeHandler{TCPSocket: &corev1.TCPSocketAction{Port: intstr.FromString("http")}}},
				StartupProbe:     &corev1.Probe{InitialDelaySeconds: 0, FailureThreshold: 30, PeriodSeconds: 5, TimeoutSeconds: 1, ProbeHandler: corev1.ProbeHandler{TCPSocket: &corev1.TCPSocketAction{Port: intstr.FromString("http")}}},
			},
		},
		ExternalSecrets: []k8sapp.ApplicationExternalSecret{
			{Name: "traefik-forward-auth", RemoteRefs: externalSecretRemoteRefs},
		},
	})
	scope.AddApiObject(&traefikv1alpha1.Middleware{
		ObjectMeta: metav1.ObjectMeta{Name: "traefik-forward-auth"},
		Spec: traefikv1alpha1.MiddlewareSpec{
			ForwardAuth: &traefikv1alpha1.ForwardAuth{
				Address:             fmt.Sprintf("http://%s.%s:4181", "traefik-forward-auth", scope.Namespace()),
				AuthResponseHeaders: []string{"X-Forwarded-User"},
			},
		},
	})
}
