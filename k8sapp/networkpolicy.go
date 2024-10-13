package k8sapp

import (
	"github.com/blesswinsamuel/kgen"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
)

// https://editor.networkpolicy.io/
// https://www.suse.com/c/rancher_blog/k3s-network-policy/
// https://ranchermanager.docs.rancher.com/reference-guides/rancher-security/hardening-guides/k3s-hardening-guide
// https://docs.k3s.io/security/hardening-guide#networkpolicies

func NewNamespaceDefaultNetworkPolicies(scope kgen.Scope) {
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
			Name: "allow-egress-to-coredns",
		},
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: v1.LabelSelector{},
			Egress: []networkingv1.NetworkPolicyEgressRule{
				{
					To: []networkingv1.NetworkPolicyPeer{
						{
							NamespaceSelector: &v1.LabelSelector{MatchLabels: map[string]string{"kubernetes.io/metadata.name": "kube-system"}},
							PodSelector:       &v1.LabelSelector{MatchLabels: map[string]string{"k8s-app": "kube-dns"}},
						},
					},
					Ports: []networkingv1.NetworkPolicyPort{{Port: ptr.To(intstr.FromInt(53)), Protocol: ptr.To(corev1.ProtocolUDP)}},
				},
			},
		},
	})
}

func NewGlobalNetworkPolicies(scope kgen.Scope) {
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

type NetworkPolicyIngress struct {
	AllowFromPods    []NetworkPolicyPeer
	AllowFromTraefik []intstr.IntOrString
}

type NetworkPolicyPeer struct {
	Namespace string
	Pod       string
	Ports     []intstr.IntOrString
}

type NetworkPolicyEgressIP struct {
	CidrIPBlocks []string
	Ports        []intstr.IntOrString
}

type NetworkPolicyEgress struct {
	AllowToPods        []NetworkPolicyPeer
	AllowToAllInternet []int
	AllowToIPs         []NetworkPolicyEgressIP
}

type NetworkPolicy struct {
	Name    string
	Ingress NetworkPolicyIngress
	Egress  NetworkPolicyEgress
}

func NewNetworkPolicy(scope kgen.Scope, props *NetworkPolicy) kgen.ApiObject {
	var ingressRules []networkingv1.NetworkPolicyIngressRule
	if len(props.Ingress.AllowFromTraefik) > 0 {
		rule := NetworkPolicyPeer{
			Namespace: "ingress",
			Pod:       "traefik",
			Ports:     props.Ingress.AllowFromTraefik,
		}
		props.Ingress.AllowFromPods = append(props.Ingress.AllowFromPods, rule)
		// rule := networkingv1.NetworkPolicyIngressRule{
		// 	From: []networkingv1.NetworkPolicyPeer{
		// 		{NamespaceSelector: &v1.LabelSelector{MatchLabels: map[string]string{"kubernetes.io/metadata.name": "ingress"}}},
		// 		{PodSelector: &v1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": "traefik"}}},
		// 	},
		// 	Ports: []networkingv1.NetworkPolicyPort{},
		// }
		// for _, port := range props.Ingress.AllowFromTraefikToPorts {
		// 	rule.Ports = append(rule.Ports, networkingv1.NetworkPolicyPort{Port: ptr.To(port)})
		// }
		// ingressRules = append(ingressRules, rule)
	}
	for _, ingressPod := range props.Ingress.AllowFromPods {
		peer := networkingv1.NetworkPolicyPeer{}
		peer.PodSelector = &v1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": ingressPod.Pod}}
		if ingressPod.Namespace != "" {
			peer.NamespaceSelector = &v1.LabelSelector{MatchLabels: map[string]string{"kubernetes.io/metadata.name": ingressPod.Namespace}}
		}
		ingressRule := networkingv1.NetworkPolicyIngressRule{
			From:  []networkingv1.NetworkPolicyPeer{peer},
			Ports: []networkingv1.NetworkPolicyPort{},
		}
		for _, port := range ingressPod.Ports {
			ingressRule.Ports = append(ingressRule.Ports, networkingv1.NetworkPolicyPort{Port: ptr.To(port)})
		}
		ingressRules = append(ingressRules, ingressRule)
	}

	var egressRules []networkingv1.NetworkPolicyEgressRule
	allowCorednsRule := networkingv1.NetworkPolicyEgressRule{
		To: []networkingv1.NetworkPolicyPeer{
			{
				NamespaceSelector: &v1.LabelSelector{MatchLabels: map[string]string{"kubernetes.io/metadata.name": "kube-system"}},
				PodSelector:       &v1.LabelSelector{MatchLabels: map[string]string{"k8s-app": "kube-dns"}},
			},
		},
		Ports: []networkingv1.NetworkPolicyPort{{Port: ptr.To(intstr.FromInt(53)), Protocol: ptr.To(corev1.ProtocolUDP)}},
	}
	egressRules = append(egressRules, allowCorednsRule)

	if len(props.Egress.AllowToAllInternet) > 0 {
		rule := NetworkPolicyEgressIP{
			CidrIPBlocks: []string{"0.0.0.0/0"},
			Ports:        []intstr.IntOrString{},
		}
		for _, port := range props.Egress.AllowToAllInternet {
			rule.Ports = append(rule.Ports, intstr.FromInt(port))
		}
		props.Egress.AllowToIPs = append(props.Egress.AllowToIPs, rule)
	}
	for _, egressPod := range props.Egress.AllowToPods {
		peer := networkingv1.NetworkPolicyPeer{}
		peer.PodSelector = &v1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": egressPod.Pod}}
		if egressPod.Namespace != "" {
			peer.NamespaceSelector = &v1.LabelSelector{MatchLabels: map[string]string{"kubernetes.io/metadata.name": egressPod.Namespace}}
		}
		egressRule := networkingv1.NetworkPolicyEgressRule{
			To:    []networkingv1.NetworkPolicyPeer{peer},
			Ports: []networkingv1.NetworkPolicyPort{},
		}
		for _, port := range egressPod.Ports {
			egressRule.Ports = append(egressRule.Ports, networkingv1.NetworkPolicyPort{Port: ptr.To(port)})
		}
		egressRules = append(egressRules, egressRule)
	}
	for _, egressIP := range props.Egress.AllowToIPs {
		egressRule := networkingv1.NetworkPolicyEgressRule{
			To:    []networkingv1.NetworkPolicyPeer{},
			Ports: []networkingv1.NetworkPolicyPort{},
		}
		for _, ipBlock := range egressIP.CidrIPBlocks {
			egressRule.To = append(egressRule.To, networkingv1.NetworkPolicyPeer{IPBlock: &networkingv1.IPBlock{CIDR: ipBlock}})
		}
		for _, port := range egressIP.Ports {
			egressRule.Ports = append(egressRule.Ports, networkingv1.NetworkPolicyPort{Port: ptr.To(port)})
		}
		egressRules = append(egressRules, egressRule)
	}

	return scope.AddApiObject(&networkingv1.NetworkPolicy{
		ObjectMeta: v1.ObjectMeta{
			Name: props.Name,
		},
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: v1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/name": props.Name,
				},
			},
			Ingress: ingressRules,
			Egress:  egressRules,
			PolicyTypes: []networkingv1.PolicyType{
				networkingv1.PolicyTypeIngress,
				networkingv1.PolicyTypeEgress,
			},
		},
	})
}
