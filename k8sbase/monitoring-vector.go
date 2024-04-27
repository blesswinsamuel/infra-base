package k8sbase

import (
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"

	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/muesli/reflow/dedent"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type VectorProps struct {
	ImageInfo    k8sapp.ImageInfo `json:"image"`
	SyslogServer struct {
		Enabled          bool   `json:"enabled"`
		VrlDecoderSource string `json:"vrlDecoderSource"`
	} `json:"syslogServer"`
}

// https://github.com/vectordotdev/helm-charts/tree/develop/charts/vector
// https://helm.vector.dev/index.yaml

func (props *VectorProps) Render(scope kubegogen.Scope) {
	syslogOpts := map[string]any{
		"decoding": map[string]any{
			"codec": "vrl",
			"vrl": map[string]any{
				"source":   strings.TrimSpace(props.SyslogServer.VrlDecoderSource),
				"timezone": "local",
			},
		},
	}

	applicationProps := &k8sapp.ApplicationProps{
		Name:      "vector",
		Kind:      "DaemonSet",
		DNSPolicy: corev1.DNSClusterFirst,
		Containers: []k8sapp.ApplicationContainer{
			{
				Name:  "vector",
				Image: props.ImageInfo,
				Args:  []string{"--config-dir", "/etc/vector/"},
				Ports: []k8sapp.ContainerPort{
					{Name: "api", Port: 8686},
					{Name: "prom-exporter", Port: 9090, PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{}},
				},
				Env: map[string]string{
					"VECTOR_LOG":  "info",
					"PROCFS_ROOT": "/host/proc",
					"SYSFS_ROOT":  "/host/sys",
				},
				ExtraEnvs: []corev1.EnvVar{
					{Name: "VECTOR_SELF_NODE_NAME", ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{FieldPath: "spec.nodeName"}}},
					{Name: "VECTOR_SELF_POD_NAME", ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{FieldPath: "metadata.name"}}},
					{Name: "VECTOR_SELF_POD_NAMESPACE", ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{FieldPath: "metadata.namespace"}}},
				},
				ExtraVolumeMounts: []corev1.VolumeMount{
					{Name: "data", MountPath: "/vector-data-dir"},
					{Name: "var-log", MountPath: "/var/log/", ReadOnly: true},
					{Name: "var-lib", MountPath: "/var/lib", ReadOnly: true},
					{Name: "procfs", MountPath: "/host/proc", ReadOnly: true},
					{Name: "sysfs", MountPath: "/host/sys", ReadOnly: true},
				},
			},
		},
		ExtraVolumes: []corev1.Volume{
			{Name: "data", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/var/lib/vector"}}},
			{Name: "var-log", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/var/log/"}}},
			{Name: "var-lib", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/var/lib/"}}},
			{Name: "procfs", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/proc"}}},
			{Name: "sysfs", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/sys"}}},
		},
		ServiceAccountName:   "vector",
		CreateServiceAccount: true,
		ConfigMaps: []k8sapp.ApplicationConfigMap{
			{
				Name:      "vector",
				MountName: "config",
				MountPath: "/etc/vector/",
				ReadOnly:  true,
				Data: map[string]string{
					"vector.yaml": infrahelpers.ToYamlString(map[string]interface{}{
						"data_dir": "/vector-data-dir",
						"api": map[string]interface{}{
							"enabled":    true,
							"address":    "0.0.0.0:8686",
							"playground": false,
						},
						"sources": infrahelpers.MergeMaps(
							infrahelpers.Ternary(props.SyslogServer.Enabled, map[string]interface{}{
								"syslog_server_tcp": infrahelpers.MergeMaps(map[string]interface{}{
									"type":    "socket",
									"address": "0.0.0.0:514",
									"mode":    "tcp",
								}, syslogOpts),
								"syslog_server_udp": infrahelpers.MergeMaps(map[string]interface{}{
									"type":       "socket",
									"address":    "0.0.0.0:514",
									"max_length": 102400,
									"mode":       "udp",
								}, syslogOpts),
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
									"inputs": []string{"syslog_server_tcp", "syslog_server_udp"},
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
							// https://playground.vrl.dev/
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
										  .level = .level || .severity || "unknown"
										} else {
											.level = "unknown"
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
									"container_name": "{{ kubernetes.container_name }}",
									"pod_name":       "{{ kubernetes.pod_name }}",
									"pod_node_name":  "{{ kubernetes.pod_node_name }}",
									"pod_namespace":  "{{ kubernetes.pod_namespace }}",
									"level":          "{{ level }}",
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
					}),
				},
			},
		},
	}
	if props.SyslogServer.Enabled {
		applicationProps.Containers[0].Ports = append(applicationProps.Containers[0].Ports, k8sapp.ContainerPort{Name: "syslog-server-t", Port: 514, Protocol: "TCP"})
		applicationProps.Containers[0].Ports = append(applicationProps.Containers[0].Ports, k8sapp.ContainerPort{Name: "syslog-server-u", Port: 514, Protocol: "UDP"})
	}
	k8sapp.NewApplication(scope, applicationProps)
	scope.AddApiObject(&rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{Name: "vector"},
		Subjects:   []rbacv1.Subject{{Kind: "ServiceAccount", Name: "vector", Namespace: scope.Namespace()}},
		RoleRef:    rbacv1.RoleRef{Kind: "ClusterRole", Name: "vector", APIGroup: "rbac.authorization.k8s.io"},
	})
	scope.AddApiObject(&rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{Name: "vector"},
		Rules: []rbacv1.PolicyRule{
			{APIGroups: []string{""}, Resources: []string{"namespaces", "nodes", "pods"}, Verbs: []string{"list", "watch"}},
		},
	})

	if props.SyslogServer.Enabled {
		scope.AddApiObject(&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name: "vector-syslog-server",
			},
			Spec: corev1.ServiceSpec{
				Type: corev1.ServiceTypeNodePort,
				Ports: []corev1.ServicePort{
					{
						Name:       "syslog-tcp",
						Port:       int32(514),
						Protocol:   "TCP",
						TargetPort: intstr.FromInt(514),
						NodePort:   int32(30514),
					},
					{
						Name:       "syslog-udp",
						Port:       int32(514),
						Protocol:   "UDP",
						TargetPort: intstr.FromInt(514),
						NodePort:   int32(30514),
					},
				},
				Selector: map[string]string{
					"app.kubernetes.io/name": "vector",
				},
			},
		})
	}
}
