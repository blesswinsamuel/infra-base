package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
)

type GlobalProps struct {
	Domain                         string `json:"domain"`
	CertIssuer                     string `json:"clusterCertIssuerName"`
	ClusterExternalSecretStoreName string `json:"clusterExternalSecretStoreName"`
	InternetAuthType               string `json:"internetAuthType"`
}

func GetGlobal(scope constructs.Construct) GlobalProps {
	globalValues := scope.Node().TryGetContext(jsii.String("global")).(string)
	return infrahelpers.FromYamlString[GlobalProps](globalValues)
}

func GetCertIssuer(scope constructs.Construct) string {
	return GetGlobal(scope).CertIssuer
}

func GetCertIssuerAnnotation(scope constructs.Construct) map[string]string {
	return map[string]string{"cert-manager.io/cluster-issuer": GetCertIssuer(scope)}
}

func SetGlobalContext(scope constructs.Construct, props GlobalProps) {
	scope.Node().SetContext(jsii.String("global"), infrahelpers.ToYamlString(props))
}

func GetDomain(scope constructs.Construct) string {
	return GetGlobal(scope).Domain
}
