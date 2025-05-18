package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
)

type Redis struct {
	ImageInfo k8sapp.ImageInfo `json:"image"`
	Metrics   struct {
		ImageInfo k8sapp.ImageInfo `json:"image"`
	} `json:"metrics"`
	Resources            corev1.ResourceRequirements `json:"resources"`
	PersistentVolumeName string                      `json:"persistentVolumeName"`
	Tolerations          []corev1.Toleration         `json:"tolerations"`
}

// https://github.com/bitnami/charts/tree/main/bitnami/redis

func (props *Redis) Render(scope kgen.Scope) {
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Kind:                         k8sapp.ApplicationKindStatefulSet,
		Name:                         "redis-master",
		HeadlessServiceNames:         []string{"redis-headless"},
		StatefulSetServiceName:       "redis-headless",
		ServiceAccountName:           "redis-master",
		CreateServiceAccount:         true,
		AutomountServiceAccountToken: ptr.To(false),
		Tolerations:                  props.Tolerations,
		Affinity: &corev1.Affinity{
			PodAntiAffinity: &corev1.PodAntiAffinity{
				PreferredDuringSchedulingIgnoredDuringExecution: []corev1.WeightedPodAffinityTerm{
					{
						PodAffinityTerm: corev1.PodAffinityTerm{
							LabelSelector: &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "redis"}},
							TopologyKey:   "kubernetes.io/hostname",
						},
						Weight: 1,
					},
				},
			},
		},
		Containers: []k8sapp.ApplicationContainer{
			{
				Name:  "redis",
				Image: props.ImageInfo,
				Ports: []k8sapp.ContainerPort{{Name: "tcp-redis", Port: 6379, ServiceName: "redis-master"}},
				Env: map[string]string{
					"REDIS_REPLICATION_MODE": "master",
					"BITNAMI_DEBUG":          "false",
					"ALLOW_EMPTY_PASSWORD":   "yes",
					// "REDIS_DISABLE_COMMANDS": "FLUSHDB,FLUSHALL",
				},
				Resources: props.Resources,
				ExtraVolumeMounts: []corev1.VolumeMount{
					{Name: "empty-dir", MountPath: "/opt/bitnami/redis/tmp", SubPath: "tmp-dir"},
					{Name: "empty-dir", MountPath: "/opt/bitnami/redis/logs", SubPath: "log-dir"},
					{Name: "empty-dir", MountPath: "/opt/bitnami/redis/etc", SubPath: "app-conf-dir"},
					// {Name: "empty-dir", MountPath: "/tmp", SubPath: "tmp-dir"},
				},
			},
			{
				Name:  "metrics",
				Image: props.Metrics.ImageInfo,
				Ports: []k8sapp.ContainerPort{{Name: "http-metrics", Port: 9121, ServiceName: "redis-metrics", PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{}}},
				Env: map[string]string{
					"REDIS_ALIAS":                       "redis",
					"REDIS_EXPORTER_WEB_LISTEN_ADDRESS": ":9121",
				},
				LivenessProbe:  &corev1.Probe{FailureThreshold: 5, InitialDelaySeconds: 10, PeriodSeconds: 10, SuccessThreshold: 1, ProbeHandler: corev1.ProbeHandler{TCPSocket: &corev1.TCPSocketAction{Port: intstr.FromString("http-metrics")}}, TimeoutSeconds: 5},
				ReadinessProbe: &corev1.Probe{FailureThreshold: 3, InitialDelaySeconds: 5, PeriodSeconds: 10, SuccessThreshold: 1, ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/", Port: intstr.FromString("http-metrics")}}, TimeoutSeconds: 1},
				ExtraVolumeMounts: []corev1.VolumeMount{
					{Name: "empty-dir", MountPath: "/tmp", SubPath: "app-tmp-dir"},
				},
			},
		},
		Security: &k8sapp.ApplicationSecurity{User: 1001, Group: 1001, FSGroup: 1001},
		ExtraVolumes: []corev1.Volume{
			{Name: "empty-dir", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}},
		},
		PersistentVolumes: []k8sapp.ApplicationPersistentVolume{
			{Name: "redis", VolumeName: props.PersistentVolumeName, RequestsStorage: "1Gi", MountToContainers: []string{"redis"}, MountName: "data", MountPath: "/bitnami/redis/data"},
		},
		NetworkPolicy: &k8sapp.ApplicationNetworkPolicy{
			Ingress: k8sapp.NetworkPolicyIngress{
				AllowFromAllNamespaces: []intstr.IntOrString{intstr.FromString("tcp-redis")},
			},
		},
	})
	scope.AddApiObject(&corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "redis-headless",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{Name: "tcp-redis", Port: 6379, Protocol: "TCP", TargetPort: intstr.FromString("redis")},
			},
			Selector: map[string]string{
				"app.kubernetes.io/name": "redis",
			},
			ClusterIP: corev1.ClusterIPNone,
		},
	})
}
