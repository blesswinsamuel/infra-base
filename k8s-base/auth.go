package resourcesbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AuthProps struct {
	Namespace          string                  `yaml:"namespace"`
	TraefikForwardAuth TraefikForwardAuthProps `yaml:"traefikForwardAuth"`
	Authelia           AutheliaProps           `yaml:"authelia"`
	LLDAP              LLDAPProps              `yaml:"lldap"`
}

func NewAuth(scope constructs.Construct, props AuthProps) constructs.Construct {
	construct := constructs.NewConstruct(scope, jsii.String("auth"))

	NewNamespace(construct, props.Namespace)

	NewTraefikForwardAuth(construct, props.TraefikForwardAuth)
	NewAuthelia(construct, props.Authelia)
	NewLLDAP(construct, props.LLDAP)

	return construct
}
