package kbaseresources

import (
	_ "embed"
	"strings"

	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	"github.com/muesli/reflow/dedent"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
)

func init() {
	RegisterModule("coredns", &CoreDNSProps{})
}

type CoreDNSProps struct {
	Image          k8sapp.ImageInfo       `json:"image"`
	ClusterDomain  string                 `json:"clusterDomain"`
	ClusterDNS     string                 `json:"clusterDNS"`
	ClusterDNSList []string               `json:"clusterDNSList"`
	IpFamilyPolicy *corev1.IPFamilyPolicy `json:"ipFamilyPolicy"`
}

func (props *CoreDNSProps) Render(scope kgen.Scope) {
	// https://github.com/k3s-io/k3s/blob/master/manifests/coredns.yaml
	scope.AddApiObject(&corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{Name: "coredns"},
	})
	scope.AddApiObject(&rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name:   "system:coredns",
			Labels: map[string]string{"kubernetes.io/bootstrapping": "rbac-defaults"},
		},
		Rules: []rbacv1.PolicyRule{
			{APIGroups: []string{""}, Resources: []string{"endpoints", "services", "pods", "namespaces"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"discovery.k8s.io"}, Resources: []string{"endpointslices"}, Verbs: []string{"list", "watch"}},
		},
	})
	scope.AddApiObject(&rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "system:coredns",
			Annotations: map[string]string{"rbac.authorization.kubernetes.io/autoupdate": "true"},
			Labels:      map[string]string{"kubernetes.io/bootstrapping": "rbac-defaults"},
		},
		RoleRef:  rbacv1.RoleRef{APIGroup: "rbac.authorization.k8s.io", Kind: "ClusterRole", Name: "system:coredns"},
		Subjects: []rbacv1.Subject{{Kind: "ServiceAccount", Name: "coredns", Namespace: scope.Namespace()}},
	})
	scope.AddApiObject(&corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{Name: "coredns"},
		Data: map[string]string{
			"Corefile": strings.TrimSpace(dedent.String(`
				.:53 {
				    errors
				    health
				    ready
				    kubernetes ` + props.ClusterDomain + ` in-addr.arpa ip6.arpa {
				        pods insecure
				        fallthrough in-addr.arpa ip6.arpa
				    }
				    hosts /etc/coredns/NodeHosts {
				        ttl 60
				        reload 15s
				        fallthrough
				    }
				    prometheus :9153
				    forward . /etc/resolv.conf
				    cache 30
				    loop
				    reload
				    loadbalance
				    import /etc/coredns/custom/*.override
				}
				import /etc/coredns/custom/*.server
			`)),
			"NodeHosts": strings.TrimSpace(dedent.String(`
			10.100.1.27 hp-chromebox-g2
			10.100.1.29 beelink-mini-s12-pro
			`)),
		},
	})
	scope.AddApiObject(&appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "coredns",
			Labels: map[string]string{
				"k8s-app":            "kube-dns",
				"kubernetes.io/name": "CoreDNS",
			},
		},
		Spec: appsv1.DeploymentSpec{
			RevisionHistoryLimit: ptr.To[int32](0),
			Strategy:             appsv1.DeploymentStrategy{Type: appsv1.RollingUpdateDeploymentStrategyType, RollingUpdate: &appsv1.RollingUpdateDeployment{MaxUnavailable: ptr.To(intstr.FromInt(1))}},
			Selector:             &metav1.LabelSelector{MatchLabels: map[string]string{"k8s-app": "kube-dns"}},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"k8s-app": "kube-dns"},
				},
				Spec: corev1.PodSpec{
					PriorityClassName:  "system-cluster-critical",
					ServiceAccountName: "coredns",
					Tolerations: []corev1.Toleration{
						{Key: "CriticalAddonsOnly", Operator: corev1.TolerationOpExists},
						{Key: "node-role.kubernetes.io/control-plane", Operator: corev1.TolerationOpExists, Effect: corev1.TaintEffectNoSchedule},
						{Key: "node-role.kubernetes.io/master", Operator: corev1.TolerationOpExists, Effect: corev1.TaintEffectNoSchedule},
					},
					NodeSelector:              map[string]string{"kubernetes.io/os": "linux"},
					TopologySpreadConstraints: []corev1.TopologySpreadConstraint{{MaxSkew: 1, TopologyKey: "kubernetes.io/hostname", WhenUnsatisfiable: corev1.DoNotSchedule, LabelSelector: &metav1.LabelSelector{MatchLabels: map[string]string{"k8s-app": "kube-dns"}}}},
					Containers: []corev1.Container{{
						Name:            "coredns",
						Image:           props.Image.String(),
						ImagePullPolicy: corev1.PullIfNotPresent,
						Resources: corev1.ResourceRequirements{
							Limits:   corev1.ResourceList{"memory": resource.MustParse("250Mi")},
							Requests: corev1.ResourceList{"cpu": resource.MustParse("100m"), "memory": resource.MustParse("70Mi")},
						},
						Args: []string{"-conf", "/etc/coredns/Corefile"},
						VolumeMounts: []corev1.VolumeMount{
							{Name: "config-volume", MountPath: "/etc/coredns", ReadOnly: true},
							{Name: "custom-config-volume", MountPath: "/etc/coredns/custom", ReadOnly: true},
						},
						Ports: []corev1.ContainerPort{
							{Name: "dns", ContainerPort: 53, Protocol: corev1.ProtocolUDP},
							{Name: "dns-tcp", ContainerPort: 53, Protocol: corev1.ProtocolTCP},
							{Name: "metrics", ContainerPort: 9153, Protocol: corev1.ProtocolTCP},
						},
						SecurityContext: &corev1.SecurityContext{
							AllowPrivilegeEscalation: ptr.To(false),
							Capabilities: &corev1.Capabilities{
								Add:  []corev1.Capability{"NET_BIND_SERVICE"},
								Drop: []corev1.Capability{"all"},
							},
							ReadOnlyRootFilesystem: ptr.To(true),
						},
						LivenessProbe:  &corev1.Probe{ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/health", Port: intstr.FromInt(8080), Scheme: corev1.URISchemeHTTP}}, InitialDelaySeconds: 60, PeriodSeconds: 10, TimeoutSeconds: 1, SuccessThreshold: 1, FailureThreshold: 3},
						ReadinessProbe: &corev1.Probe{ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/ready", Port: intstr.FromInt(8181), Scheme: corev1.URISchemeHTTP}}, InitialDelaySeconds: 0, PeriodSeconds: 2, TimeoutSeconds: 1, SuccessThreshold: 1, FailureThreshold: 3},
					}},
					DNSPolicy: corev1.DNSDefault,
					Volumes: []corev1.Volume{
						{Name: "config-volume", VolumeSource: corev1.VolumeSource{ConfigMap: &corev1.ConfigMapVolumeSource{
							LocalObjectReference: corev1.LocalObjectReference{Name: "coredns"},
							Items:                []corev1.KeyToPath{{Key: "Corefile", Path: "Corefile"}, {Key: "NodeHosts", Path: "NodeHosts"}},
						}}},
						{Name: "custom-config-volume", VolumeSource: corev1.VolumeSource{ConfigMap: &corev1.ConfigMapVolumeSource{LocalObjectReference: corev1.LocalObjectReference{Name: "coredns-custom"}, Optional: ptr.To(true)}}},
					},
				},
			},
		},
	})
	scope.AddApiObject(&corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "kube-dns",
			Annotations: map[string]string{
				"prometheus.io/port":   "9153",
				"prometheus.io/scrape": "true",
			},
			Labels: map[string]string{
				"k8s-app":                       "kube-dns",
				"kubernetes.io/cluster-service": "true",
				"kubernetes.io/name":            "CoreDNS",
			},
		},
		Spec: corev1.ServiceSpec{
			Selector:   map[string]string{"k8s-app": "kube-dns"},
			ClusterIP:  props.ClusterDNS,
			ClusterIPs: props.ClusterDNSList,
			Ports: []corev1.ServicePort{
				{Name: "dns", Port: 53, Protocol: corev1.ProtocolUDP},
				{Name: "dns-tcp", Port: 53, Protocol: corev1.ProtocolTCP},
				{Name: "metrics", Port: 9153, Protocol: corev1.ProtocolTCP},
			},
			IPFamilyPolicy: props.IpFamilyPolicy,
		},
	})
}
