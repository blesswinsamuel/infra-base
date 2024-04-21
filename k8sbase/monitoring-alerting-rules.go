package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type AlertingRulesProps struct {
	Rules map[string]k8sapp.AlertingRule
}

func (props *AlertingRulesProps) Chart(scope kubegogen.Scope) kubegogen.Scope {
	cprops := kubegogen.ScopeProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.CreateScope("alerting-rules", cprops)

	k8sapp.NewAlertingRules(chart, props.Rules)

	return chart
}
