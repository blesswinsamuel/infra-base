package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type AlertingRulesProps struct {
	Rules map[string]k8sapp.AlertingRuleProps
}

func (props *AlertingRulesProps) Chart(scope kubegogen.Construct) kubegogen.Construct {
	cprops := kubegogen.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("alerting-rules", cprops)

	k8sapp.NewAlertingRules(chart, props.Rules)

	return chart
}
