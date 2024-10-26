package k8sapp

import (
	"slices"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/kgen"
	"github.com/rs/zerolog/log"
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
	AllowFromAppRefs       map[string][]intstr.IntOrString
	AllowFromApps          []NetworkPolicyPeer
	AllowFromIPs           map[string][]intstr.IntOrString
	AllowFromAllNamespaces []intstr.IntOrString
}

type NetworkPolicyPeer struct {
	Namespace string
	App       string
	Ports     []intstr.IntOrString
}

type NetworkPolicyEgressIP struct {
	CidrIPBlocks []string
	Ports        []int
	Protocol     corev1.Protocol
}

type NetworkPolicyEgress struct {
	AllowToKubeAPIServer bool
	AllowToAppRefs       []string
	AllowToApps          []NetworkPolicyPeer
	AllowToAllInternet   []int
	AllowToAllNamespaces bool
	AllowToIPs           []NetworkPolicyEgressIP
}

type NetworkPolicy struct {
	Name    string
	Ingress NetworkPolicyIngress
	Egress  NetworkPolicyEgress
}

func NewNetworkPolicy(scope kgen.Scope, props *NetworkPolicy) kgen.ApiObject {
	globals := GetGlobals(scope)

	// Ingress
	var ingressRules []networkingv1.NetworkPolicyIngressRule
	allowFromPods := slices.Clone(props.Ingress.AllowFromApps)
	for _, app := range infrahelpers.MapKeysSorted(props.Ingress.AllowFromAppRefs) {
		ports := props.Ingress.AllowFromAppRefs[app]
		if appRef, ok := globals.AppRefs[app]; ok {
			allowFromPods = append(allowFromPods, NetworkPolicyPeer{Namespace: appRef.Namespace, App: appRef.Name, Ports: ports})
		} else {
			log.Panic().Msg("AppRef not found: " + app)
		}
	}
	for _, ingressPod := range allowFromPods {
		peer := networkingv1.NetworkPolicyPeer{}
		peer.PodSelector = &v1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": ingressPod.App}}
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
	for _, cidrBlock := range infrahelpers.MapKeysSorted(props.Ingress.AllowFromIPs) {
		ports := props.Ingress.AllowFromIPs[cidrBlock]
		ingressRule := networkingv1.NetworkPolicyIngressRule{
			From: []networkingv1.NetworkPolicyPeer{{IPBlock: &networkingv1.IPBlock{CIDR: cidrBlock}}},
		}
		for _, port := range ports {
			ingressRule.Ports = append(ingressRule.Ports, networkingv1.NetworkPolicyPort{Port: ptr.To(port)})
		}
		ingressRules = append(ingressRules, ingressRule)
	}
	if len(props.Ingress.AllowFromAllNamespaces) > 0 {
		ingressRule := networkingv1.NetworkPolicyIngressRule{
			// From: []networkingv1.NetworkPolicyPeer{{NamespaceSelector: &v1.LabelSelector{}}},
		}
		if !(len(props.Ingress.AllowFromAllNamespaces) == 1 && props.Ingress.AllowFromAllNamespaces[0] == intstr.FromInt(0)) {
			for _, port := range props.Ingress.AllowFromAllNamespaces {
				ingressRule.Ports = append(ingressRule.Ports, networkingv1.NetworkPolicyPort{Port: ptr.To(port)})
			}
		}
		ingressRules = append(ingressRules, ingressRule)
	}

	// Egress
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
	if props.Egress.AllowToAllNamespaces {
		egressRule := networkingv1.NetworkPolicyEgressRule{
			To: []networkingv1.NetworkPolicyPeer{{NamespaceSelector: &v1.LabelSelector{}}},
		}
		egressRules = append(egressRules, egressRule)
	}
	allowToPods := slices.Clone(props.Egress.AllowToApps)

	for _, app := range props.Egress.AllowToAppRefs {
		if appRef, ok := globals.AppRefs[app]; ok {
			allowToPods = append(allowToPods, NetworkPolicyPeer{Namespace: appRef.Namespace, App: appRef.Name, Ports: []intstr.IntOrString{appRef.Port}})
		} else {
			log.Panic().Msg("AppRef not found: " + app)
		}
	}
	allowToIPs := slices.Clone(props.Egress.AllowToIPs)
	if props.Egress.AllowToKubeAPIServer {
		// allowToPods = append(allowToPods, NetworkPolicyPeer{Namespace: "kube-system", Pod: "kube-apiserver", Ports: []intstr.IntOrString{intstr.FromString("https")}})
		allowToIPs = append(allowToIPs, NetworkPolicyEgressIP{CidrIPBlocks: []string{globals.KubeApiServer.IP}, Ports: []int{globals.KubeApiServer.Port}})
	}

	if len(props.Egress.AllowToAllInternet) > 0 {
		allowToIPs = append(allowToIPs, NetworkPolicyEgressIP{CidrIPBlocks: []string{"0.0.0.0/0"}, Ports: props.Egress.AllowToAllInternet})
	}
	for _, egressPod := range allowToPods {
		peer := networkingv1.NetworkPolicyPeer{}
		if egressPod.App != "" {
			peer.PodSelector = &v1.LabelSelector{MatchLabels: map[string]string{"app.kubernetes.io/name": egressPod.App}}
		}
		if egressPod.Namespace != "" {
			peer.NamespaceSelector = &v1.LabelSelector{MatchLabels: map[string]string{"kubernetes.io/metadata.name": egressPod.Namespace}}
		}
		egressRule := networkingv1.NetworkPolicyEgressRule{
			To: []networkingv1.NetworkPolicyPeer{peer},
		}
		for _, port := range egressPod.Ports {
			egressRule.Ports = append(egressRule.Ports, networkingv1.NetworkPolicyPort{Port: ptr.To(port)})
		}
		egressRules = append(egressRules, egressRule)
	}
	for _, egressIP := range allowToIPs {
		egressRule := networkingv1.NetworkPolicyEgressRule{}
		for _, ipBlock := range egressIP.CidrIPBlocks {
			egressRule.To = append(egressRule.To, networkingv1.NetworkPolicyPeer{IPBlock: &networkingv1.IPBlock{CIDR: ipBlock}})
		}
		for _, port := range egressIP.Ports {
			egressRule.Ports = append(egressRule.Ports, networkingv1.NetworkPolicyPort{Port: ptr.To(intstr.FromInt(port)), Protocol: infrahelpers.PtrIfNonEmpty(egressIP.Protocol)})
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
