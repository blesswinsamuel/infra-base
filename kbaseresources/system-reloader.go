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
	RegisterModule("reloader", &ReloaderProps{})
}

type ReloaderProps struct {
	ImageInfo k8sapp.ImageInfo `json:"image"`
}

// https://github.com/stakater/Reloader/blob/master/deployments/kubernetes/chart/reloader

func (props *ReloaderProps) Render(scope kgen.Scope) {
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Name:                 "reloader-reloader",
		ServiceAccountName:   "reloader-reloader",
		CreateServiceAccount: true,
		Containers: []k8sapp.ApplicationContainer{
			{
				Name:           "reloader",
				Args:           []string{"--reload-strategy=annotations"},
				Image:          props.ImageInfo,
				Ports:          []k8sapp.ContainerPort{{Name: "http", Port: 9090, PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{}}},
				LivenessProbe:  &v1.Probe{FailureThreshold: 5, ProbeHandler: v1.ProbeHandler{HTTPGet: &v1.HTTPGetAction{Path: "/live", Port: intstr.FromString("http")}}, InitialDelaySeconds: 10, PeriodSeconds: 10, SuccessThreshold: 1, TimeoutSeconds: 5},
				ReadinessProbe: &v1.Probe{FailureThreshold: 5, ProbeHandler: v1.ProbeHandler{HTTPGet: &v1.HTTPGetAction{Path: "/metrics", Port: intstr.FromString("http")}}, InitialDelaySeconds: 10, PeriodSeconds: 10, SuccessThreshold: 1, TimeoutSeconds: 5},
			},
		},
		PodSecurityContext: &v1.PodSecurityContext{
			RunAsNonRoot:   ptr.To(true),
			RunAsUser:      ptr.To(int64(65534)),
			SeccompProfile: &v1.SeccompProfile{Type: v1.SeccompProfileTypeRuntimeDefault},
		},
	})
	scope.AddApiObject(&rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{Name: "reloader-reloader-role"},
		Rules: []rbacv1.PolicyRule{
			{APIGroups: []string{""}, Resources: []string{"secrets", "configmaps"}, Verbs: []string{"list", "get", "watch"}},
			{APIGroups: []string{"apps"}, Resources: []string{"deployments", "daemonsets", "statefulsets"}, Verbs: []string{"list", "get", "update", "patch"}},
			{APIGroups: []string{"extensions"}, Resources: []string{"deployments", "daemonsets"}, Verbs: []string{"list", "get", "update", "patch"}},
			{APIGroups: []string{"batch"}, Resources: []string{"cronjobs"}, Verbs: []string{"list", "get"}},
			{APIGroups: []string{"batch"}, Resources: []string{"jobs"}, Verbs: []string{"create"}},
			{APIGroups: []string{""}, Resources: []string{"events"}, Verbs: []string{"create", "patch"}},
		},
	})
	scope.AddApiObject(&rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{Name: "reloader-reloader-role-binding"},
		Subjects:   []rbacv1.Subject{{Kind: "ServiceAccount", Name: "reloader-reloader", Namespace: scope.Namespace()}},
		RoleRef:    rbacv1.RoleRef{Kind: "ClusterRole", Name: "reloader-reloader-role", APIGroup: "rbac.authorization.k8s.io"},
	})
}
