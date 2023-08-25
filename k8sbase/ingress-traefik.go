package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TraefikProps struct {
	ChartInfo        k8sapp.ChartInfo `json:"helm"`
	TrustedIPs       []string         `json:"trustedIPs"`
	DashboardIngress struct {
		Enabled   bool   `json:"enabled"`
		SubDomain string `json:"subDomain"`
	} `json:"dashboardIngress"`
	CreateMiddlewares struct {
		StripPrefix struct {
			Enabled bool `json:"enabled"`
		} `json:"stripPrefix"`
	} `json:"createMiddlewares"`
	DefaultMiddlewares []string `json:"defaultMiddlewares"`
}

// https://github.com/traefik/traefik-helm-chart/tree/master/traefik
func (props *TraefikProps) Chart(scope packager.Construct) packager.Construct {
	cprops := packager.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("traefik", cprops)

	k8sapp.NewHelm(chart, "traefik", &k8sapp.HelmProps{
		ChartInfo:   props.ChartInfo,
		ReleaseName: "traefik",
		Namespace:   chart.Namespace(),
		Values: map[string]interface{}{
			// "deployment": map[string]any{
			// 	"podAnnotations": map[string]any{
			// 		"prometheus.io/port":   "8082",
			// 		"prometheus.io/scrape": "true",
			// 	},
			// },
			// above is already set in the helm chart
			"providers": map[string]any{
				"kubernetesCRD": map[string]any{
					"allowCrossNamespace":       true,
					"allowExternalNameServices": true,
				},
				"kubernetesIngress": map[string]any{
					"allowExternalNameServices": true,
					"publishedService": map[string]any{
						"enabled": true,
					},
				},
			},
			"priorityClassName": "system-cluster-critical",
			"tolerations": []map[string]any{
				{
					"key":      "CriticalAddonsOnly",
					"operator": "Exists",
				},
				{
					"key":      "node-role.kubernetes.io/control-plane",
					"operator": "Exists",
					"effect":   "NoSchedule",
				},
				{
					"key":      "node-role.kubernetes.io/master",
					"operator": "Exists",
					"effect":   "NoSchedule",
				},
			},

			"ingressClass": map[string]any{
				"enabled":        true,
				"isDefaultClass": true,
			},
			"ingressRoute": map[string]any{
				"dashboard": map[string]any{
					"enabled": false,
				},
			},
			// "tlsOptions": map[string]any{
			// 	// add traefik-ssh to the default alpnProtocols
			// 	"default": map[string]any{
			// 		"alpnProtocols": []string{
			// 			"h2", "http/1.1", "acme-tls/1", "traefik-ssh",
			// 		},
			// 	},
			// },
			"additionalArguments": []string{
				"--api.insecure=true", // to expose the api for homepage dashboard via kubernetes service created below
				"--accesslog=true",
				"--accesslog.format=json",
				"--log.format=json",
				"--experimental.plugins.htransformation.modulename=github.com/tomMoulard/htransformation",
				"--experimental.plugins.htransformation.version=v0.2.7",
			},
			"experimental": map[string]any{
				"plugins": map[string]any{
					"enabled": true,
				},
			},
			"ports": map[string]any{
				// "traefik": map[string]any{
				// 	"expose": true,
				// },
				"web": map[string]any{
					"redirectTo":       "websecure",
					"forwardedHeaders": map[string]any{"trustedIPs": props.TrustedIPs},
					"proxyProtocol":    map[string]any{"trustedIPs": props.TrustedIPs},
				},
				"websecure": map[string]any{
					"forwardedHeaders": map[string]any{"trustedIPs": props.TrustedIPs},
					"proxyProtocol":    map[string]any{"trustedIPs": props.TrustedIPs},
					// ## Enable this entrypoint as a default entrypoint. When a service doesn't explicity set an entrypoint it will only use this entrypoint.
					// # works only from traefik v3
					// # asDefault: true
					"middlewares": props.DefaultMiddlewares,
				},
			},
			"service": map[string]any{
				"ipFamilyPolicy": "PreferDualStack",
				"spec": map[string]any{
					"externalTrafficPolicy": "Local", // So that traefik gets the real IP - https://github.com/k3s-io/k3s/discussions/2997#discussioncomment-413904
				},
			},
			"extraObjects": []map[string]any{
				{
					"apiVersion": "v1",
					"kind":       "Service",
					"metadata": map[string]any{
						"name": "traefik-api",
					},
					"spec": map[string]any{
						"type": "ClusterIP",
						"selector": map[string]any{
							"app.kubernetes.io/name":     "traefik",
							"app.kubernetes.io/instance": "traefik-ingress",
						},
						"ports": []map[string]any{
							{
								"port":       8080,
								"name":       "traefik",
								"targetPort": 9000,
								"protocol":   "TCP",
							},
						},
					},
				},
			},
		},
	})

	if props.DashboardIngress.Enabled {
		k8sapp.NewIngress(chart, "traefik-dashboard-external", &k8sapp.IngressProps{
			Name: "traefik-dashboard",
			Hosts: []k8sapp.IngressHost{
				{Host: props.DashboardIngress.SubDomain + "." + GetDomain(chart), Paths: []k8sapp.IngressHostPath{{Path: "/", ServiceName: "api@internal"}}, Tls: true},
			},
			IngressType: "traefik",
		})
	}

	if props.CreateMiddlewares.StripPrefix.Enabled {
		k8sapp.NewK8sObject(chart, "traefik-strip-prefix", &traefikv1alpha1.Middleware{
			ObjectMeta: v1.ObjectMeta{Name: "traefik-strip-prefix"},
			Spec: traefikv1alpha1.MiddlewareSpec{
				StripPrefixRegex: &dynamic.StripPrefixRegex{
					Regex: []string{"^/[^/]+"},
				},
			},
		})
	}

	return chart
}

// # # https://github.com/traefik/traefik/issues/5571#issuecomment-539393453 - affects wss in goatcounter
// # apiVersion: traefik.containo.us/v1alpha1
// # kind: Middleware
// # metadata:
// #   name: ssl-header
// #   namespace: ingress
// # spec:
// #   headers:
// #     customRequestHeaders:
// #       X-Forwarded-Proto: https,wss
