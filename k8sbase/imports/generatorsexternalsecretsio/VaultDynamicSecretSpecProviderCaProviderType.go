// generatorsexternal-secretsio
package generatorsexternalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type VaultDynamicSecretSpecProviderCaProviderType string

const (
	// Secret.
	VaultDynamicSecretSpecProviderCaProviderType_SECRET VaultDynamicSecretSpecProviderCaProviderType = "SECRET"
	// ConfigMap.
	VaultDynamicSecretSpecProviderCaProviderType_CONFIG_MAP VaultDynamicSecretSpecProviderCaProviderType = "CONFIG_MAP"
)

