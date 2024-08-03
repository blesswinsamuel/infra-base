package kbaseresources

import (
	_ "embed"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
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
	PersistentVolumeName string              `json:"persistentVolumeName"`
	Tolerations          []corev1.Toleration `json:"tolerations"`
}

//go:embed vmagent-config.yaml
var vmagentConfig string

// https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-agent
func (props *VmagentProps) Render(scope kgen.Scope) {
	vmagentConfig := infrahelpers.FromYamlString[map[string]any](vmagentConfig)

	extraScrapeConfigs := []any{}
	for _, extraScrapeConfig := range props.ExtraScrapeConfigs {
		extraScrapeConfigs = append(extraScrapeConfigs, extraScrapeConfig)
	}
	vmagentConfig["scrape_configs"] = append(vmagentConfig["scrape_configs"].([]any), extraScrapeConfigs...)

	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Name:               "vmagent",
		Kind:               "StatefulSet",
		ServiceAccountName: "vmagent",
		Containers: []k8sapp.ApplicationContainer{{
			Name:  "vmagent",
			Image: props.ImageInfo,
			Ports: []k8sapp.ContainerPort{
				{Name: "http", Port: 8429, Ingress: infrahelpers.If(props.Ingress.Enabled, &k8sapp.ApplicationIngress{Host: props.Ingress.SubDomain + "." + k8sapp.GetDomain(scope)}, nil)},
			},
			Args: []string{
				"-promscrape.config=/config/scrape.yml",
				"-remoteWrite.tmpDataPath=/tmpdata",
				"-remoteWrite.url=http://victoriametrics:8428/api/v1/write",
				"-envflag.enable=true",
				"-envflag.prefix=VM_",
				"-loggerFormat=json",
				"-promscrape.streamParse=true",
			},
			LivenessProbe:  &corev1.Probe{InitialDelaySeconds: 5, PeriodSeconds: 15, TimeoutSeconds: 5, ProbeHandler: corev1.ProbeHandler{TCPSocket: &corev1.TCPSocketAction{Port: intstr.FromString("http")}}},
			ReadinessProbe: &corev1.Probe{InitialDelaySeconds: 5, PeriodSeconds: 15, TimeoutSeconds: 5, ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/health", Port: intstr.FromString("http")}}},
			Resources:      props.Resources,
		}},
		ConfigMaps: []k8sapp.ApplicationConfigMap{
			{Name: "vmagent", Data: map[string]string{"scrape.yml": infrahelpers.ToYamlString(vmagentConfig)}, MountPath: "/config", MountName: "config"},
		},
		PersistentVolumes: []k8sapp.ApplicationPersistentVolume{
			{Name: "vmagent-tmpdata", RequestsStorage: "1Gi", MountName: "tmpdata", MountPath: "/tmpdata", VolumeName: props.PersistentVolumeName},
		},
		Tolerations: props.Tolerations,
	})
	scope.AddApiObject(&rbacv1.ClusterRole{
		ObjectMeta: v1.ObjectMeta{Name: "vmagent"},
		Rules: []rbacv1.PolicyRule{
			{APIGroups: []string{"discovery.k8s.io"}, Resources: []string{"endpointslices"}, Verbs: []string{"get", "list", "watch"}},
			{APIGroups: []string{""}, Resources: []string{"nodes", "nodes/proxy", "nodes/metrics", "services", "endpoints", "pods"}, Verbs: []string{"get", "list", "watch"}},
			{APIGroups: []string{"extensions", "networking.k8s.io"}, Resources: []string{"ingresses"}, Verbs: []string{"get", "list", "watch"}},
			{NonResourceURLs: []string{"/metrics"}, Verbs: []string{"get"}},
		},
	})
	scope.AddApiObject(&rbacv1.ClusterRoleBinding{
		ObjectMeta: v1.ObjectMeta{Name: "vmagent"},
		Subjects:   []rbacv1.Subject{{Kind: "ServiceAccount", Name: "vmagent", Namespace: scope.Namespace()}},
		RoleRef:    rbacv1.RoleRef{Kind: "ClusterRole", Name: "vmagent", APIGroup: "rbac.authorization.k8s.io"},
	})
	scope.AddApiObject(&corev1.ServiceAccount{
		ObjectMeta: v1.ObjectMeta{Name: "vmagent"},
	})
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
