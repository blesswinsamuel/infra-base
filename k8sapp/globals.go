package k8sapp

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type Globals struct {
	DefaultSecretStoreName               string
	DefaultSecretStoreKind               string
	DefaultExternalSecretRefreshInterval string

	DefaultCertIssuerName string
	DefaultCertIssuerKind string

	CacheDir string
}

func GetGlobalContext(scope kubegogen.Scope) Globals {
	globalValues := scope.GetContext("globals").(string)
	return infrahelpers.FromYamlString[Globals](globalValues)
}

func SetGlobalContext(scope kubegogen.Scope, props Globals) {
	scope.SetContext("globals", infrahelpers.ToYamlString(props))
}
