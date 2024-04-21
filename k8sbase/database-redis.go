package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	corev1 "k8s.io/api/core/v1"
)

type Redis struct {
	HelmChartInfo k8sapp.ChartInfo             `json:"helm"`
	Resources     *corev1.ResourceRequirements `json:"resources"`
}

// https://github.com/bitnami/charts/tree/main/bitnami/redis

func (props *Redis) Render(scope kubegogen.Scope) {
	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "redis",
		Namespace:   scope.Namespace(),
		Values: map[string]interface{}{
			"architecture": "standalone",
			"auth": map[string]interface{}{
				"enabled": false,
			},
			"metrics": map[string]interface{}{
				"enabled":   true,
				"resources": props.Resources,
			},
			"master": map[string]interface{}{
				"resources": props.Resources,
			},
		},
	})
}
