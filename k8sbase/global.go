package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

func GetGlobals(scope kubegogen.Scope) k8sapp.ValuesGlobal {
	return scope.GetContext("global").(k8sapp.ValuesGlobal)
}

func SetGlobals(scope kubegogen.Scope, props k8sapp.ValuesGlobal) {
	scope.SetContext("global", props)
}

func GetDomain(scope kubegogen.Scope) string {
	return GetGlobals(scope).Domain
}
