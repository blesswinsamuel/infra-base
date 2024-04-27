package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type VictoriaMetrics struct {
	HelmChartInfo k8sapp.ChartInfo            `json:"helm"`
	Resources     corev1.ResourceRequirements `json:"resources"`
	Ingress       struct {
		Enabled   bool   `json:"enabled"`
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
	RetentionPeriod        string         `json:"retentionPeriod"`
	PersistentVolume       map[string]any `json:"persistentVolume"`
	PersistentVolumeName   string         `json:"persistentVolumeName"`
	NodePortServiceEnabled bool           `json:"nodePortServiceEnabled"`
}

// https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-single
func (props *VictoriaMetrics) Render(scope kubegogen.Scope) {
	if props.PersistentVolume == nil {
		props.PersistentVolume = map[string]any{}
	}
	if props.PersistentVolumeName != "" {
		k8sapp.NewPersistentVolumeClaim(scope, &k8sapp.PersistentVolumeClaim{
			Name:            "victoriametrics",
			StorageClass:    infrahelpers.Ternary(props.PersistentVolumeName != "", "-", ""),
			RequestsStorage: "1Gi",
			VolumeName:      props.PersistentVolumeName,
		})
		props.PersistentVolume["existingClaim"] = "victoriametrics"
	}
	// TODO: remove helm dependency
	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "victoriametrics",
		Namespace:   scope.Namespace(),
		Values: map[string]any{
			"server": map[string]any{
				"retentionPeriod": props.RetentionPeriod,
				"statefulSet": map[string]any{
					"service": map[string]any{
						"annotations": map[string]any{
							"prometheus.io/scrape": "true",
							"prometheus.io/port":   "8428",
						},
					},
				},
				"ingress": map[string]any{
					"enabled":     props.Ingress.Enabled,
					"annotations": GetCertIssuerAnnotation(scope),
					"hosts": []map[string]any{
						{
							"name": props.Ingress.SubDomain + "." + GetDomain(scope),
							"path": "/",
							"port": "http",
						},
					},
					"tls": []map[string]any{
						{
							"hosts": []string{
								props.Ingress.SubDomain + "." + GetDomain(scope),
							},
							"secretName": "victoriametrics-tls",
						},
					},
					"pathType": "Prefix",
				},
				"extraArgs": map[string]any{
					"vmalert.proxyURL": `http://vmalert:8880`,
				},
				"resources":        props.Resources,
				"persistentVolume": props.PersistentVolume,
			},
		},
	})

	if props.NodePortServiceEnabled {
		scope.AddApiObject(&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: "victoriametrics-np"},
			Spec: corev1.ServiceSpec{
				Type:     corev1.ServiceTypeNodePort,
				Ports:    []corev1.ServicePort{{Name: "http", Port: 8428, TargetPort: intstr.FromInt32(8428), NodePort: 30428}},
				Selector: map[string]string{"app.kubernetes.io/name": "victoria-metrics-single"},
			},
		})
	}

	k8sapp.NewConfigMap(scope, &k8sapp.ConfigmapProps{
		Name: "grafana-datasource-victoriametrics",
		Labels: map[string]string{
			"grafana_datasource": "1",
		},
		Data: map[string]string{
			"victoriametrics.yaml": infrahelpers.ToYamlString(map[string]interface{}{
				"apiVersion": 1,
				"deleteDatasources": []map[string]interface{}{
					{
						"name":  "VictoriaMetrics",
						"orgId": 1,
					},
				},
				"datasources": []map[string]interface{}{
					{
						"name":      "VictoriaMetrics",
						"type":      "prometheus",
						"access":    "proxy",
						"orgId":     1,
						"uid":       "victoriametrics",
						"url":       "http://victoriametrics-victoria-metrics-single-server:8428",
						"isDefault": true,
						"version":   1,
						"editable":  false,
						"jsonData": map[string]interface{}{
							"timeInterval": "60s",
						},
						// # jsonData:
						// #   alertmanagerUid: alertmanager
					},
				},
			}),
		},
	})
}
