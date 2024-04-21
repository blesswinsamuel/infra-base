package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type CrowdsecTraefikBouncerProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/crowdsecurity/helm-charts/tree/main/charts/crowdsec-traefik-bouncer
func (props *CrowdsecTraefikBouncerProps) Render(scope kubegogen.Scope) {
	k8sapp.NewHelm(scope, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "crowdsec-traefik-bouncer",
		Namespace:   scope.Namespace(),
		Values: map[string]any{
			"bouncer": map[string]any{
				"crowdsec_bouncer_api_key": "test", // TODO
				"crowdsec_agent_host":      "crowdsec-service." + scope.Namespace() + ".svc.cluster.local:8080",
			},
		},
	})
}
