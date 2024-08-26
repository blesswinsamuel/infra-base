package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TraefikProps struct {
	ChartInfo        k8sapp.ChartInfo `json:"helm"`
	TrustedIPs       []string         `json:"trustedIPs"`
	ServiceType      string           `json:"serviceType"`
	DashboardIngress struct {
		Enabled   bool   `json:"enabled"`
		SubDomain string `json:"subDomain"`
	} `json:"dashboardIngress"`
	CreateMiddlewares struct {
		HSTS struct {
			Enabled bool `json:"enabled"`
		} `json:"hsts"`
		StripPrefix struct {
			Enabled bool `json:"enabled"`
		} `json:"stripPrefix"`
	} `json:"createMiddlewares"`
	DefaultMiddlewares         []string `json:"defaultMiddlewares"`
	Plugins                    []string `json:"plugins"`
	DisableHttpToHttpsRedirect bool     `json:"disableHttpToHttpsRedirect"`
	DefaultTlsStore            string   `json:"defaultTlsStore"`
	// HostPathMountForLogs bool     `json:"hostPathMountForLogs"`
}

// https://github.com/traefik/traefik-helm-chart/tree/master/traefik
func (props *TraefikProps) Render(scope kgen.Scope) {
	if props.ServiceType == "" {
		props.ServiceType = "LoadBalancer"
	}

	values := map[string]interface{}{
		"deployment": map[string]any{
			// 	"podAnnotations": map[string]any{
			// 		"prometheus.io/port":   "8082",
			// 		"prometheus.io/scrape": "true",
			// 	},
			"initContainers": []map[string]any{
				{
					"name":    "sleeper",
					"image":   "alpine:3.12",
					"command": []string{"sh", "-c", "sleep 1"}, // to wait for networking to be ready to download plugins
				},
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
		},
		"experimental": map[string]any{
			"plugins": map[string]any{},
		},
		"ports": map[string]any{
			// "traefik": map[string]any{
			// 	"expose": true,
			// },
			"web": map[string]any{
				// "expose":           map[string]any{"default": true},
				// "exposedPort":      80,
				"redirectTo":       map[string]any{"port": "websecure"},
				"forwardedHeaders": map[string]any{"trustedIPs": props.TrustedIPs},
				"proxyProtocol":    map[string]any{"trustedIPs": props.TrustedIPs},
			},
			"websecure": map[string]any{
				// "expose":           map[string]any{"default": true},
				// "exposedPort":      443,
				"forwardedHeaders": map[string]any{"trustedIPs": props.TrustedIPs},
				"proxyProtocol":    map[string]any{"trustedIPs": props.TrustedIPs},
				// ## Enable this entrypoint as a default entrypoint. When a service doesn't explicity set an entrypoint it will only use this entrypoint.
				// # works only from traefik v3
				// # asDefault: true
				"middlewares": props.DefaultMiddlewares,
			},
		},
		"service": map[string]any{
			"type":           props.ServiceType,
			"ipFamilyPolicy": "PreferDualStack",
			"spec": map[string]any{
				"externalTrafficPolicy": infrahelpers.If(props.ServiceType == "LoadBalancer", "Local", ""), // So that traefik gets the real IP - https://github.com/k3s-io/k3s/discussions/2997#discussioncomment-413904
				// also see https://www.authelia.com/integration/kubernetes/introduction/#external-traffic-policy
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
		"logs": map[string]any{
			"general": map[string]any{"format": "json"},
		},
	}
	if props.DefaultTlsStore != "" {
		values["tlsStore"] = map[string]any{
			"default": map[string]any{
				"defaultCertificate": map[string]any{
					"secretName": props.DefaultTlsStore,
				},
			},
		}
	}
	if props.DisableHttpToHttpsRedirect {
		delete(values["ports"].(map[string]any)["web"].(map[string]any), "redirectTo")
		delete(values["ports"].(map[string]any), "websecure")
	}
	values["logs"].(map[string]any)["access"] = map[string]any{
		"enabled": true,
		"format":  "json",
	}

	plugins := values["experimental"].(map[string]any)["plugins"].(map[string]any)
	for _, plugin := range props.Plugins {
		switch plugin {
		case "htransformation":
			plugins[plugin] = map[string]any{
				"moduleName": "github.com/tomMoulard/htransformation",
				"version":    "v0.2.8",
			}
		case "crowdsec-bouncer":
			plugins[plugin] = map[string]any{
				"moduleName": "github.com/maxlerebourg/crowdsec-bouncer-traefik-plugin",
				"version":    "v1.3.3",
			}
		case "request-headers":
			plugins[plugin] = map[string]any{
				"moduleName": "github.com/traefik/plugindemo",
				"version":    "v0.2.2",
			}
			// plugins[plugin] = map[string]any{
			// 	"moduleName": "github.com/blesswinsamuel/traefik-request-headers-middleware",
			// 	"version":    "v0.0.6",
			// }
		}
	}

	// if props.HostPathMountForLogs {
	// 	values["logs"].(map[string]any)["access"] = map[string]any{
	// 		"enabled": true,
	// 		"format":  "json",
	// 		// "bufferingSize": 100,
	// 		// "filters":       map[string]any{"statuscodes": "204-299,400-499,500-59"},
	// 		// "filePath":      "/var/log/traefik/access.log",
	// 	}
	// 	values["additionalVolumeMounts"] = []corev1.VolumeMount{
	// 		{Name: "access-logs", MountPath: "/var/log/traefik"},
	// 	}
	// 	values["additionalVolumes"] = []corev1.Volume{
	// 		{Name: "access-logs", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/var/log/traefik"}}},
	// 	}
	// }

	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
		ChartInfo:   props.ChartInfo,
		ReleaseName: "traefik",
		Values:      values,
	})

	if props.DashboardIngress.Enabled {
		k8sapp.NewIngress(scope, &k8sapp.IngressProps{
			Name: "traefik-dashboard",
			Hosts: []k8sapp.IngressHost{
				{Host: props.DashboardIngress.SubDomain + "." + k8sapp.GetDomain(scope), Paths: []k8sapp.IngressHostPath{{Path: "/", ServiceName: "api@internal"}}, Tls: true},
			},
			IngressType:    "traefik",
			UseDefaultCert: true,
			Annotations: k8sapp.GetHomepageAnnotations(&k8sapp.ApplicationHomepage{
				Name:        "Traefik",
				Description: "Reverse proxy",
				Widget: map[string]string{
					"type": "traefik",
					"url":  "http://traefik-api." + scope.Namespace() + ".svc.cluster.local:8080",
				},
				Href:        "https://" + props.DashboardIngress.SubDomain + "." + k8sapp.GetDomain(scope),
				PodSelector: "app.kubernetes.io/name=traefik",
				Icon:        "traefik",
				Group:       "Infra",
			}),
		})
	}

	if props.CreateMiddlewares.StripPrefix.Enabled {
		scope.AddApiObject(&traefikv1alpha1.Middleware{
			ObjectMeta: metav1.ObjectMeta{Name: "traefik-strip-prefix"},
			Spec: traefikv1alpha1.MiddlewareSpec{
				StripPrefixRegex: &dynamic.StripPrefixRegex{
					Regex: []string{"^/[^/]+"},
				},
			},
		})
	}
	// https://docs.nextcloud.com/server/29/admin_manual/installation/harden_server.html#enable-http-strict-transport-security
	if props.CreateMiddlewares.HSTS.Enabled {
		scope.AddApiObject(&traefikv1alpha1.Middleware{
			ObjectMeta: metav1.ObjectMeta{Name: "traefik-hsts"},
			Spec: traefikv1alpha1.MiddlewareSpec{
				Headers: &dynamic.Headers{
					STSSeconds:           15552000,
					STSIncludeSubdomains: true,
					STSPreload:           true,
					ForceSTSHeader:       true,
				},
			},
		})
	}

	// scope.AddApiObject(&networkingv1.NetworkPolicy{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name: "allow-all-svclbtraefik-ingress",
	// 	},
	// 	Spec: networkingv1.NetworkPolicySpec{
	// 		PodSelector: v1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "traefik"}},
	// 		Ingress: []networkingv1.NetworkPolicyIngressRule{
	// 			{},
	// 		},
	// 		PolicyTypes: []networkingv1.PolicyType{networkingv1.PolicyTypeIngress},
	// 	},
	// })
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
