package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type CrowdsecProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/crowdsecurity/helm-charts/tree/main/charts/crowdsec
func (props *CrowdsecProps) Chart(scope kubegogen.Scope) kubegogen.Scope {
	cprops := kubegogen.ScopeProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.CreateScope("crowdsec", cprops)

	k8sapp.NewHelm(chart, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "crowdsec",
		Namespace:   chart.Namespace(),
		Values: map[string]any{
			"container_runtime": "containerd",
			"lapi": map[string]any{
				"service": map[string]any{
					"annotations": map[string]any{
						"prometheus.io/scrape": "true",
						"prometheus.io/port":   "6060",
					},
				},
				"metrics": map[string]any{
					"enabled": true,
				},
				"env": []map[string]any{
					// {"name": "ENROLL_KEY", "valueFrom": map[string]any{"secretKeyRef": map[string]any{"name": "crowdsec-keys", "key": "ENROLL_KEY"}}},
					// {"name": "ENROLL_INSTANCE_NAME", "value": "homelab"},
					// {"name": "DISABLE_ONLINE_API", "value": "false"},
				},
				"ingress": map[string]any{
					"enabled":     false,
					"annotations": GetCertIssuerAnnotation(scope),
					"host":        "crowdsec-lapi" + "." + GetDomain(scope),
					"tls": []map[string]any{
						{
							"hosts":      []string{"crowdsec-lapi" + "." + GetDomain(scope)},
							"secretName": "crowdsec-lapi-tls",
						},
					},
				},
				"dashboard": map[string]any{
					"enabled": false,
					"ingress": map[string]any{
						"enabled":     true,
						"annotations": GetCertIssuerAnnotation(scope),
						"host":        "crowdsec-lapi-dashboard" + "." + GetDomain(scope),
						"tls": []map[string]any{
							{
								"hosts":      []string{"crowdsec-lapi-dashboard" + "." + GetDomain(scope)},
								"secretName": "crowdsec-lapi-dashboard-tls",
							},
						},
					},
				},
			},
			"agent": map[string]any{
				"service": map[string]any{
					"annotations": map[string]any{
						"prometheus.io/scrape": "true",
						"prometheus.io/port":   "6060",
					},
				},
				"metrics": map[string]any{
					"enabled": true,
				},
				"acquisition": []map[string]any{
					{
						"namespace": "ingress",
						"podName":   "traefik-*",
						"program":   "traefik",
					},
				},
				// "additionalAcquisition": []map[string]any{
				// 	{
				// 		"filenames":     []string{"/var/log/containers/traefik-*_ingress_*.log"},
				// 		"force_inotify": true,
				// 		"labels": map[string]any{
				// 			"type":    "containerd",
				// 			"program": "traefik",
				// 		},
				// 	},
				// },
				"env": []map[string]any{
					{"name": "COLLECTIONS", "value": "crowdsecurity/traefik"},
					{"name": "PARSERS", "value": "crowdsecurity/cri-logs"},
					{"name": "DISABLE_PARSERS", "value": "crowdsecurity/whitelists"},
					// {"name": "DISABLE_ONLINE_API", "value": "false"},
				},
			},
			"secrets": map[string]any{
				"username": "crowdsec",
				"password": "crowdsec@123",
			},
		},
	})
	// for _, obj := range *helmResources.ApiObjects() {
	// 	if *obj.Metadata().Name() == "crowdsec-agent" && *obj.Kind() == "DaemonSet" {
	// 		obj.AddJsonPatch(
	// 			packager.JsonPatch_Test(("/spec/template/spec/containers/0/env/1/name"), "DISABLE_ONLINE_API"),
	// 			packager.JsonPatch_Replace(("/spec/template/spec/containers/0/env/1"), map[string]any{
	// 				"name":  "DISABLE_ONLINE_API",
	// 				"value": "false",
	// 			}),
	// 		)
	// 	}
	// }

	k8sapp.NewExternalSecret(chart, &k8sapp.ExternalSecretProps{
		Name: "crowdsec-keys",
		RemoteRefs: map[string]string{
			"ENROLL_KEY": "CROWDSEC_ENROLL_KEY",
		},
	})

	return chart
}
