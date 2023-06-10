package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type VmalertProps struct {
	Enabled   bool             `yaml:"enabled"`
	ImageInfo k8sapp.ImageInfo `yaml:"image"`
	Resources *Resources       `yaml:"resources"`
	Ingress   struct {
		Enabled   bool   `yaml:"enabled"`
		SubDomain string `yaml:"subDomain"`
	} `yaml:"ingress"`
}

// https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-agent
func NewVmalert(scope constructs.Construct, props VmalertProps) constructs.Construct {
	if !props.Enabled {
		return nil
	}

	return k8sapp.NewApplicationChart(scope, "vmalert", &k8sapp.ApplicationProps{
		Name: "vmalert",
		Containers: []k8sapp.ApplicationContainer{{
			Name:  "vmalert",
			Image: props.ImageInfo,
			Args: []string{
				`-rule=/config/alert-rules.yaml`,
				`-datasource.url=http://victoriametrics-victoria-metrics-single-server:8428`,
				`-notifier.url=http://alertmanager:9093`,
				`-remoteRead.url=http://victoriametrics-victoria-metrics-single-server:8428`,
				`-remoteWrite.url=http://victoriametrics-victoria-metrics-single-server:8428`,
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
		}},
		ExtraVolumes: []corev1.Volume{
			{Name: "alerts-config", VolumeSource: corev1.VolumeSource{ConfigMap: &corev1.ConfigMapVolumeSource{LocalObjectReference: corev1.LocalObjectReference{Name: "alerting-rules"}}}},
		},
	})
}
