package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type VictoriaMetrics struct {
	ImageInfo k8sapp.ImageInfo            `json:"image"`
	Resources corev1.ResourceRequirements `json:"resources"`
	Ingress   struct {
		Enabled   bool   `json:"enabled"`
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
	RetentionPeriod        string              `json:"retentionPeriod"`
	PersistentVolumeName   string              `json:"persistentVolumeName"`
	NodePortServiceEnabled bool                `json:"nodePortServiceEnabled"`
	Tolerations            []corev1.Toleration `json:"tolerations"`
}

// https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-single
func (props *VictoriaMetrics) Render(scope kgen.Scope) {
	vcts := []k8sapp.ApplicationPersistentVolume{}
	pvs := []k8sapp.ApplicationPersistentVolume{}
	if props.PersistentVolumeName != "" {
		pvs = []k8sapp.ApplicationPersistentVolume{{Name: "victoriametrics", VolumeName: props.PersistentVolumeName, MountName: "server-volume", MountPath: "/storage"}}
	} else {
		vcts = []k8sapp.ApplicationPersistentVolume{{Name: "server-volume", RequestsStorage: "16Gi", MountName: "server-volume", MountPath: "/storage"}}
	}
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		DeploymentUpdateStrategy: appsv1.DeploymentStrategy{Type: appsv1.RecreateDeploymentStrategyType},
		Name:                     "victoriametrics",
		Containers: []k8sapp.ApplicationContainer{{
			Name:  "victoriametrics",
			Image: props.ImageInfo,
			Ports: []k8sapp.ContainerPort{
				{Name: "http", Port: 8428, Ingress: &k8sapp.ApplicationIngress{Host: props.Ingress.SubDomain + "." + k8sapp.GetDomain(scope)}, PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{}},
			},
			Args: []string{
				"--retentionPeriod=" + props.RetentionPeriod,
				"--storageDataPath=/storage",
				"--envflag.enable=true",
				"--envflag.prefix=VM_",
				"--loggerFormat=json",
				"--vmalert.proxyURL=http://vmalert:8880",
			},
			ReadinessProbe: &corev1.Probe{FailureThreshold: 10, InitialDelaySeconds: 30, PeriodSeconds: 30, TimeoutSeconds: 5, ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Port: intstr.FromString("http"), Path: "/health"}}},
			LivenessProbe:  &corev1.Probe{FailureThreshold: 10, InitialDelaySeconds: 30, PeriodSeconds: 30, TimeoutSeconds: 5, ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Port: intstr.FromString("http"), Path: "/health"}}},
			Resources:      props.Resources,
		}},
		Security:                        &k8sapp.ApplicationSecurity{User: 65534, Group: 65534, FSGroup: 65534},
		PersistentVolumes:               pvs,
		StatefulSetVolumeClaimTemplates: vcts,
		Tolerations:                     props.Tolerations,
		Homepage: &k8sapp.ApplicationHomepage{
			Name:        "VictoriaMetrics",
			Description: "Metrics storage",
			SiteMonitor: "http://victoriametrics." + scope.Namespace() + ".svc.cluster.local:8428/health",
			Group:       "Infra",
			Icon:        "si-victoriametrics",
		},
		NetworkPolicy: &k8sapp.ApplicationNetworkPolicy{
			Ingress: k8sapp.NetworkPolicyIngress{
				AllowFromIPs: map[string][]intstr.IntOrString{"0.0.0.0/0": {intstr.FromString("http")}}, // TODO: move to vmagent
				AllowFromAppRefs: map[string][]intstr.IntOrString{
					"grafana": {intstr.FromString("http")},
					"vmalert": {intstr.FromString("http")},
				},
			},
			Egress: k8sapp.NetworkPolicyEgress{
				AllowToAppRefs: []string{"vmalert"},
			},
		},
	})

	if props.NodePortServiceEnabled {
		scope.AddApiObject(&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: "victoriametrics-np"},
			Spec: corev1.ServiceSpec{
				Type:     corev1.ServiceTypeNodePort,
				Ports:    []corev1.ServicePort{{Name: "http", Port: 8428, TargetPort: intstr.FromInt32(8428), NodePort: 30428}},
				Selector: map[string]string{"app.kubernetes.io/name": "victoriametrics"},
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
						"url":       "http://victoriametrics:8428",
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
