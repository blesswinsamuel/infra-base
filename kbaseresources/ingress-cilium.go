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
	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "cilium",
		Values: map[string]any{
			// "installCRDs": "true",
			"routingMode":           "native",
			"ipv4NativeRoutingCIDR": "10.0.0.0/8",
			"autoDirectNodeRoutes":  true,
			"ipam": map[string]any{
				"mode": "kubernetes",
				"operator": map[string]any{
					"clusterPoolIPv4PodCIDRList": []string{"10.0.0.0/8"},
				},
			},
			"operator": map[string]any{
				"replicas": 1,
			},
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
