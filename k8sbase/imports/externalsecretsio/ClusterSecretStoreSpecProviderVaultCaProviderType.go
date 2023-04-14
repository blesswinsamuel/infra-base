// external-secretsio
package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type ClusterSecretStoreSpecProviderVaultCaProviderType string

const (
	// Secret.
	ClusterSecretStoreSpecProviderVaultCaProviderType_SECRET ClusterSecretStoreSpecProviderVaultCaProviderType = "SECRET"
	// ConfigMap.
	ClusterSecretStoreSpecProviderVaultCaProviderType_CONFIG_MAP ClusterSecretStoreSpecProviderVaultCaProviderType = "CONFIG_MAP"
)

