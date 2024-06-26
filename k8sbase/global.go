package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
)

func GetGlobals(scope kgen.Scope) k8sapp.ValuesGlobal {
	return scope.GetContext("global").(k8sapp.ValuesGlobal)
}

func SetGlobals(scope kgen.Scope, props k8sapp.ValuesGlobal) {
	scope.SetContext("global", props)
}

func GetDomain(scope kgen.Scope) string {
	return GetGlobals(scope).Domain
}
