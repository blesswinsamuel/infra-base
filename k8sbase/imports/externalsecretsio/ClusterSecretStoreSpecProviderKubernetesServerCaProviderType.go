// external-secretsio
package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type ClusterSecretStoreSpecProviderKubernetesServerCaProviderType string

const (
	// Secret.
	ClusterSecretStoreSpecProviderKubernetesServerCaProviderType_SECRET ClusterSecretStoreSpecProviderKubernetesServerCaProviderType = "SECRET"
	// ConfigMap.
	ClusterSecretStoreSpecProviderKubernetesServerCaProviderType_CONFIG_MAP ClusterSecretStoreSpecProviderKubernetesServerCaProviderType = "CONFIG_MAP"
)

