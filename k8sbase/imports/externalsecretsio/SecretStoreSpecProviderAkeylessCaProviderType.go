package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type SecretStoreSpecProviderAkeylessCaProviderType string

const (
	// Secret.
	SecretStoreSpecProviderAkeylessCaProviderType_SECRET SecretStoreSpecProviderAkeylessCaProviderType = "SECRET"
	// ConfigMap.
	SecretStoreSpecProviderAkeylessCaProviderType_CONFIG_MAP SecretStoreSpecProviderAkeylessCaProviderType = "CONFIG_MAP"
)

