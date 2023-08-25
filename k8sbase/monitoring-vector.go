package k8sbase

import (
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/packager"

	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/muesli/reflow/dedent"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type VectorProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
	SyslogServer  struct {
		Enabled bool `json:"enabled"`
	} `json:"syslogServer"`
	Ingress struct {
		Enabled   bool   `json:"enabled"`
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
}

// https://github.com/vectordotdev/helm-charts/tree/develop/charts/vector
// https://helm.vector.dev/index.yaml

func (props *VectorProps) Chart(scope packager.Construct) packager.Construct {
	cprops := packager.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("vector", cprops)

	k8sapp.NewHelm(chart, "helm", &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "vector",
		Namespace:   chart.Namespace(),
		Values: map[string]interface{}{
			"role": "Agent",
			// Prometheus scrape is disabled because it's creating duplicate metrics. Also, there is a peer_addr which blows up the cardinality
			"podAnnotations": map[string]interface{}{
				"prometheus.io/port":   "9090",
				"prometheus.io/scrape": "true",
			},
			"ingress": infrahelpers.Ternary(props.Ingress.Enabled, map[string]interface{}{
				"enabled":     true,
				"annotations": GetCertIssuerAnnotation(scope),
				"hosts": []map[string]interface{}{
					{
						"host": props.Ingress.SubDomain + "." + GetDomain(scope),
						"paths": []map[string]interface{}{
							{
								"path":     "/",
								"pathType": "ImplementationSpecific",
								"port":     map[string]interface{}{"name": "api"},
							},
						},
					},
				},
				"tls": []map[string]interface{}{
					{
						"hosts": []string{
							props.Ingress.SubDomain + "." + GetDomain(scope),
						},
						"secretName": "vector-tls",
					},
				},
			}, map[string]any{}),
			"customConfig": map[string]interface{}{
				"data_dir": "/vector-data-dir",
				"api": map[string]interface{}{
					"enabled":    true,
					"address":    "0.0.0.0:8686",
					"playground": false,
				},
				"sources": infrahelpers.MergeMaps(
					infrahelpers.Ternary(props.SyslogServer.Enabled, map[string]interface{}{
						"syslog_server": map[string]interface{}{
							"type":       "syslog",
							"address":    "0.0.0.0:514",
							"max_length": 102400,
							"mode":       "tcp",
							"path":       "/syslog-socket",
						},
					}, nil),
					map[string]interface{}{
						"kubernetes_logs": map[string]interface{}{
							"type": "kubernetes_logs",
						},
						// // # vector_logs:
						// // #   type: internal_logs
						// "host_metrics": map[string]interface{}{
						// 	"type": "host_metrics",
						// 	"filesystem": map[string]interface{}{
						// 		"devices": map[string]interface{}{
						// 			"excludes": []string{"binfmt_misc"},
						// 		},
						// 		"filesystems": map[string]interface{}{
						// 			"excludes": []string{"binfmt_misc"},
						// 		},
						// 		"mountPoints": map[string]interface{}{
						// 			"excludes": []string{"*/proc/sys/fs/binfmt_misc"},
						// 		},
						// 	},
						// },
						"internal_metrics": map[string]interface{}{
							"type": "internal_metrics",
						},
					},
				),
				"transforms": infrahelpers.MergeMaps(
					infrahelpers.Ternary(props.SyslogServer.Enabled, map[string]interface{}{
						"syslog_transform": map[string]interface{}{
							"type":   "remap",
							"inputs": []string{"syslog_server"},
							"source": strings.TrimSpace(dedent.String(`
								.kubernetes = {}
								.kubernetes.container_name = .appname
								.kubernetes.pod_name = .appname
								.kubernetes.pod_node_name = .host
								.kubernetes.pod_namespace = "syslog"
								.level = .severity
							`)),
						},
					}, nil),
					map[string]interface{}{
						"kubernetes_parse_and_merge_log_message": map[string]interface{}{
							"type":   "remap",
							"inputs": []string{"kubernetes_logs"},
							"source": strings.TrimSpace(dedent.String(`
								parsed_message, err = parse_json(.message) # ?? parse_common_log(.message) ?? parse_logfmt(.message) # ?? parse_syslog(.message)
								if err == null {
								  del(.message)
								  ., err = merge(., parsed_message)
								  if err != null {
								    log("Failed to merge message into log: " + err, level: "error")
								  }
								}
							`)),
						},
						"kubernetes_log_transform": map[string]interface{}{
							"type":   "remap",
							"inputs": []string{"kubernetes_parse_and_merge_log_message"},
							"source": strings.TrimSpace(dedent.String(`
								# .@timestamp = del(.timestamp)
								del(.kubernetes.pod_labels)
								del(.kubernetes.pod_annotations)
								del(.kubernetes.node_labels)
								del(.kubernetes.namespace_labels)
								del(.kubernetes.container_id)
								del(.kubernetes.pod_uid)
								del(.kubernetes.pod_ip)
								del(.kubernetes.pod_ips)
								del(.file)
							`)),
						},
					},
				),
				"sinks": map[string]interface{}{
					"prom_exporter": map[string]interface{}{
						"type": "prometheus_exporter",
						"inputs": []string{
							// "host_metrics",
							"internal_metrics",
						},
						"address": "0.0.0.0:9090",
					},
					"loki_sink": map[string]interface{}{
						"type": "loki",
						"inputs": infrahelpers.MergeLists(
							[]string{"kubernetes_log_transform"},
							infrahelpers.Ternary(props.SyslogServer.Enabled, []string{"syslog_transform"}, nil),
						),
						"endpoint": "http://loki:3100",
						"labels": map[string]interface{}{
							"container_name": "{{`{{ kubernetes.container_name }}`}}",
							"pod_name":       "{{`{{ kubernetes.pod_name }}`}}",
							"pod_node_name":  "{{`{{ kubernetes.pod_node_name }}`}}",
							"pod_namespace":  "{{`{{ kubernetes.pod_namespace }}`}}",
							"level":          "{{`{{ level }}`}}",
						},
						"encoding": map[string]interface{}{
							"timestamp_format": "rfc3339",
							"codec":            "json",
						},
						"out_of_order_action": "accept",
						// # debug_sink:
						// #   type: console
						// #   inputs:
						// #     - syslog_server
						// #   target: stdout
						// #   encoding:
						// #     codec: json
						//   # healthcheck:
						//   #   enabled: true
					},
				},
			},
		},
	})

	if props.SyslogServer.Enabled {
		k8sapp.NewK8sObject(chart, "syslog-service", &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name: "vector-syslog-server",
			},
			Spec: corev1.ServiceSpec{
				Type: corev1.ServiceType("NodePort"),
				Ports: []corev1.ServicePort{
					{
						Name:       "syslog-server",
						Port:       int32(514),
						Protocol:   "TCP",
						TargetPort: intstr.FromInt(514),
						NodePort:   int32(30514),
					},
				},
				Selector: map[string]string{
					"app.kubernetes.io/component": "Agent",
					"app.kubernetes.io/instance":  "vector",
					"app.kubernetes.io/name":      "vector",
				},
			},
		})
	}

	return chart
}
