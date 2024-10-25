package k8sapp

import "github.com/blesswinsamuel/kgen"

type ValuesGlobalCert struct {
	CertIssuerName string `json:"certIssuerName"`
	CertIssuerKind string `json:"certIssuerKind"`
}

type ValuesGlobalIngress struct {
	DisableTls bool `json:"disableTls"`
}

type ValuesGlobalExternalSecret struct {
	SecretsProvider string `json:"secretsProvider"`

	SecretStoreName string `json:"secretStoreName"`
	SecretStoreKind string `json:"secretStoreKind"`
	RefreshInterval string `json:"refreshInterval"`
}

type ValuesGlobal struct {
	Domain      string `json:"domain"`
	ClusterName string `json:"clusterName"`

	Cert           ValuesGlobalCert             `json:"cert"`
	Ingress        ValuesGlobalIngress          `json:"ingress"`
	ExternalSecret ValuesGlobalExternalSecret   `json:"externalSecret"`
	AppRefs        map[string]NameNamespacePort `json:"appRefs"`
}

var defaultValuesGlobal = ValuesGlobal{
	Domain:      "",
	ClusterName: "",
	ExternalSecret: ValuesGlobalExternalSecret{
		SecretsProvider: "doppler",
		SecretStoreName: "secretstore",
		SecretStoreKind: "ClusterSecretStore",
		RefreshInterval: "1m",
	},
	Cert: ValuesGlobalCert{
		CertIssuerName: "letsencrypt-prod",
		CertIssuerKind: "ClusterIssuer",
	},
	Ingress: ValuesGlobalIngress{
		DisableTls: false,
	},
}

var globalContextKey = kgen.GenerateContextKey()

func GetGlobals(scope kgen.Scope) ValuesGlobal {
	return scope.GetContext(globalContextKey).(ValuesGlobal)
}

func SetGlobals(scope kgen.Scope, props ValuesGlobal) {
	scope.SetContext(globalContextKey, props)
}

func GetDomain(scope kgen.Scope) string {
	return GetGlobals(scope).Domain
}
