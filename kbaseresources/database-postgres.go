package kbaseresources

import (
	"fmt"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
)

// https://github.com/bitnami/charts/tree/main/bitnami/postgresql

type PostgresGrafanaDatasourceProps struct {
	Database string `json:"database"`
}

type Postgres struct {
	ImagePullSecrets string           `json:"imagePullSecrets"`
	ImageInfo        k8sapp.ImageInfo `json:"image"`
	Metrics          struct {
		ImageInfo k8sapp.ImageInfo `json:"image"`
	} `json:"metrics"`
	ImagePullPolicy        corev1.PullPolicy `json:"imagePullPolicy"`
	Database               string            `json:"database"`
	Username               string            `json:"username"`
	SharedPreloadLibraries []string          `json:"sharedPreloadLibraries"`
	LoadBalancer           struct {
		Enabled bool `json:"enabled"`
		Port    int  `json:"port"`
	} `json:"loadBalancer"`
	PersistentVolumeName string                           `json:"persistentVolumeName"`
	GrafanaDatasources   []PostgresGrafanaDatasourceProps `json:"grafana_datasources"`
	Resources            *corev1.ResourceRequirements     `json:"resources"`
	Tolerations          []corev1.Toleration              `json:"tolerations"`
	CreateService        string                           `json:"serviceName"`
}

