package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type PostgresProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
	ImageInfo     k8sapp.ImageInfo `json:"image"`
	Database      string           `json:"database"`
	Username      string           `json:"username"`
	LoadBalancer  struct {
		Enabled bool `json:"enabled"`
		Port    int  `json:"port"`
	} `json:"loadBalancer"`
}

func (props *PostgresProps) Chart(scope packager.Construct) packager.Construct {
	cprops := packager.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("postgres", cprops)

	k8sapp.NewHelm(chart, "helm", &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "postgres",
		Namespace:   chart.Namespace(),
		Values: map[string]interface{}{
			"nameOverride": "postgres",
			"image":        props.ImageInfo.ToMap(),
			"auth": map[string]interface{}{
				"database":       props.Database,
				"username":       props.Username,
				"existingSecret": "postgres-passwords",
			},
			"metrics": map[string]interface{}{"enabled": true},
		},
	})

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

	return chart
}
