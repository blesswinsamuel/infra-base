package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
)

func init() {
	RegisterModule("cilium", &CiliumProps{})
}

type CiliumProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://docs.cilium.io/en/stable/helm-reference/
func (props *CiliumProps) Render(scope kgen.Scope) {
	// https://github.com/cilium/cilium/tree/main/install/kubernetes/cilium
	// https://blogs.learningdevops.com/the-complete-guide-to-setting-up-cilium-on-k3s-with-kubernetes-gateway-api-8f78adcddb4d
	// https://github.com/cilium/cilium/blob/main/install/kubernetes/cilium/values.yaml
	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "cilium",
		Values: map[string]any{
			// "installCRDs": "true",
			"nodeIPAM": map[string]any{
				"enabled": true,
			},
			"autoDirectNodeRoutes": true,
			"defaultLBServiceIPAM": "nodeipam",

			// "routingMode":           "native",
			// "ipv4NativeRoutingCIDR": "10.42.0.0/16",
			"ipam": map[string]any{
				"mode": "kubernetes",
				"operator": map[string]any{
					"clusterPoolIPv4PodCIDRList": []string{"10.42.0.0/16"},
				},
			},
			"operator": map[string]any{
				"replicas":          1,
				"rollOutPods":       true,
				"rollOutCiliumPods": true,
			},
			"nodePort": map[string]any{
				"enabled": true,
			},
			// "gatewayAPI": map[string]any{
			// 	"enabled": true,
			// },
			// "securityContext": map[string]any{
			// 	"privileged": true,
			// },
			// "externalIPs": map[string]any{
			// 	"enabled": true,
			// },
			// "hubble": map[string]any{
			// 	"tls": map[string]any{
			// 		"enabled": false,
			// 		"auto": map[string]any{
			// 			"method": "cronJob",
			// 			// "method": "certmanager",
			// 			// "certManagerIssuerRef": map[string]any{
			// 			// 	"group": "cert-manager.io",
			// 			// 	"kind":  k8sapp.GetGlobals(scope).Cert.CertIssuerKind,
			// 			// 	"name":  k8sapp.GetGlobals(scope).Cert.CertIssuerName,
			// 			// },
			// 		},
			// 	},
			// 	"relay": map[string]any{
			// 		"enabled": true,
			// 	},
			// 	"ui": map[string]any{
			// 		"enabled": true,
			// 	},
			// },
			// "secretsBackend": map[string]any{
			// 	"secretsBackend": "k8s",
			// },
		},
		PatchObject: helmPatchCleanLabelsAndAnnotations,
	})
}
