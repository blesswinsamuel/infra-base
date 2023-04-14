// external-secretsio
package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type SecretStoreSpecProviderVaultCaProviderType string

const (
	// Secret.
	SecretStoreSpecProviderVaultCaProviderType_SECRET SecretStoreSpecProviderVaultCaProviderType = "SECRET"
	// ConfigMap.
	SecretStoreSpecProviderVaultCaProviderType_CONFIG_MAP SecretStoreSpecProviderVaultCaProviderType = "CONFIG_MAP"
)

