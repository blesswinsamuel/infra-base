package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type LokiProps struct {
	Enabled       bool             `json:"enabled"`
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
	Storage       string           `json:"storage"`
	S3            struct {
		Endpoint        string `json:"endpoint"`
		SecretAccessKey string `json:"secret_access_key"`
		AccessKeyID     string `json:"access_key_id"`
	} `json:"s3"`
}

// https://github.com/grafana/loki/tree/main/production/helm/loki
func NewLoki(scope constructs.Construct, props LokiProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: k8sapp.GetNamespaceContextPtr(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("loki"), &cprops)

	k8sapp.NewHelm(chart, jsii.String("helm"), &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("loki"),
		Namespace:   chart.Namespace(),
		Values: &map[string]any{
			"singleBinary": map[string]any{
				"replicas": 1,
			},
			"monitoring": map[string]any{
				"dashboards": map[string]any{
					"enabled": false,
				},
				"serviceMonitor": map[string]any{
					"enabled": false,
					"metricsInstance": map[string]any{
						"enabled": false,
					},
				},
				"alerts": map[string]any{
					"enabled": false,
				},
				"rules": map[string]any{
					"enabled":  false,
					"alerting": false,
				},
				"selfMonitoring": map[string]any{
					"enabled": false,
					"grafanaAgent": map[string]any{
						"installOperator": false,
					},
				},
				"lokiCanary": map[string]any{
					"enabled": false,
				},
			},
			"test": map[string]any{
				"enabled": false,
			},
			"gateway": map[string]any{
				"enabled": false,
			},
			"memberlist": map[string]any{
				"service": map[string]any{
					// https://github.com/grafana/loki/issues/7907#issuecomment-1445336799
					"publishNotReadyAddresses": true,
				},
			},
			"loki": map[string]any{
				"auth_enabled": false,
				"commonConfig": map[string]any{
					"replication_factor": 1,
				},
				"compactor": map[string]any{
					"retention_enabled": true,
				},
				"rulerConfig": map[string]any{
					"alertmanager_url": "http://alertmanager:9093",
				},
				"storage": infrahelpers.MergeMaps(
					infrahelpers.Ternary(
						props.Storage == "local",
						map[string]any{
							"type": "filesystem",
						},
						nil,
					),
					infrahelpers.Ternary(
						props.Storage == "s3",
						map[string]any{
							"type": "s3",
							"bucketNames": map[string]string{
								"chunks": "loki-chunks",
								"ruler":  "loki-ruler",
								"admin":  "loki-admin", // never used
							},
							"s3": map[string]any{
								"endpoint":         props.S3.Endpoint,
								"secretAccessKey":  props.S3.SecretAccessKey,
								"accessKeyId":      props.S3.AccessKeyID,
								"s3ForcePathStyle": true,
								// insecure: true,
							},
						},
						nil,
					),
				),
			},
		},
	})

	return chart
}
