package externalsecretsio


// Version is the Vault KV secret engine version.
//
// This can be either "v1" or "v2". Version defaults to "v2".
type ClusterSecretStoreSpecProviderVaultVersion string

const (
	// v1.
	ClusterSecretStoreSpecProviderVaultVersion_V1 ClusterSecretStoreSpecProviderVaultVersion = "V1"
	// v2.
	ClusterSecretStoreSpecProviderVaultVersion_V2 ClusterSecretStoreSpecProviderVaultVersion = "V2"
)

