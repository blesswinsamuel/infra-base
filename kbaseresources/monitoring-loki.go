package kbaseresources

import (
	"log"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type LokiProps struct {
	ImageInfo            k8sapp.ImageInfo `json:"image"`
	Storage              string           `json:"storage"`
	PersistentVolumeName string           `json:"persistentVolumeName"`
	S3                   struct {
		Endpoint        string `json:"endpoint"`
		SecretAccessKey string `json:"secret_access_key"`
		AccessKeyID     string `json:"access_key_id"`
	} `json:"s3"`
}

// https://github.com/grafana/loki/tree/main/production/helm/loki
func (props *LokiProps) Render(scope kgen.Scope) {
	var vcts []k8sapp.ApplicationPersistentVolume
	vols := []corev1.Volume{}
	volMnts := []corev1.VolumeMount{}
	if props.PersistentVolumeName != "" {
		// k8sapp.NewPersistentVolumeClaim(scope, &k8sapp.PersistentVolumeClaim{
		// 	Name:            "loki",
		// 	StorageClass:    "-",
		// 	RequestsStorage: "10Gi",
		// 	VolumeName:      props.PersistentVolumeName,
		// })
		// vols = []corev1.Volume{{Name: "storage", VolumeSource: corev1.VolumeSource{PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{ClaimName: "loki"}}}}
		// volMnts = []corev1.VolumeMount{{Name: "storage", MountPath: "/var/loki"}}
		vcts = []k8sapp.ApplicationPersistentVolume{{Name: "storage", StorageClass: "-", VolumeName: props.PersistentVolumeName, RequestsStorage: "10Gi", MountName: "storage", MountPath: "/var/loki"}}
	} else {
		vcts = []k8sapp.ApplicationPersistentVolume{{Name: "storage", RequestsStorage: "16Gi", MountName: "storage", MountPath: "/var/loki"}}
	}
	lokiConfig := map[string]any{
		"server": map[string]any{
			"grpc_listen_port": 9095,
			"http_listen_port": 3100,
			"log_format":       "json",
		},
		"analytics": map[string]any{
			"reporting_enabled": false,
		},
		"auth_enabled": false,
		"common": map[string]any{
			"compactor_address":  "http://loki:3100",
			"path_prefix":        "/var/loki",
			"replication_factor": 1,
			"ring": map[string]any{
				"kvstore": map[string]any{"store": "inmemory"},
			},
		},
		"compactor": map[string]any{
			"retention_enabled":    true,
			"delete_request_store": "filesystem", // ?
		},
		"frontend": map[string]any{
			"scheduler_address": "",
			// "encoding":          "protobuf",
		},
		"frontend_worker": map[string]any{
			"scheduler_address": "",
		},
		"index_gateway": map[string]any{
			"mode": "ring",
		},
		"limits_config": map[string]any{
			"max_cache_freshness_per_query": "10m",
			"reject_old_samples":            true,
			"reject_old_samples_max_age":    "168h",
			"split_queries_by_interval":     "15m",
		},
		"memberlist": map[string]any{
			"join_members": []string{
				"loki-memberlist",
			},
		},
		"query_range": map[string]any{
			"align_queries_with_step": true,
			"results_cache": map[string]any{
				"cache": map[string]any{
					"embedded_cache": map[string]any{"enabled": true, "max_size_mb": 100},
				},
			},
		},
		"ruler": map[string]any{
			"alertmanager_url": "http://alertmanager:9093",
			"storage": map[string]any{
				"type": "local",
			},
		},
		"runtime_config": map[string]any{
			"file": "/etc/loki/runtime-config/runtime-config.yaml",
		},
		"schema_config": map[string]any{
			"configs": []map[string]any{
				{
					"from": "2022-01-11",
					"index": map[string]any{
						"period": "24h",
						"prefix": "loki_index_",
					},
					"object_store": "filesystem",
					"schema":       "v12",
					"store":        "boltdb-shipper",
				},
				{
					"from": "2024-05-13",
					"index": map[string]any{
						"period": "24h",
						"prefix": "index_",
					},
					"object_store": "filesystem",
					"schema":       "v13",
					"store":        "tsdb",
				},
			},
		},
		"storage_config": map[string]any{
			"hedging": map[string]any{
				"at":             "250ms",
				"max_per_second": 20,
				"up_to":          3,
			},
			"tsdb_shipper": map[string]any{
				"active_index_directory": "/var/loki/tsdb-index",
				"cache_location":         "/var/loki/tsdb-cache",
			},
		},
		"tracing": map[string]any{
			"enabled": false,
		},
	}
	switch props.Storage {
	case "local":
		lokiConfig["common"].(map[string]any)["storage"] = map[string]any{
			"filesystem": map[string]any{
				"chunks_directory": "/var/loki/chunks",
				"rules_directory":  "/var/loki/rules",
			},
		}
	case "s3":
		log.Panic("s3 storage not implemented")
		// lokiConfig["storage"] = map[string]any{
		// 	"s3": map[string]any{
		// 		"bucketNames": map[string]string{
		// 			"chunks": "loki-chunks",
		// 			"ruler":  "loki-ruler",
		// 			"admin":  "loki-admin", // never used
		// 		},
		// 		"endpoint":         props.S3.Endpoint,
		// 		"secretAccessKey":  props.S3.SecretAccessKey,
		// 		"accessKeyId":      props.S3.AccessKeyID,
		// 		"s3ForcePathStyle": true,
		// 		// insecure: true,
		// 	},
		// }
	}
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Kind:                         "StatefulSet",
		Name:                         "loki",
		ServiceAccountName:           "loki",
		CreateServiceAccount:         true,
		AutomountServiceAccountToken: true,
		CreateHeadlessService:        true,
		EnableServiceLinks:           infrahelpers.Ptr(true),
		StatefulSetServiceName:       "loki-headless",
		StatefulSetUpdateStrategy:    v1.StatefulSetUpdateStrategy{RollingUpdate: &v1.RollingUpdateStatefulSetStrategy{Partition: infrahelpers.Ptr(int32(0))}},
		PodSecurityContext: &corev1.PodSecurityContext{
			FSGroup:      infrahelpers.Ptr(int64(10001)),
			RunAsGroup:   infrahelpers.Ptr(int64(10001)),
			RunAsUser:    infrahelpers.Ptr(int64(10001)),
			RunAsNonRoot: infrahelpers.Ptr(true),
		},
		Containers: []k8sapp.ApplicationContainer{{
			Name:  "loki",
			Image: props.ImageInfo,
			Ports: []k8sapp.ContainerPort{
				{Name: "http-metrics", Port: 3100}, // loki-headless should have only this
				{Name: "grpc", Port: 9095},
			},
			Args: []string{
				"-config.file=/etc/loki/config/config.yaml",
				"-target=all",
				"-validation.allow-structured-metadata=false", // TODO temporary
			},
			ReadinessProbe: &corev1.Probe{InitialDelaySeconds: 30, TimeoutSeconds: 1, ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Port: intstr.FromString("http-metrics"), Path: "/ready"}}},
			ExtraVolumeMounts: infrahelpers.MergeLists(volMnts, []corev1.VolumeMount{
				{Name: "tmp", MountPath: "/tmp"},
			}),
			SecurityContext: &corev1.SecurityContext{Capabilities: &corev1.Capabilities{Drop: []corev1.Capability{"ALL"}}, AllowPrivilegeEscalation: infrahelpers.Ptr(false), ReadOnlyRootFilesystem: infrahelpers.Ptr(true)},
		}},
		ExtraVolumes: infrahelpers.MergeLists(vols, []corev1.Volume{
			{Name: "tmp", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}},
		}),
		StatefulSetVolumeClaimTemplates: vcts,
		ConfigMaps: []k8sapp.ApplicationConfigMap{
			{
				Name: "loki",
				Data: map[string]string{
					"config.yaml": infrahelpers.ToYamlString(lokiConfig),
				},
				MountName: "config",
				MountPath: "/etc/loki/config",
				ReadOnly:  true,
			},
			{
				Name: "loki-runtime",
				Data: map[string]string{
					"runtime-config.yaml": infrahelpers.ToYamlString(map[string]any{}),
				},
				MountName: "runtime-config",
				MountPath: "/etc/loki/runtime-config",
				ReadOnly:  true,
			},
		},
	})

	scope.AddApiObject(&corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: "loki-memberlist"},
		Spec: corev1.ServiceSpec{
			Ports:                    []corev1.ServicePort{{Name: "http-memberlist", Port: 7946, TargetPort: intstr.FromInt(7946)}},
			Selector:                 map[string]string{"app.kubernetes.io/name": "loki"},
			PublishNotReadyAddresses: true,
		},
	})

	k8sapp.NewConfigMap(scope, &k8sapp.ConfigmapProps{
		Name: "grafana-datasource-loki",
		Labels: map[string]string{
			"grafana_datasource": "1",
		},
		Data: map[string]string{
			"loki.yaml": infrahelpers.ToYamlString(map[string]interface{}{
				"apiVersion": 1,
				"deleteDatasources": []map[string]interface{}{
					{
						"name":  "Loki",
						"orgId": 1,
					},
				},
				"datasources": []map[string]interface{}{
					{
						"name":   "Loki",
						"type":   "loki",
						"access": "proxy",
						"orgId":  1,
						"uid":    "loki",
						"url":    "http://loki:3100",
						"jsonData": map[string]interface{}{
							"maxLines": 1000,
							// # alertmanagerUid: alertmanager
						},
					},
				},
			}),
		},
	})
}
