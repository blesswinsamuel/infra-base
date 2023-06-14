package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
)

type CrowdsecTraefikBouncerProps struct {
	Enabled       bool             `json:"enabled"`
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/crowdsecurity/helm-charts/tree/main/charts/crowdsec-traefik-bouncer
func NewCrowdsecTraefikBouncer(scope packager.Construct, props CrowdsecTraefikBouncerProps) packager.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := packager.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("crowdsec-traefik-bouncer", cprops)

	k8sapp.NewHelm(chart, "helm", &k8sapp.HelmProps{
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