func (props *Postgres) Render(scope kgen.Scope) {
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Name:                 scope.ID(),
		Kind:                 "StatefulSet",
		HeadlessServiceNames: []string{scope.ID() + "-hl"},
		Tolerations:          props.Tolerations,
		Containers: []k8sapp.ApplicationContainer{
			{
				Name:            "postgresql",
				Image:           props.ImageInfo,
				ImagePullPolicy: corev1.PullPolicy(props.ImagePullPolicy),
				Ports: []k8sapp.ContainerPort{
					{Name: "tcp-postgresql", Port: 5432},
				},
				Env: map[string]string{
					"POSTGRESQL_VOLUME_DIR":               "/bitnami/postgresql",
					"PGDATA":                              "/bitnami/postgresql/data",
					"POSTGRESQL_SHARED_PRELOAD_LIBRARIES": infrahelpers.If(props.SharedPreloadLibraries != nil, strings.Join(props.SharedPreloadLibraries, ","), ""),
				},
				EnvFromSecretRef: []string{scope.ID() + "-passwords"},
				LivenessProbe: &corev1.Probe{ProbeHandler: corev1.ProbeHandler{Exec: &corev1.ExecAction{Command: []string{
					"/bin/sh", "-c", fmt.Sprintf(`exec pg_isready -U "%s" -d "dbname=%s" -h 127.0.0.1 -p 5432`, props.Username, props.Database),
				}}}, FailureThreshold: 6, InitialDelaySeconds: 30, PeriodSeconds: 10, SuccessThreshold: 1, TimeoutSeconds: 5},
				ReadinessProbe: &corev1.Probe{ProbeHandler: corev1.ProbeHandler{Exec: &corev1.ExecAction{Command: []string{
					"/bin/sh", "-c", "-e", fmt.Sprintf(`exec pg_isready -U "%s" -d "dbname=%s" -h 127.0.0.1 -p 5432`, props.Username, props.Database),
				}}}, FailureThreshold: 6, InitialDelaySeconds: 5, PeriodSeconds: 10, SuccessThreshold: 1, TimeoutSeconds: 5},
				Resources: *props.Resources,
				SecurityContext: &corev1.SecurityContext{
					AllowPrivilegeEscalation: ptr.To(false),
					Capabilities:             &corev1.Capabilities{Drop: []corev1.Capability{"ALL"}},
					Privileged:               ptr.To(false),
					ReadOnlyRootFilesystem:   ptr.To(true),
					RunAsGroup:               ptr.To(int64(1001)),
					RunAsNonRoot:             ptr.To(true),
					RunAsUser:                ptr.To(int64(1001)),
					SELinuxOptions:           &corev1.SELinuxOptions{},
					SeccompProfile:           &corev1.SeccompProfile{Type: "RuntimeDefault"},
				},
				ExtraVolumeMounts: []corev1.VolumeMount{
					{Name: "empty-dir", MountPath: "/tmp", SubPath: "tmp-dir"},
					{Name: "empty-dir", MountPath: "/opt/bitnami/postgresql/conf", SubPath: "app-conf-dir"},
					{Name: "empty-dir", MountPath: "/opt/bitnami/postgresql/tmp", SubPath: "app-tmp-dir"},
					{Name: "dshm", MountPath: "/dev/shm"},
				},
			},
			{
				Name:  "metrics",
				Image: props.Metrics.ImageInfo,
				Ports: []k8sapp.ContainerPort{{Name: "http-metrics", Port: 9187, ServiceName: scope.ID() + "-metrics", PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{}}},
				Env: map[string]string{
					"DATA_SOURCE_URI":  "127.0.0.1:5432/postgres?sslmode=disable",
					"DATA_SOURCE_USER": "postgres-exporter",
				},
				ExtraEnvs: []corev1.EnvVar{
					{Name: "DATA_SOURCE_PASS", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: "POSTGRES_PASSWORD_POSTGRES_EXPORTER", LocalObjectReference: corev1.LocalObjectReference{Name: scope.ID() + "-passwords"}}}},
				},
				LivenessProbe:  &corev1.Probe{ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/", Port: intstr.FromString("http-metrics")}}, FailureThreshold: 6, InitialDelaySeconds: 5, PeriodSeconds: 10, SuccessThreshold: 1, TimeoutSeconds: 5},
				ReadinessProbe: &corev1.Probe{ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/", Port: intstr.FromString("http-metrics")}}, FailureThreshold: 6, InitialDelaySeconds: 5, PeriodSeconds: 10, SuccessThreshold: 1, TimeoutSeconds: 5},
				Resources: corev1.ResourceRequirements{
					Limits: corev1.ResourceList{
						"cpu":    resource.MustParse("300m"),
						"memory": resource.MustParse("200Mi"),
					},
					Requests: corev1.ResourceList{
						"cpu":    resource.MustParse("100m"),
						"memory": resource.MustParse("128Mi"),
					},
				},
				SecurityContext: &corev1.SecurityContext{
					AllowPrivilegeEscalation: ptr.To(false),
					Capabilities:             &corev1.Capabilities{Drop: []corev1.Capability{"ALL"}},
					Privileged:               ptr.To(false),
					ReadOnlyRootFilesystem:   ptr.To(true),
					RunAsGroup:               ptr.To(int64(1001)),
					RunAsNonRoot:             ptr.To(true),
					RunAsUser:                ptr.To(int64(1001)),
					SELinuxOptions:           &corev1.SELinuxOptions{},
					SeccompProfile:           &corev1.SeccompProfile{Type: "RuntimeDefault"},
				},
				ExtraVolumeMounts: []corev1.VolumeMount{
					{Name: "empty-dir", MountPath: "/tmp", SubPath: "tmp-dir"},
				},
			},
		},
		PodSecurityContext: &corev1.PodSecurityContext{
			FSGroup:             ptr.To[int64](1001),
			FSGroupChangePolicy: ptr.To(corev1.FSGroupChangeAlways),
		},
		ExtraVolumes: []corev1.Volume{
			{Name: "empty-dir", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}},
			{Name: "dshm", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{Medium: corev1.StorageMediumMemory}}},
		},
		StatefulSetUpdateStrategy:    v1.StatefulSetUpdateStrategy{Type: v1.RollingUpdateStatefulSetStrategyType, RollingUpdate: &v1.RollingUpdateStatefulSetStrategy{}},
		CreateServiceAccount:         true,
		AutomountServiceAccountToken: ptr.To(false),
		ServiceAccountName:           scope.ID(),
		StatefulSetServiceName:       scope.ID() + "-hl",
		Affinity: &corev1.Affinity{
			PodAntiAffinity: &corev1.PodAntiAffinity{
				PreferredDuringSchedulingIgnoredDuringExecution: []corev1.WeightedPodAffinityTerm{
					{
						PodAffinityTerm: corev1.PodAffinityTerm{
							LabelSelector: &metav1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": scope.ID()}},
							TopologyKey:   "kubernetes.io/hostname",
						},
						Weight: 1,
					},
				},
			},
		},
		PersistentVolumes: []k8sapp.ApplicationPersistentVolume{
			{Name: scope.ID(), VolumeName: props.PersistentVolumeName, RequestsStorage: "8Gi", MountName: "data", MountPath: "/bitnami/postgresql"},
		},
		ExternalSecrets: []k8sapp.ApplicationExternalSecret{
			{
				Name: scope.ID() + "-passwords",
				RemoteRefs: map[string]string{
					"POSTGRES_PASSWORD_POSTGRES_EXPORTER": "POSTGRES_PASSWORD_POSTGRES_EXPORTER",
					"POSTGRESQL_PASSWORD":                 "POSTGRES_PASSWORD_POSTGRES",
				},
			},
		},
		ImagePullSecrets: props.ImagePullSecrets,
		// TerminationGracePeriodSeconds: ptr.To(int64(300)),
		NetworkPolicy: &k8sapp.ApplicationNetworkPolicy{
			Ingress: k8sapp.NetworkPolicyIngress{
				AllowFromAllNamespaces: []intstr.IntOrString{intstr.FromString("tcp-postgresql")},
			},
		},
	})
	scope.AddApiObject(&corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        scope.ID() + "-hl",
			Annotations: map[string]string{"service.alpha.kubernetes.io/tolerate-unready-endpoints": "true"},
		},
		Spec: corev1.ServiceSpec{
			PublishNotReadyAddresses: true,
			Ports: []corev1.ServicePort{
				{Name: "tcp-postgresql", Port: 5432, Protocol: "TCP", TargetPort: intstr.FromString("tcp-postgresql")},
			},
			Selector: map[string]string{
				"app.kubernetes.io/name": scope.ID(),
			},
			ClusterIP: corev1.ClusterIPNone,
		},
	})
	if props.CreateService != "" {
		scope.AddApiObject(&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: props.CreateService},
			Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{Name: "tcp-postgresql", Port: 5432, Protocol: "TCP", TargetPort: intstr.FromString("tcp-postgresql")},
				},
				Selector: map[string]string{"app.kubernetes.io/name": scope.ID()},
			},
		})
	}

	if props.LoadBalancer.Enabled {
		scope.AddApiObject(&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name: scope.ID() + "-lb",
			},
			Spec: corev1.ServiceSpec{
				Type: corev1.ServiceTypeLoadBalancer,
				Ports: []corev1.ServicePort{
					{Name: "tcp-postgresql", Port: 5432, Protocol: "TCP", TargetPort: intstr.FromString("tcp-postgresql")},
				},
				Selector: map[string]string{"app.kubernetes.io/name": scope.ID()},
			},
		})
	}

	for _, grafanaDatasource := range props.GrafanaDatasources {
		k8sapp.NewExternalSecret(scope, &k8sapp.ExternalSecretProps{
			Name: fmt.Sprintf("grafana-datasource-postgres-%s", grafanaDatasource.Database),
			SecretLabels: map[string]string{
				"grafana_datasource": "1",
			},
			RemoteRefs: map[string]string{
				"password": "POSTGRES_PASSWORD_POSTGRES",
			},
			Template: map[string]string{
				"postgres.yaml": infrahelpers.ToYamlString(map[string]interface{}{
					"apiVersion": 1,
					"deleteDatasources": []map[string]interface{}{
						{
							"name":  fmt.Sprintf("Postgres %s", grafanaDatasource.Database),
							"orgId": 1,
						},
					},
					"datasources": []map[string]interface{}{
						{
							"name":   fmt.Sprintf("Postgres %s", grafanaDatasource.Database),
							"type":   "postgres",
							"orgId":  1,
							"uid":    fmt.Sprintf("postgres-%s", grafanaDatasource.Database),
							"url":    "postgres.database.svc.cluster.local:5432",
							"access": "proxy",
							// TODO: use readonly user, use secret
							"user": props.Username,
							"secureJsonData": map[string]any{
								"password": "{{ .password }}",
							},
							"editable": false,
							"jsonData": map[string]interface{}{
								"sslmode":          "disable",
								"connMaxLifetime":  14400,
								"database":         grafanaDatasource.Database,
								"maxIdleConns":     100,
								"maxIdleConnsAuto": true,
								"maxOpenConns":     100,
								"postgresVersion":  1400,
								"timescaledb":      false,
							},
						},
					},
				}),
			},
		})
	}
}
