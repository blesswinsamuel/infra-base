// external-secretsio
package externalsecretsio


// Vault configures this store to sync secrets using Hashi provider.
type SecretStoreV1Beta1SpecProviderVault struct {
	// Auth configures how secret-manager authenticates with the Vault server.
	Auth *SecretStoreV1Beta1SpecProviderVaultAuth `field:"required" json:"auth" yaml:"auth"`
	// Server is the connection address for the Vault server, e.g: "https://vault.example.com:8200".
	Server *string `field:"required" json:"server" yaml:"server"`
	// PEM encoded CA bundle used to validate Vault server certificate.
	//
	// Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.
	CaBundle *string `field:"optional" json:"caBundle" yaml:"caBundle"`
	// The provider for the CA bundle to use to validate Vault server certificate.
	CaProvider *SecretStoreV1Beta1SpecProviderVaultCaProvider `field:"optional" json:"caProvider" yaml:"caProvider"`
	// ForwardInconsistent tells Vault to forward read-after-write requests to the Vault leader instead of simply retrying within a loop.
	//
	// This can increase performance if the option is enabled serverside. https://www.vaultproject.io/docs/configuration/replication#allow_forwarding_via_header
	ForwardInconsistent *bool `field:"optional" json:"forwardInconsistent" yaml:"forwardInconsistent"`
	// Name of the vault namespace.
	//
	// Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: "ns1". More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Path is the mount path of the Vault KV backend endpoint, e.g: "secret". The v2 KV secret engine version specific "/data" path suffix for fetching secrets from Vault is optional and will be appended if not present in specified path.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// ReadYourWrites ensures isolated read-after-write semantics by providing discovered cluster replication states in each request.
	//
	// More information about eventual consistency in Vault can be found here https://www.vaultproject.io/docs/enterprise/consistency
	ReadYourWrites *bool `field:"optional" json:"readYourWrites" yaml:"readYourWrites"`
	// Version is the Vault KV secret engine version.
	//
	// This can be either "v1" or "v2". Version defaults to "v2".
	Version SecretStoreV1Beta1SpecProviderVaultVersion `field:"optional" json:"version" yaml:"version"`
}

