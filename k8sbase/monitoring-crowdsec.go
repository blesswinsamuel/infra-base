package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/helpers"
	"github.com/blesswinsamuel/infra-base/k8sbase/infraglobal"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type CrowdsecProps struct {
	Enabled       bool              `yaml:"enabled"`
	HelmChartInfo helpers.ChartInfo `yaml:"helm"`
}

// https://github.com/crowdsecurity/helm-charts/tree/main/charts/crowdsec
func NewCrowdsec(scope constructs.Construct, props CrowdsecProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: helpers.GetNamespace(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("crowdsec"), &cprops)

	helpers.NewHelmCached(chart, jsii.String("helm"), &helpers.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("crowdsec"),
		Namespace:   chart.Namespace(),
		Values: &map[string]any{
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
					"annotations": infraglobal.GetCertIssuerAnnotation(scope),
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
						"annotations": infraglobal.GetCertIssuerAnnotation(scope),
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
	// 			cdk8s.JsonPatch_Test(jsii.String("/spec/template/spec/containers/0/env/1/name"), "DISABLE_ONLINE_API"),
	// 			cdk8s.JsonPatch_Replace(jsii.String("/spec/template/spec/containers/0/env/1"), map[string]any{
	// 				"name":  "DISABLE_ONLINE_API",
	// 				"value": "false",
	// 			}),
	// 		)
	// 	}
	// }

	NewExternalSecret(chart, jsii.String("crowdsec-keys"), &ExternalSecretProps{
		Name:      jsii.String("crowdsec-keys"),
		Namespace: chart.Namespace(),
		Secrets: map[string]string{
			"ENROLL_KEY": "CROWDSEC_ENROLL_KEY",
		},
	})

	return chart
}
