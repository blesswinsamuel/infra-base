package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
)

type AuthProps struct {
	Enabled            bool                    `json:"enabled"`
	TraefikForwardAuth TraefikForwardAuthProps `json:"traefikForwardAuth"`
	Authelia           AutheliaProps           `json:"authelia"`
	LLDAP              LLDAPProps              `json:"lldap"`
}

func NewAuth(scope packager.Construct, props AuthProps) packager.Construct {
	if !props.Enabled {
		return nil
	}
	defer logModuleTiming("auth")()
	chart := k8sapp.NewNamespaceChart(scope, "auth")

	NewTraefikForwardAuth(chart, props.TraefikForwardAuth)
	NewAuthelia(chart, props.Authelia)
	NewLLDAP(chart, props.LLDAP)

	return chart
}
