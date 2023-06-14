package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/packager"
)

type GlobalProps struct {
	Domain                         string `json:"domain"`
	CertIssuer                     string `json:"clusterCertIssuerName"`
	ClusterExternalSecretStoreName string `json:"clusterExternalSecretStoreName"`
}

func GetGlobal(scope packager.Construct) GlobalProps {
	globalValues := scope.GetContext("global").(string)
	return infrahelpers.FromYamlString[GlobalProps](globalValues)
}

func GetCertIssuer(scope packager.Construct) string {
	return GetGlobal(scope).CertIssuer
}

func GetCertIssuerAnnotation(scope packager.Construct) map[string]string {
	return map[string]string{"cert-manager.io/cluster-issuer": GetCertIssuer(scope)}
}

func SetGlobalContext(scope packager.Construct, props GlobalProps) {
	scope.SetContext("global", infrahelpers.ToYamlString(props))
}

func GetDomain(scope packager.Construct) string {
	return GetGlobal(scope).Domain
}
