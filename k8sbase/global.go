package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

func GetGlobal(scope kubegogen.Scope) k8sapp.ValuesGlobal {
	globalValues := scope.GetContext("global").(string)
	return infrahelpers.FromYamlString[k8sapp.ValuesGlobal](globalValues)
}

func GetCertIssuer(scope kubegogen.Scope) string {
	return GetGlobal(scope).CertIssuer
}

func GetCertIssuerAnnotation(scope kubegogen.Scope) map[string]string {
	return map[string]string{"cert-manager.io/cluster-issuer": GetCertIssuer(scope)}
}

func SetGlobalContext(scope kubegogen.Scope, props k8sapp.ValuesGlobal) {
	scope.SetContext("global", infrahelpers.ToYamlString(props))
}

func GetDomain(scope kubegogen.Scope) string {
	return GetGlobal(scope).Domain
}
