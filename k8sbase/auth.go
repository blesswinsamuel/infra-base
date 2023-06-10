package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type AuthProps struct {
	Namespace          string                  `json:"namespace"`
	TraefikForwardAuth TraefikForwardAuthProps `json:"traefikForwardAuth"`
	Authelia           AutheliaProps           `json:"authelia"`
	LLDAP              LLDAPProps              `json:"lldap"`
}

func NewAuth(scope constructs.Construct, props AuthProps) constructs.Construct {
	defer logModuleTiming("auth")()
	chart := k8sapp.NewNamespaceChart(scope, props.Namespace)

	NewTraefikForwardAuth(chart, props.TraefikForwardAuth)
	NewAuthelia(chart, props.Authelia)
	NewLLDAP(chart, props.LLDAP)

	return chart
}
