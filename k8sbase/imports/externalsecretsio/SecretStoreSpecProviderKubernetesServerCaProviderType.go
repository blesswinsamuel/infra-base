// external-secretsio
package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type SecretStoreSpecProviderKubernetesServerCaProviderType string

const (
	// Secret.
	SecretStoreSpecProviderKubernetesServerCaProviderType_SECRET SecretStoreSpecProviderKubernetesServerCaProviderType = "SECRET"
	// ConfigMap.
	SecretStoreSpecProviderKubernetesServerCaProviderType_CONFIG_MAP SecretStoreSpecProviderKubernetesServerCaProviderType = "CONFIG_MAP"
)

