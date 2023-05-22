package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type ClusterSecretStoreV1Beta1SpecProviderVaultCaProviderType string

const (
	// Secret.
	ClusterSecretStoreV1Beta1SpecProviderVaultCaProviderType_SECRET ClusterSecretStoreV1Beta1SpecProviderVaultCaProviderType = "SECRET"
	// ConfigMap.
	ClusterSecretStoreV1Beta1SpecProviderVaultCaProviderType_CONFIG_MAP ClusterSecretStoreV1Beta1SpecProviderVaultCaProviderType = "CONFIG_MAP"
)

