package k8sapp

import (
	"github.com/blesswinsamuel/kgen"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
)

type NetworkPolicy struct {
	Name string
}

// https://editor.networkpolicy.io/
// https://www.suse.com/c/rancher_blog/k3s-network-policy/
// https://ranchermanager.docs.rancher.com/reference-guides/rancher-security/hardening-guides/k3s-hardening-guide
// https://docs.k3s.io/security/hardening-guide#networkpolicies

func NewNamespaceNetworkPolicies(scope kgen.Scope) {
	scope.AddApiObject(&networkingv1.NetworkPolicy{
		ObjectMeta: v1.ObjectMeta{
			Name: "default-deny-all-ingress",
		},
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: v1.LabelSelector{},
			Ingress:     []networkingv1.NetworkPolicyIngressRule{},
		},
	})
	scope.AddApiObject(&networkingv1.NetworkPolicy{
		ObjectMeta: v1.ObjectMeta{
			Name: "allow-egress-to-kube-dns",
		},
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: v1.LabelSelector{},
			Egress: []networkingv1.NetworkPolicyEgressRule{
				{
					To:    []networkingv1.NetworkPolicyPeer{{PodSelector: &v1.LabelSelector{MatchLabels: map[string]string{"k8s-app": "kube-dns"}}}},
					Ports: []networkingv1.NetworkPolicyPort{{Port: ptr.To(intstr.FromInt(53)), Protocol: ptr.To(corev1.ProtocolUDP)}},
				},
			},
		},
	})
}

func NewKubeSystemNetworkPolicies(scope kgen.Scope) {
	// https://docs.k3s.io/security/hardening-guide#networkpolicies
	scope.AddApiObject(&networkingv1.NetworkPolicy{
		ObjectMeta: v1.ObjectMeta{
			Name: "allow-intra-namespace",
		},
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: v1.LabelSelector{},
			Ingress: []networkingv1.NetworkPolicyIngressRule{
				{
					From: []networkingv1.NetworkPolicyPeer{
						{NamespaceSelector: &v1.LabelSelector{MatchLabels: map[string]string{"name": "kube-system"}}},
					},
				},
			},
		},
	})
	scope.AddApiObject(&networkingv1.NetworkPolicy{
		ObjectMeta: v1.ObjectMeta{
			Name: "allow-ingress-to-coredns",
		},
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: v1.LabelSelector{MatchLabels: map[string]string{"k8s-app": "kube-dns"}},
			Ingress: []networkingv1.NetworkPolicyIngressRule{
				{
					Ports: []networkingv1.NetworkPolicyPort{
						{Port: ptr.To(intstr.FromInt(53)), Protocol: ptr.To(corev1.ProtocolTCP)},
						{Port: ptr.To(intstr.FromInt(53)), Protocol: ptr.To(corev1.ProtocolUDP)},
					},
				},
			},
		},
	})
	scope.AddApiObject(&networkingv1.NetworkPolicy{
		ObjectMeta: v1.ObjectMeta{
			Name: "allow-ingress-to-metrics-server",
		},
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: v1.LabelSelector{MatchLabels: map[string]string{"k8s-app": "metrics-server"}},
			Ingress: []networkingv1.NetworkPolicyIngressRule{
				{},
			},
			PolicyTypes: []networkingv1.PolicyType{networkingv1.PolicyTypeIngress},
		},
	})
}

func NewNetworkPolicy(scope kgen.Scope, props *NetworkPolicy) kgen.ApiObject {
	return scope.AddApiObject(&networkingv1.NetworkPolicy{
		ObjectMeta: v1.ObjectMeta{
			Name: props.Name,
		},
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: v1.LabelSelector{},
			Ingress:     []networkingv1.NetworkPolicyIngressRule{},
			Egress:      []networkingv1.NetworkPolicyEgressRule{},
			PolicyTypes: []networkingv1.PolicyType{},
		},
	})
}
