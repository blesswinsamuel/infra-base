package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type AlertingRulesProps struct {
	Rules map[string]k8sapp.AlertingRules
}

func (props *AlertingRulesProps) Render(scope kubegogen.Scope) {
	k8sapp.NewAlertingRules(scope, props.Rules)
}
