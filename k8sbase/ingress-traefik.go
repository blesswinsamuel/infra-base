package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/k8simports/middlewares_traefikio"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type TraefikProps struct {
	Enabled          bool             `yaml:"enabled"`
	ChartInfo        k8sapp.ChartInfo `yaml:"helm"`
	TrustedIPs       []string         `yaml:"trustedIPs"`
	DashboardIngress struct {
		Enabled   bool   `yaml:"enabled"`
		SubDomain string `yaml:"subDomain"`
	} `yaml:"dashboardIngress"`
	Middlewares struct {
		StripPrefix struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"stripPrefix"`
	} `yaml:"middlewares"`
}

func getTraefikAuthMiddlewareName(scope constructs.Construct) string {
	switch GetGlobal(scope).InternetAuthType {
	case "traefik-forward-auth":
		return "auth-traefik-forward-auth@kubernetescrd"
	case "authelia":
		return "auth-forwardauth-authelia@kubernetescrd"
	}
	panic("Invalid internetAuthType")
}

// https://github.com/traefik/traefik-helm-chart/tree/master/traefik
func NewTraefik(scope constructs.Construct, props TraefikProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: k8sapp.GetNamespaceContextPtr(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("traefik"), &cprops)

	k8sapp.NewHelm(chart, jsii.String("traefik"), &k8sapp.HelmProps{
		ChartInfo:   props.ChartInfo,
		ReleaseName: jsii.String("traefik"),
		Namespace:   chart.Namespace(),
		Values: &map[string]interface{}{
			"deployment": map[string]any{
				"podAnnotations": map[string]any{
					"prometheus.io/port":   "8082",
					"prometheus.io/scrape": "true",
				},
			},
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
					"middlewares": []string{
						getTraefikAuthMiddlewareName(scope),
					},
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
		k8sapp.NewIngress(chart, jsii.String("traefik-dashboard-external"), &k8sapp.IngressProps{
			Name: "traefik-dashboard",
			Hosts: []k8sapp.IngressHost{
				{Host: props.DashboardIngress.SubDomain + "." + GetDomain(chart), Paths: []k8sapp.IngressHostPath{{Path: "/", ServiceName: "api@internal"}}, Tls: true},
			},
			IngressType: "traefik",
		})
	}

	if props.Middlewares.StripPrefix.Enabled {
		middlewares_traefikio.NewMiddleware(chart, jsii.String("traefik-strip-prefix"), &middlewares_traefikio.MiddlewareProps{
			Metadata: &cdk8s.ApiObjectMetadata{
				Name: jsii.String("traefik-strip-prefix"),
			},
			Spec: &middlewares_traefikio.MiddlewareSpec{
				StripPrefixRegex: &middlewares_traefikio.MiddlewareSpecStripPrefixRegex{
					Regex: jsii.PtrSlice("^/[^/]+"),
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
