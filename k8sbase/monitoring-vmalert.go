package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/helpers"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/k8s"
)

type VmalertProps struct {
	Enabled   bool              `yaml:"enabled"`
	ImageInfo helpers.ImageInfo `yaml:"image"`
	Resources *Resources        `yaml:"resources"`
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

	return NewApplication(scope, jsii.String("vmalert"), &ApplicationProps{
		Name: "vmalert",
		Containers: []ApplicationContainer{{
			Name:      "vmalert",
			ImageInfo: props.ImageInfo,
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
			Ports: []ContainerPort{{Name: "http", Port: 8880}},
			LivenessProbe: &k8s.Probe{
				InitialDelaySeconds: jsii.Number(5),
				PeriodSeconds:       jsii.Number(15),
				TcpSocket:           &k8s.TcpSocketAction{Port: k8s.IntOrString_FromString(jsii.String("http"))},
				TimeoutSeconds:      jsii.Number(5),
			},
			ReadinessProbe: &k8s.Probe{
				InitialDelaySeconds: jsii.Number(5),
				PeriodSeconds:       jsii.Number(15),
				HttpGet:             &k8s.HttpGetAction{Port: k8s.IntOrString_FromString(jsii.String("http")), Path: jsii.String("/health")},
			},
			ExtraVolumeMounts: []*k8s.VolumeMount{
				{Name: jsii.String("alerts-config"), MountPath: jsii.String("/config")},
			},
		}},
		Ingress: []ApplicationIngress{{
			Host:     props.Ingress.SubDomain + "." + GetDomain(scope),
			PortName: "http",
		}},
		ExtraVolumes: []*k8s.Volume{
			{Name: jsii.String("alerts-config"), ConfigMap: &k8s.ConfigMapVolumeSource{Name: jsii.String("alerting-rules")}},
		},
	})
}
