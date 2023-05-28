package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/helpers"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/middlewares_traefikio"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type TraefikProps struct {
	Enabled          bool              `yaml:"enabled"`
	ChartInfo        helpers.ChartInfo `yaml:"helm"`
	TrustedIPs       []string          `yaml:"trustedIPs"`
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

// https://github.com/traefik/traefik-helm-chart/tree/master/traefik
func NewTraefik(scope constructs.Construct, props TraefikProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: helpers.GetNamespace(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("traefik"), &cprops)

	helpers.NewHelmCached(chart, jsii.String("traefik"), &helpers.HelmProps{
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
				"--accesslog=true",
				"--accesslog.format=json",
				"--log.format=json",
			},
			"ports": map[string]any{
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
						GetTraefikAuthMiddlewareName(scope),
					},
				},
			},
			"service": map[string]any{
				"ipFamilyPolicy": "PreferDualStack",
				"spec": map[string]any{
					"externalTrafficPolicy": "Local", // So that traefik gets the real IP - https://github.com/k3s-io/k3s/discussions/2997#discussioncomment-413904
				},
			},
		},
	})

	if props.DashboardIngress.Enabled {
		helpers.NewIngress(chart, jsii.String("traefik-dashboard-external"), &helpers.IngressProps{
			Name: "traefik-dashboard",
			Hosts: []helpers.Host{
				{Host: props.DashboardIngress.SubDomain + "." + GetDomain(chart), Paths: []helpers.Path{{Path: "/", ServiceName: "api@internal"}}, Tls: true},
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
