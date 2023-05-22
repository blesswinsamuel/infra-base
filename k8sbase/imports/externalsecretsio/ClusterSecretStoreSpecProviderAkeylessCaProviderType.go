package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type ClusterSecretStoreSpecProviderAkeylessCaProviderType string

const (
	// Secret.
	ClusterSecretStoreSpecProviderAkeylessCaProviderType_SECRET ClusterSecretStoreSpecProviderAkeylessCaProviderType = "SECRET"
	// ConfigMap.
	ClusterSecretStoreSpecProviderAkeylessCaProviderType_CONFIG_MAP ClusterSecretStoreSpecProviderAkeylessCaProviderType = "CONFIG_MAP"
)

