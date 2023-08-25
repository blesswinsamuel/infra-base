package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
)

func GetGlobal(scope packager.Construct) k8sapp.GlobalProps {
	globalValues := scope.GetContext("global").(string)
	return infrahelpers.FromYamlString[k8sapp.GlobalProps](globalValues)
}

func GetCertIssuer(scope packager.Construct) string {
	return GetGlobal(scope).CertIssuer
}

func GetCertIssuerAnnotation(scope packager.Construct) map[string]string {
	return map[string]string{"cert-manager.io/cluster-issuer": GetCertIssuer(scope)}
}

func SetGlobalContext(scope packager.Construct, props k8sapp.GlobalProps) {
	scope.SetContext("global", infrahelpers.ToYamlString(props))
}

func GetDomain(scope packager.Construct) string {
	return GetGlobal(scope).Domain
}
