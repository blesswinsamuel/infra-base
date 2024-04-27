package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type VmalertProps struct {
	ImageInfo k8sapp.ImageInfo `json:"image"`
	Sidecar   struct {
		ImageInfo k8sapp.ImageInfo `json:"image"`
	} `json:"sidecar"`
	Resources corev1.ResourceRequirements `json:"resources"`
	Ingress   struct {
		Enabled   bool   `json:"enabled"`
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
}

// https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-agent
func (props *VmalertProps) Render(scope kubegogen.Scope) {
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Name: "vmalert",
		Containers: []k8sapp.ApplicationContainer{
			{
				Name:  "vmalert",
				Image: props.ImageInfo,
				Args: []string{
					`-rule=/config/alert-rules.yaml`,
					`-datasource.url=http://victoriametrics:8428`,
					`-notifier.url=http://alertmanager:9093`,
					`-remoteRead.url=http://victoriametrics:8428`,
					`-remoteWrite.url=http://victoriametrics:8428`,
					`-envflag.enable=true`,
					`-envflag.prefix=VM_`,
					// # external.alert.source: explore?orgId=1&left=["now-1h","now","VictoriaMetrics",{"expr":""},{"mode":"Metrics"},{"ui":[true,true,true,"none"]}]
					// # external.alert.source: {{ `explore?orgId=1&left=["now-1h","now","VictoriaMetrics",{"expr":"{{$expr|quotesEscape|pathEscape}}"}]` }}
					// https://github.com/VictoriaMetrics/VictoriaMetrics/blob/8edb390e215cbffe9bb267eea8337dbf1df1c76f/deployment/docker/docker-compose.yml#L75
					`-external.alert.source=explore?orgId=1&left={"datasource":"VictoriaMetrics","queries":[{"expr":"{{$expr|quotesEscape|crlfEscape|queryEscape}}","refId":"A"}],"range":{"from":"now-1h","to":"now"}}`,
					`-external.url=https://grafana.` + GetDomain(scope),
					`-loggerFormat=json`,
					`-rule="/config/*.yaml"`,
					// # - "-external.label=env=${ENV_NAME}"
					// # - "-evaluationInterval=30s"
					// # - "-rule=/config/alert_rules.yml"
				},
				Ports: []k8sapp.ContainerPort{{Name: "http", Port: 8880, Ingress: &k8sapp.ApplicationIngress{Host: props.Ingress.SubDomain + "." + GetDomain(scope)}}},
				LivenessProbe: &corev1.Probe{
					InitialDelaySeconds: int32(5),
					PeriodSeconds:       int32(15),
					ProbeHandler:        corev1.ProbeHandler{TCPSocket: &corev1.TCPSocketAction{Port: intstr.FromString("http")}},
					TimeoutSeconds:      int32(5),
				},
				ReadinessProbe: &corev1.Probe{
					InitialDelaySeconds: int32(5),
					PeriodSeconds:       int32(15),
					ProbeHandler:        corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Port: intstr.FromString("http"), Path: "/health"}},
				},
				ExtraVolumeMounts: []corev1.VolumeMount{
					{Name: "alerts-config", MountPath: "/config"},
				},
			},
			{
				Name:  "sidecar",
				Image: props.Sidecar.ImageInfo,
				Env: map[string]string{
					"LABEL":             "alerting_rule",
					"LABEL_VALUE":       "1",
					"FOLDER":            "/config",
					"FOLDER_ANNOTATION": "alerting-rule-folder",
					"NAMESPACE":         "ALL",
					"RESOURCE":          "configmap",
					"REQ_URL":           "http://localhost:8880/-/reload",
				},
				ExtraVolumeMounts: []corev1.VolumeMount{
					{Name: "alerts-config", MountPath: "/config"},
				},
			},
		},
		ExtraVolumes: []corev1.Volume{
			// {Name: "alerts-config", VolumeSource: corev1.VolumeSource{ConfigMap: &corev1.ConfigMapVolumeSource{LocalObjectReference: corev1.LocalObjectReference{Name: "alerting-rules"}}}},
			{Name: "alerts-config", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}},
		},
		ServiceAccountName:   "vmalert",
		CreateServiceAccount: true,
	})
	scope.AddApiObject(&rbacv1.ClusterRole{
		ObjectMeta: v1.ObjectMeta{Name: "vmalert"},
		Rules: []rbacv1.PolicyRule{
			{APIGroups: []string{""}, Resources: []string{"configmaps", "secrets"}, Verbs: []string{"get", "list", "watch"}},
		},
	})
	scope.AddApiObject(&rbacv1.ClusterRoleBinding{
		ObjectMeta: v1.ObjectMeta{Name: "vmalert"},
		RoleRef:    rbacv1.RoleRef{Kind: "ClusterRole", Name: "vmalert", APIGroup: "rbac.authorization.k8s.io"},
		Subjects:   []rbacv1.Subject{{Kind: "ServiceAccount", Name: "vmalert", Namespace: scope.Namespace()}},
	})
}
