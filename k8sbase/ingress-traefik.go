package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/certmanagerio"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/ingressroute_traefikcontainous"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/middlewares_traefikcontainous"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type TraefikProps struct {
	ChartInfo        ChartInfo `yaml:"helm"`
	TrustedIPs       []string  `yaml:"trustedIPs"`
	DashboardIngress struct {
		Enabled   bool   `yaml:"enabled"`
		SubDomain string `yaml:"subDomain"`
	} `yaml:"dashboardIngress"`
	Middlewares struct {
		BasicAuth struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"basicAuth"`
		StripPrefix struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"stripPrefix"`
	} `yaml:"middlewares"`
}

// https://github.com/traefik/traefik-helm-chart/tree/master/traefik
func NewTraefik(scope constructs.Construct, props TraefikProps) cdk8s.Chart {
	cprops := cdk8s.ChartProps{
		Namespace: jsii.String("kube-system"),
	}
	chart := cdk8s.NewChart(scope, jsii.String("traefik"), &cprops)

	NewHelmCached(chart, jsii.String("traefik"), &HelmProps{
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
					"allowCrossNamespace": true,
				},
				"kubernetesIngress": map[string]any{
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
		certmanagerio.NewCertificate(chart, jsii.String("traefik-dashboard-cert"), &certmanagerio.CertificateProps{
			Metadata: &cdk8s.ApiObjectMetadata{
				Name:      jsii.String("traefik-dashboard"),
				Namespace: jsii.String("kube-system"),
			},
			Spec: &certmanagerio.CertificateSpec{
				DnsNames: jsii.PtrSlice(
					props.DashboardIngress.SubDomain + "." + GetDomain(chart),
				),
				SecretName: jsii.String("traefik-dashboard-tls"),
				IssuerRef: &certmanagerio.CertificateSpecIssuerRef{
					Name: jsii.String(GetCertIssuer(scope)),
					Kind: jsii.String("ClusterIssuer"),
				},
			},
		})

		ingressroute_traefikcontainous.NewIngressRoute(chart, jsii.String("traefik-dashboard-external"), &ingressroute_traefikcontainous.IngressRouteProps{
			Metadata: &cdk8s.ApiObjectMetadata{
				Name:      jsii.String("traefik-dashboard-external"),
				Namespace: jsii.String("kube-system"),
			},
			Spec: &ingressroute_traefikcontainous.IngressRouteSpec{
				EntryPoints: jsii.PtrSlice("websecure"),
				Routes: &[]*ingressroute_traefikcontainous.IngressRouteSpecRoutes{
					{
						Match: jsii.String("Host(`" + props.DashboardIngress.SubDomain + "." + GetDomain(chart) + "`) && (PathPrefix(`/dashboard`) || PathPrefix(`/api`))"),
						Kind:  ingressroute_traefikcontainous.IngressRouteSpecRoutesKind_RULE,
						Services: &[]*ingressroute_traefikcontainous.IngressRouteSpecRoutesServices{
							{
								Name: jsii.String("api@internal"),
								Kind: ingressroute_traefikcontainous.IngressRouteSpecRoutesServicesKind_TRAEFIK_SERVICE,
							},
						},
						Middlewares: &[]*ingressroute_traefikcontainous.IngressRouteSpecRoutesMiddlewares{
							GetTraefikCRAuthMiddleware(chart),
						},
					},
				},
				Tls: &ingressroute_traefikcontainous.IngressRouteSpecTls{
					SecretName: jsii.String("traefik-dashboard-tls"),
					Domains: &[]*ingressroute_traefikcontainous.IngressRouteSpecTlsDomains{
						{
							Main: jsii.String(props.DashboardIngress.SubDomain + "." + GetDomain(chart)),
						},
					},
				},
			},
		})
	}

	if props.Middlewares.BasicAuth.Enabled {
		middlewares_traefikcontainous.NewMiddleware(chart, jsii.String("traefik-basic-auth"), &middlewares_traefikcontainous.MiddlewareProps{
			Metadata: &cdk8s.ApiObjectMetadata{
				Name:      jsii.String("traefik-basic-auth"),
				Namespace: jsii.String("kube-system"),
			},
			Spec: &middlewares_traefikcontainous.MiddlewareSpec{
				BasicAuth: &middlewares_traefikcontainous.MiddlewareSpecBasicAuth{
					Secret:       jsii.String("traefik-basic-auth"),
					RemoveHeader: jsii.Bool(true),
				},
			},
		})
		NewExternalSecret(chart, jsii.String("traefik-basic-auth-secret"), &ExternalSecretProps{
			Name:            jsii.String("traefik-basic-auth"),
			Namespace:       jsii.String("kube-system"),
			RefreshInterval: jsii.String("2m"),
			Secrets: map[string]string{
				"users": "TRAEFIK_BASIC_AUTH_USERS",
			},
		})
	}

	if props.Middlewares.StripPrefix.Enabled {
		middlewares_traefikcontainous.NewMiddleware(chart, jsii.String("traefik-strip-prefix"), &middlewares_traefikcontainous.MiddlewareProps{
			Metadata: &cdk8s.ApiObjectMetadata{
				Name:      jsii.String("traefik-strip-prefix"),
				Namespace: jsii.String("kube-system"),
			},
			Spec: &middlewares_traefikcontainous.MiddlewareSpec{
				StripPrefixRegex: &middlewares_traefikcontainous.MiddlewareSpecStripPrefixRegex{
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
// #   namespace: kube-system
// # spec:
// #   headers:
// #     customRequestHeaders:
// #       X-Forwarded-Proto: https,wss
