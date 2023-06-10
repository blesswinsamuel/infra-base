package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type CrowdsecTraefikBouncerProps struct {
	Enabled       bool             `json:"enabled"`
	HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/crowdsecurity/helm-charts/tree/main/charts/crowdsec-traefik-bouncer
func NewCrowdsecTraefikBouncer(scope constructs.Construct, props CrowdsecTraefikBouncerProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: k8sapp.GetNamespaceContextPtr(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("crowdsec-traefik-bouncer"), &cprops)

	k8sapp.NewHelm(chart, jsii.String("helm"), &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("crowdsec-traefik-bouncer"),
		Namespace:   chart.Namespace(),
		Values: map[string]any{
			"bouncer": map[string]any{
				"crowdsec_bouncer_api_key": "test", // TODO
				"crowdsec_agent_host":      "crowdsec-service." + *k8sapp.GetNamespaceContextPtr(scope) + ".svc.cluster.local:8080",
			},
		},
	})

	return chart
}
