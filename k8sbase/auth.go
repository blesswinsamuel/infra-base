package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/helpers"
)

type AuthProps struct {
	Namespace          string                  `yaml:"namespace"`
	TraefikForwardAuth TraefikForwardAuthProps `yaml:"traefikForwardAuth"`
	Authelia           AutheliaProps           `yaml:"authelia"`
	LLDAP              LLDAPProps              `yaml:"lldap"`
}

func NewAuth(scope constructs.Construct, props AuthProps) constructs.Construct {
	defer logModuleTiming("auth")()
	construct := constructs.NewConstruct(scope, jsii.String("auth"))

	helpers.NewNamespace(construct, props.Namespace)

	NewTraefikForwardAuth(construct, props.TraefikForwardAuth)
	NewAuthelia(construct, props.Authelia)
	NewLLDAP(construct, props.LLDAP)

	return construct
}
