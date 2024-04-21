package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type CrowdsecTraefikBouncerProps struct {
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/crowdsecurity/helm-charts/tree/main/charts/crowdsec-traefik-bouncer
func (props *CrowdsecTraefikBouncerProps) Chart(scope kubegogen.Scope) kubegogen.Scope {
	cprops := kubegogen.ScopeProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.CreateScope("crowdsec-traefik-bouncer", cprops)

	k8sapp.NewHelm(chart, &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: "crowdsec-traefik-bouncer",
		Namespace:   chart.Namespace(),
		Values: map[string]any{
			"bouncer": map[string]any{
				"crowdsec_bouncer_api_key": "test", // TODO
				"crowdsec_agent_host":      "crowdsec-service." + k8sapp.GetNamespaceContext(scope) + ".svc.cluster.local:8080",
			},
		},
	})

	return chart
}
