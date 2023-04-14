// external-secretsio
package externalsecretsio


// Version is the Vault KV secret engine version.
//
// This can be either "v1" or "v2". Version defaults to "v2".
type SecretStoreSpecProviderVaultVersion string

const (
	// v1.
	SecretStoreSpecProviderVaultVersion_V1 SecretStoreSpecProviderVaultVersion = "V1"
	// v2.
	SecretStoreSpecProviderVaultVersion_V2 SecretStoreSpecProviderVaultVersion = "V2"
)

