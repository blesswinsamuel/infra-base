package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type SecretStoreV1Beta1SpecProviderVaultCaProviderType string

const (
	// Secret.
	SecretStoreV1Beta1SpecProviderVaultCaProviderType_SECRET SecretStoreV1Beta1SpecProviderVaultCaProviderType = "SECRET"
	// ConfigMap.
	SecretStoreV1Beta1SpecProviderVaultCaProviderType_CONFIG_MAP SecretStoreV1Beta1SpecProviderVaultCaProviderType = "CONFIG_MAP"
)

