package infraglobal

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/utils"
)

type GlobalProps struct {
	Domain                         string `yaml:"domain"`
	CertIssuer                     string `yaml:"clusterCertIssuerName"`
	ClusterExternalSecretStoreName string `yaml:"clusterExternalSecretStoreName"`
	InternetAuthType               string `yaml:"internetAuthType"`
}

func GetGlobal(scope constructs.Construct) GlobalProps {
	globalValues := scope.Node().TryGetContext(jsii.String("global")).(string)
	return utils.FromYamlString[GlobalProps](globalValues)
}

func GetCertIssuer(scope constructs.Construct) string {
	return GetGlobal(scope).CertIssuer
}

func GetCertIssuerAnnotation(scope constructs.Construct) map[string]string {
	return map[string]string{"cert-manager.io/cluster-issuer": GetCertIssuer(scope)}
}
