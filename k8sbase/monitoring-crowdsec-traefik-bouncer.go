package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type CrowdsecTraefikBouncerProps struct {
	Enabled       bool      `yaml:"enabled"`
	HelmChartInfo ChartInfo `yaml:"helm"`
}

// https://github.com/crowdsecurity/helm-charts/tree/main/charts/crowdsec-traefik-bouncer
func NewCrowdsecTraefikBouncer(scope constructs.Construct, props CrowdsecTraefikBouncerProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: GetNamespace(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("crowdsec-traefik-bouncer"), &cprops)

	NewHelmCached(chart, jsii.String("helm"), &HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("crowdsec-traefik-bouncer"),
		Namespace:   chart.Namespace(),
		Values: &map[string]any{
			"bouncer": map[string]any{
				"crowdsec_bouncer_api_key": "test", // TODO
				"crowdsec_agent_host":      "crowdsec-service." + *GetNamespace(scope) + ".svc.cluster.local:8080",
			},
		},
	})

	return chart
}
