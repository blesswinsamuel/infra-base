package externalsecretsio


// Version is the Vault KV secret engine version.
//
// This can be either "v1" or "v2". Version defaults to "v2".
type ClusterSecretStoreV1Beta1SpecProviderVaultVersion string

const (
	// v1.
	ClusterSecretStoreV1Beta1SpecProviderVaultVersion_V1 ClusterSecretStoreV1Beta1SpecProviderVaultVersion = "V1"
	// v2.
	ClusterSecretStoreV1Beta1SpecProviderVaultVersion_V2 ClusterSecretStoreV1Beta1SpecProviderVaultVersion = "V2"
)

