package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
)

type AlertingRulesProps struct {
	Rules map[string]k8sapp.AlertingRules
}

func (props *AlertingRulesProps) Render(scope kgen.Scope) {
	k8sapp.NewAlertingRules(scope, props.Rules)
}
