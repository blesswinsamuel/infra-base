package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type AuthProps struct {
	Namespace          string                  `yaml:"namespace"`
	TraefikForwardAuth TraefikForwardAuthProps `yaml:"traefikForwardAuth"`
	Authelia           AutheliaProps           `yaml:"authelia"`
	LLDAP              LLDAPProps              `yaml:"lldap"`
}

func NewAuth(scope constructs.Construct, props AuthProps) constructs.Construct {
	defer logModuleTiming("auth")()
	chart := k8sapp.NewNamespaceChart(scope, props.Namespace)

	NewTraefikForwardAuth(chart, props.TraefikForwardAuth)
	NewAuthelia(chart, props.Authelia)
	NewLLDAP(chart, props.LLDAP)

	return chart
}
