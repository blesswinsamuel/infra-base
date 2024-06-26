package k8sbase

import (
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/kgen"

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
		Enabled          bool              `json:"enabled"`
		HostnameMappings map[string]string `json:"hostnameMappings"`
		Debug            bool              `json:"debug"`
	} `json:"syslogServer"`
}

// https://github.com/vectordotdev/helm-charts/tree/develop/charts/vector
// https://helm.vector.dev/index.yaml

func (props *VectorProps) Render(scope kgen.Scope) {
	sources := map[string]any{
		"kubernetes_logs": map[string]any{
			"type": "kubernetes_logs",
		},
		// // # vector_logs:
		// // #   type: internal_logs
		// "host_metrics": map[string]any{
		// 	"type": "host_metrics",
		// 	"filesystem": map[string]any{
		// 		"devices": map[string]any{
		// 			"excludes": []string{"binfmt_misc"},
		// 		},
		// 		"filesystems": map[string]any{
		// 			"excludes": []string{"binfmt_misc"},
		// 		},
		// 		"mountPoints": map[string]any{
		// 			"excludes": []string{"*/proc/sys/fs/binfmt_misc"},
		// 		},
		// 	},
		// },
		"internal_metrics": map[string]any{
			"type": "internal_metrics",
		},
	}

	// https://playground.vrl.dev/
	transforms := map[string]any{
		"kubernetes_parse_and_merge_log_message": map[string]any{
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
				  # normalize msg and message
				  if .message == null && .msg != null {
				    .message = .msg
				  }
				} else {
				  .level = "unknown"
				}
			`)),
		},
		"kubernetes_log_transform": map[string]any{
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
				del(.kubernetes.container_image_id)
				del(.file)
			`)),
		},
	}

	lokiCommonOpts := map[string]any{
		"type":     "loki",
		"endpoint": "http://loki:3100",
		"encoding": map[string]any{
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
	}
	sinks := map[string]any{
		"prom_exporter": map[string]any{
			"type": "prometheus_exporter",
			"inputs": []string{
				// "host_metrics",
				"internal_metrics",
			},
			"address": "0.0.0.0:9090",
		},
		"loki_sink_kubernetes": infrahelpers.MergeMaps(lokiCommonOpts, map[string]any{
			"inputs": []string{"kubernetes_log_transform"},
			"labels": map[string]any{
				"stream":         "kubernetes",
				"pod_node_name":  "{{ kubernetes.pod_node_name }}",
				"pod_namespace":  "{{ kubernetes.pod_namespace }}",
				"pod_name":       "{{ kubernetes.pod_name }}",
				"container_name": "{{ kubernetes.container_name }}",
				"level":          "{{ level }}",
			},
			"encoding": map[string]any{
				"codec":         "json",
				"except_fields": []string{"source_type"},
			},
		}),
	}
	if props.SyslogServer.Enabled {
		// Sources
		hostnameMappingStr := ""
		for _, k := range infrahelpers.MapKeysSorted(props.SyslogServer.HostnameMappings) {
			v := props.SyslogServer.HostnameMappings[k]
			hostnameMappingStr += `    .hostname = replace(.hostname, "` + k + `", "` + v + `") ?? .hostname` + "\n"
		}
		syslogOpts := map[string]any{
			"decoding": map[string]any{
				"codec": "vrl",
				"vrl": map[string]any{
					// https://vector.dev/docs/reference/configuration/sources/syslog/#examples-syslog-event
					// https://vector.dev/docs/reference/vrl/functions/#parse_syslog
					"source": strings.TrimSpace(`
raw_message = .message
parsed_message, err = parse_syslog(.message)
if err != null {
  .message = string!(.message)
  # <XXXMODELNAMEXXX> - <13> syslog: Do not receive BRAS LCP-echo-request. Failure count: 1/3
  # <XXXMODELNAMEXXX> - <4> kernel: 1714246742.659541: [mapd][wapp_wait_recv_parse_wapp_resp][1326][wapp_wait_recv_parse_wapp_resp]wait for event timeout
  parsed_message, err = parse_regex(.message, r'<(?P<hostname>\w+)> - <(?P<priority>\d+)> (?P<appname>\w+): (?P<message>.*)')
  if err != null {
    log("Failed to parse message: " + err, level: "error")
    .hostname = "unknown"
    .facility = "unknown"
    .severity = "unknown"
    .appname = "unknown"
  } else {
    # https://gist.github.com/marvin/1017480?permalink_comment_id=584813#gistcomment-584813
    parsed_message.priority = to_int!(parsed_message.priority)
    . = parsed_message
` + hostnameMappingStr + `
    .facility = to_syslog_facility(to_int(.priority / 8)) ?? "invalid"
    .severity = to_syslog_level(.priority - (to_int(.priority / 8) * 8)) ?? "invalid"
  }
} else {
  . = parsed_message
  # reset timestamp to current time because the timezone is wonky sometimes - this results in "400 Bad Request" from loki
  .timestamp = format_timestamp!(now(), format: "%+")
}
.raw_message = raw_message
# .level = .severity
# del(.severity)
`),
					"timezone": "local",
				},
			},
		}
		sources["syslog_server_tcp"] = infrahelpers.MergeMaps(map[string]any{
			"type":    "socket",
			"address": "0.0.0.0:514",
			"mode":    "tcp",
		}, syslogOpts)
		sources["syslog_server_udp"] = infrahelpers.MergeMaps(map[string]any{
			"type":       "socket",
			"address":    "0.0.0.0:514",
			"max_length": 102400,
			"mode":       "udp",
		}, syslogOpts)

		transforms["syslog_transform"] = map[string]any{
			"type":   "remap",
			"inputs": []string{"syslog_server_tcp", "syslog_server_udp"},
			"source": ``,
		}

		// Sinks
		sinks["loki_sink_syslog"] = infrahelpers.MergeMaps(lokiCommonOpts, map[string]any{
			"inputs": []string{"syslog_transform"},
			"labels": map[string]any{
				"stream": "syslog",
				// "host":     "{{ host }}",
				"hostname": "{{ hostname }}",
				"program":  "{{ appname }}",
				"facility": "{{ facility }}",
				"level":    "{{ severity }}",
			},
			"encoding": map[string]any{
				"codec":         "json",
				"except_fields": []string{"raw_message", "source_type"},
			},
		})
		// props.SyslogServer.Debug = true
		if props.SyslogServer.Debug {
			sinks["console_debug_syslog"] = map[string]any{
				"type":   "console",
				"inputs": []string{"syslog_transform"},
				"encoding": map[string]any{
					"codec": "json",
				},
			}
			// {"appname":"sshd","facility":"daemon","host":"10.42.0.1","hostname":"homelab-asdf","message":"portmapper: UPnP discovered root \"http://asdf:1900/gatedesc.xml\" does not match gateway IP 192.168.1.1; repointing at gateway which is assumed to be floating","port":40549,"procid":999,"raw_message":"<30>Apr 28 14:54:12 homelab-asdf tailscaled[999]: portmapper: UPnP discovered root \"http://192.168.0.254:1900/gatedesc.xml\" does not match gateway IP 192.168.1.1; repointing at gateway which is assumed to be floating","severity":"info","source_type":"socket","timestamp":"2024-04-28T09:24:12.012889488+00:00"}
			// {"appname":"kernel","facility":"kern","host":"10.42.0.1","hostname":"isp-asdf","message":"1714296252.091443: [mapd][wapp_wait_recv_parse_wapp_resp][1326][wapp_wait_recv_parse_wapp_resp]wait for event timeout","port":1223,"priority":4,"raw_message":"<XXXMODELNAMEXXX> - <4> kernel: 1714296251.091441: [mapd][wapp_wait_recv_parse_wapp_resp][1326][wapp_wait_recv_parse_wapp_resp]wait for event timeout","severity":"warning","source_type":"socket","timestamp":"2024-04-28T09:24:11.144703156Z"}
		}
	}

	config := map[string]any{
		"data_dir": "/vector-data-dir",
		"api": map[string]any{
			"enabled":    true,
			"address":    "0.0.0.0:8686",
			"playground": false,
		},
		"sources":    sources,
		"transforms": transforms,
		"sinks":      sinks,
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
					"TZ":          "UTC",
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
					"vector.yaml": infrahelpers.ToYamlString(config),
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
