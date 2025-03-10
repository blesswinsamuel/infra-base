package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
)

type CertManagerProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/cert-manager/cert-manager/tree/master/deploy/charts/cert-manager
// https://artifacthub.io/packages/helm/cert-manager/cert-manager
func (props *CertManagerProps) Render(scope kgen.Scope) {
	values := map[string]interface{}{
		"installCRDs": true,
		"extraArgs": []string{
			"--enable-certificate-owner-ref=true",
		},
		"dns01RecursiveNameservers":     "8.8.8.8:53,1.1.1.1:53",
		"dns01RecursiveNameserversOnly": true,
		// "podDnsPolicy":                  "None",
		// "podDnsConfig": map[string]interface{}{
		// 	"nameservers": []string{"1.1.1.1", "8.8.8.8"},
		// },
	}
	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "cert-manager",
		Values:      values,
		PatchObject: helmPatchCleanLabelsAndAnnotations,
	})
}
