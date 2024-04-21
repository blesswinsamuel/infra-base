package k8sbase

import (
	_ "embed"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type VmagentProps struct {
	ImageInfo          k8sapp.ImageInfo            `json:"image"`
	ExtraScrapeConfigs []map[string]any            `json:"extraScrapeConfigs"`
	Resources          corev1.ResourceRequirements `json:"resources"`
	Ingress            struct {
		Enabled   bool   `json:"enabled"`
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
}

//go:embed vmagent-config.yaml
var vmagentConfig string

// https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-agent
func (props *VmagentProps) Chart(scope kubegogen.Scope) kubegogen.Scope {
	vmagentConfig := infrahelpers.FromYamlString[map[string]any](vmagentConfig)

	extraScrapeConfigs := []any{}
	for _, extraScrapeConfig := range props.ExtraScrapeConfigs {
		extraScrapeConfigs = append(extraScrapeConfigs, extraScrapeConfig)
	}
	vmagentConfig["scrape_configs"] = append(vmagentConfig["scrape_configs"].([]any), extraScrapeConfigs...)

	app := k8sapp.NewApplicationChart(scope, "vmagent", &k8sapp.ApplicationProps{
		Name:               "vmagent",
		ServiceAccountName: "vmagent",
		Containers: []k8sapp.ApplicationContainer{{
			Name:  "vmagent",
			Image: props.ImageInfo,
			Ports: []k8sapp.ContainerPort{
				{Name: "http", Port: 8429, Ingress: infrahelpers.If(props.Ingress.Enabled, &k8sapp.ApplicationIngress{Host: props.Ingress.SubDomain + "." + GetDomain(scope)}, nil)},
			},
			Args: []string{
				"-promscrape.config=/config/scrape.yml",
				"-remoteWrite.tmpDataPath=/tmpData",
				"-remoteWrite.url=http://victoriametrics-victoria-metrics-single-server:8428/api/v1/write",
				"-envflag.enable=true",
				"-envflag.prefix=VM_",
				"-loggerFormat=json",
				"-promscrape.streamParse=true",
			},
			LivenessProbe:  &corev1.Probe{InitialDelaySeconds: 5, PeriodSeconds: 15, TimeoutSeconds: 5, ProbeHandler: corev1.ProbeHandler{TCPSocket: &corev1.TCPSocketAction{Port: intstr.FromString("http")}}},
			ReadinessProbe: &corev1.Probe{InitialDelaySeconds: 5, PeriodSeconds: 15, TimeoutSeconds: 5, ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/health", Port: intstr.FromString("http")}}},
			Resources:      props.Resources,
			ExtraVolumeMounts: []corev1.VolumeMount{
				{Name: "tmpdata", MountPath: "/tmpdata"},
			},
		}},
		ConfigMaps: []k8sapp.ApplicationConfigMap{
			{Name: "vmagent", Data: map[string]string{"scrape.yml": infrahelpers.ToYamlString(vmagentConfig)}, MountPath: "/config", MountName: "config"},
		},
		ExtraVolumes: []corev1.Volume{
			{Name: "tmpdata", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}},
		},
	})
	app.AddApiObject(&rbacv1.ClusterRole{
		ObjectMeta: v1.ObjectMeta{Name: "vmagent"},
		Rules: []rbacv1.PolicyRule{
			{APIGroups: []string{"discovery.k8s.io"}, Resources: []string{"endpointslices"}, Verbs: []string{"get", "list", "watch"}},
			{APIGroups: []string{""}, Resources: []string{"nodes", "nodes/proxy", "nodes/metrics", "services", "endpoints", "pods"}, Verbs: []string{"get", "list", "watch"}},
			{APIGroups: []string{"extensions", "networking.k8s.io"}, Resources: []string{"ingresses"}, Verbs: []string{"get", "list", "watch"}},
			{NonResourceURLs: []string{"/metrics"}, Verbs: []string{"get"}},
		},
	})
	app.AddApiObject(&rbacv1.ClusterRoleBinding{
		ObjectMeta: v1.ObjectMeta{Name: "vmagent"},
		Subjects:   []rbacv1.Subject{{Kind: "ServiceAccount", Name: "vmagent", Namespace: app.Namespace()}},
		RoleRef:    rbacv1.RoleRef{Kind: "ClusterRole", Name: "vmagent", APIGroup: "rbac.authorization.k8s.io"},
	})
	app.AddApiObject(&corev1.ServiceAccount{
		ObjectMeta: v1.ObjectMeta{Name: "vmagent"},
	})
	return app
}

// #     extraArgs:
// #       remoteWrite.relabelConfig: /relabel-config/relabel-config.yaml
// #     extraVolumeMounts:
// #     - mountPath: /relabel-config/relabel-config.yaml
// #       subPath: relabel-config.yaml
// #       name: relabel-config
// #       readOnly: true
// #     extraVolumes:
// #     - name: relabel-config
// #       configMap:
// #         name: vmagent-relabel-config
// # ---
// # apiVersion: v1
// # kind: ConfigMap
// # metadata:
// #   name: vmagent-relabel-config
// #   namespace: monitoring
// # data:
// #   relabel-config.yaml: |
// #     - action: labeldrop
// #       regex: "(node_role_kubernetes_io_|node_kubernetes_io_|beta_kubernetes_io_|kubernetes_io_|app_kubernetes_io_|helm_sh_).+"
// #     - action: labeldrop
// #       regex: "(chart|heritage|release|pod_template_hash|objectset_rio_cattle_io_hash)"
// #     - source_labels: job
// #       action: replace
// #       regex: kubernetes-nodes
// #       replacement: apiserver
