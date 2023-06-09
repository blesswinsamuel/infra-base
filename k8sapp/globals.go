package k8sapp

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
)

type Globals struct {
	DefaultSecretStoreName               string
	DefaultSecretStoreKind               string
	DefaultExternalSecretRefreshInterval string

	DefaultCertIssuerName string
	DefaultCertIssuerKind string

	CacheDir string
}

func GetGlobalContext(scope constructs.Construct) Globals {
	globalValues := scope.Node().TryGetContext(jsii.String("globals")).(string)
	return infrahelpers.FromYamlString[Globals](globalValues)
}

func SetGlobalContext(scope constructs.Construct, props Globals) {
	scope.Node().SetContext(jsii.String("globals"), jsii.String(infrahelpers.ToYamlString(props)))
}
