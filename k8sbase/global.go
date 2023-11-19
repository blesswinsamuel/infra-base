package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

func GetGlobal(scope kubegogen.Construct) k8sapp.GlobalProps {
	globalValues := scope.GetContext("global").(string)
	return infrahelpers.FromYamlString[k8sapp.GlobalProps](globalValues)
}

func GetCertIssuer(scope kubegogen.Construct) string {
	return GetGlobal(scope).CertIssuer
}

func GetCertIssuerAnnotation(scope kubegogen.Construct) map[string]string {
	return map[string]string{"cert-manager.io/cluster-issuer": GetCertIssuer(scope)}
}

func SetGlobalContext(scope kubegogen.Construct, props k8sapp.GlobalProps) {
	scope.SetContext("global", infrahelpers.ToYamlString(props))
}

func GetDomain(scope kubegogen.Construct) string {
	return GetGlobal(scope).Domain
}
