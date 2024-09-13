package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	"go.universe.tf/metallb/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	RegisterModule("metallb", &MetalLBProps{})
}

type MetalLBProps struct {
	ChartInfo k8sapp.ChartInfo `json:"helm"`
	Speaker   struct {
		Tolerations []corev1.Toleration `json:"tolerations"`
	} `json:"speaker"`
}

// https://github.com/metallb/metallb/tree/main/charts/metallb
func (props *MetalLBProps) Render(scope kgen.Scope) {
	values := map[string]interface{}{
		"speaker": map[string]interface{}{
			"tolerations": props.Speaker.Tolerations,
		},
	}
	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
		ChartInfo:   props.ChartInfo,
		ReleaseName: "metallb",
		Values:      values,
		PatchObject: helmPatchCleanLabelsAndAnnotations,
	})
	scope.AddApiObject(&v1beta1.IPAddressPool{
		ObjectMeta: metav1.ObjectMeta{Name: "default"},
		Spec: v1beta1.IPAddressPoolSpec{
			Addresses: []string{
				"10.100.1.220/32",
				"10.100.1.221/32",
				"10.100.1.222/32",
				"10.100.1.223/32",
				"10.100.1.224/32",
			},
		},
	})
}
