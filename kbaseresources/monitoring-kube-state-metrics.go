package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
)

func init() {
	RegisterModule("kube-state-metrics", &KubeStateMetricsProps{})
}

type KubeStateMetricsProps struct {
	ImageInfo k8sapp.ImageInfo `json:"image"`
}

func (props *KubeStateMetricsProps) Render(scope kgen.Scope) {
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Name:                 "kube-state-metrics",
		ServiceAccountName:   "kube-state-metrics",
		CreateServiceAccount: true,
		Containers: []k8sapp.ApplicationContainer{
			{
				Name: "kube-state-metrics",
				Args: []string{
					"--port=8080",
					"--resources=certificatesigningrequests,configmaps,cronjobs,daemonsets,deployments,endpoints,horizontalpodautoscalers,ingresses,jobs,leases,limitranges,mutatingwebhookconfigurations,namespaces,networkpolicies,nodes,persistentvolumeclaims,persistentvolumes,poddisruptionbudgets,pods,replicasets,replicationcontrollers,resourcequotas,secrets,services,statefulsets,storageclasses,validatingwebhookconfigurations,volumeattachments",
				},
				Image:           props.ImageInfo,
				Ports:           []k8sapp.ContainerPort{{Name: "http", Port: 8080, PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{}}},
				LivenessProbe:   &v1.Probe{FailureThreshold: 3, ProbeHandler: v1.ProbeHandler{HTTPGet: &v1.HTTPGetAction{Path: "/healthz", Port: intstr.FromString("http")}}, InitialDelaySeconds: 5, PeriodSeconds: 10, SuccessThreshold: 1, TimeoutSeconds: 5},
				ReadinessProbe:  &v1.Probe{FailureThreshold: 3, ProbeHandler: v1.ProbeHandler{HTTPGet: &v1.HTTPGetAction{Path: "/", Port: intstr.FromString("http")}}, InitialDelaySeconds: 5, PeriodSeconds: 10, SuccessThreshold: 1, TimeoutSeconds: 5},
				SecurityContext: &v1.SecurityContext{AllowPrivilegeEscalation: ptr.To(false), Capabilities: &v1.Capabilities{Drop: []v1.Capability{"ALL"}}, ReadOnlyRootFilesystem: ptr.To(true)},
			},
		},
		PodSecurityContext: &v1.PodSecurityContext{
			RunAsNonRoot:   ptr.To(true),
			FSGroup:        ptr.To(int64(65534)),
			RunAsUser:      ptr.To(int64(65534)),
			RunAsGroup:     ptr.To(int64(65534)),
			SeccompProfile: &v1.SeccompProfile{Type: v1.SeccompProfileTypeRuntimeDefault},
		},
	})
	scope.AddApiObject(&rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{Name: "kube-state-metrics"},
		Rules: []rbacv1.PolicyRule{
			{APIGroups: []string{"certificates.k8s.io"}, Resources: []string{"certificatesigningrequests"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{""}, Resources: []string{"configmaps"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"batch"}, Resources: []string{"cronjobs"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"extensions", "apps"}, Resources: []string{"daemonsets"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"extensions", "apps"}, Resources: []string{"deployments"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{""}, Resources: []string{"endpoints"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"autoscaling"}, Resources: []string{"horizontalpodautoscalers"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"extensions", "networking.k8s.io"}, Resources: []string{"ingresses"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"batch"}, Resources: []string{"jobs"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"coordination.k8s.io"}, Resources: []string{"leases"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{""}, Resources: []string{"limitranges"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"admissionregistration.k8s.io"}, Resources: []string{"mutatingwebhookconfigurations"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{""}, Resources: []string{"namespaces"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"networking.k8s.io"}, Resources: []string{"networkpolicies"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{""}, Resources: []string{"nodes"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{""}, Resources: []string{"persistentvolumeclaims"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{""}, Resources: []string{"persistentvolumes"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"policy"}, Resources: []string{"poddisruptionbudgets"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{""}, Resources: []string{"pods"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"extensions", "apps"}, Resources: []string{"replicasets"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{""}, Resources: []string{"replicationcontrollers"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{""}, Resources: []string{"resourcequotas"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{""}, Resources: []string{"secrets"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{""}, Resources: []string{"services"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"apps"}, Resources: []string{"statefulsets"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"storage.k8s.io"}, Resources: []string{"storageclasses"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"admissionregistration.k8s.io"}, Resources: []string{"validatingwebhookconfigurations"}, Verbs: []string{"list", "watch"}},
			{APIGroups: []string{"storage.k8s.io"}, Resources: []string{"volumeattachments"}, Verbs: []string{"list", "watch"}},
		},
	})
	scope.AddApiObject(&rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{Name: "kube-state-metrics"},
		Subjects:   []rbacv1.Subject{{Kind: "ServiceAccount", Name: "kube-state-metrics", Namespace: scope.Namespace()}},
		RoleRef:    rbacv1.RoleRef{Kind: "ClusterRole", Name: "kube-state-metrics", APIGroup: "rbac.authorization.k8s.io"},
	})
}
