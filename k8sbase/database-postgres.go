package k8sbase

import (
	"fmt"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// https://github.com/bitnami/charts/tree/main/bitnami/postgresql

type PostgresGrafanaDatasourceProps struct {
	Database string `json:"database"`
}

type PostgresProps struct {
	HelmChartInfo          k8sapp.ChartInfo  `json:"helm"`
	ImagePullSecrets       []string          `json:"imagePullSecrets"`
	ImageInfo              k8sapp.ImageInfo  `json:"image"`
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
}

func (props *PostgresProps) Chart(scope kubegogen.Construct) kubegogen.Construct {
	cprops := kubegogen.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("postgres", cprops)

	k8sapp.NewHelm(chart, "helm", &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "postgres",
		Namespace:   chart.Namespace(),
		Values: map[string]interface{}{
			"postgresqlSharedPreloadLibraries": infrahelpers.If(props.SharedPreloadLibraries != nil, strings.Join(props.SharedPreloadLibraries, ","), ""),
			"nameOverride":                     "postgres",
			"image": infrahelpers.MergeMaps(props.ImageInfo.ToMap(), map[string]any{
				"registry":    "",
				"pullPolicy":  infrahelpers.PtrIfNonEmpty(props.ImagePullPolicy),
				"pullSecrets": props.ImagePullSecrets,
			}),
			"auth": map[string]interface{}{
				"database":       props.Database,
				"username":       props.Username,
				"existingSecret": "postgres-passwords",
			},
			"metrics": map[string]interface{}{"enabled": true},
			"primary": map[string]interface{}{
				"persistence": map[string]interface{}{
					"existingClaim": infrahelpers.Ternary(props.PersistentVolumeName != "", "postgres", ""),
					"storageClass":  infrahelpers.Ternary(props.PersistentVolumeName != "", "-", ""),
				},
			},
		},
	})
	if props.PersistentVolumeName != "" {
		k8sapp.NewPersistentVolumeClaim(chart, "postgres", &k8sapp.PersistentVolumeClaim{
			Name:            "postgres",
			StorageClass:    infrahelpers.Ternary(props.PersistentVolumeName != "", "-", ""),
			RequestsStorage: "8Gi",
			VolumeName:      props.PersistentVolumeName,
		})
	}

	k8sapp.NewExternalSecret(chart, "external-secret", &k8sapp.ExternalSecretProps{
		Name: "postgres-passwords",
		RemoteRefs: map[string]string{
			"postgres-password":    "POSTGRES_ADMIN_PASSWORD",
			"password":             "POSTGRES_USER_PASSWORD",
			"replication-password": "POSTGRES_REPLICATION_PASSWORD",
		},
	})

	if props.LoadBalancer.Enabled {
		k8sapp.NewK8sObject(scope, "postgres-lb", &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name: "postgres-lb",
			},
			Spec: corev1.ServiceSpec{
				Type: corev1.ServiceTypeLoadBalancer,
				Ports: []corev1.ServicePort{
					{Name: "tcp-postgresql", Port: 5432, Protocol: "TCP", TargetPort: intstr.FromString("tcp-postgresql")},
				},
				Selector: map[string]string{
					"app.kubernetes.io/component": "primary",
					"app.kubernetes.io/instance":  "postgres",
					"app.kubernetes.io/name":      "postgres",
				},
			},
		})
	}

	for _, grafanaDatasource := range props.GrafanaDatasources {
		k8sapp.NewExternalSecret(chart, "grafana-datasource-postgres", &k8sapp.ExternalSecretProps{
			Name: fmt.Sprintf("grafana-datasource-postgres-%s", grafanaDatasource.Database),
			SecretLabels: map[string]string{
				"grafana_datasource": "1",
			},
			RemoteRefs: map[string]string{
				"password": "POSTGRES_USER_PASSWORD",
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

	return chart
}
