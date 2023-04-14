// external-secretsio
package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type ClusterSecretStoreV1Beta1SpecProviderKubernetesServerCaProviderType string

const (
	// Secret.
	ClusterSecretStoreV1Beta1SpecProviderKubernetesServerCaProviderType_SECRET ClusterSecretStoreV1Beta1SpecProviderKubernetesServerCaProviderType = "SECRET"
	// ConfigMap.
	ClusterSecretStoreV1Beta1SpecProviderKubernetesServerCaProviderType_CONFIG_MAP ClusterSecretStoreV1Beta1SpecProviderKubernetesServerCaProviderType = "CONFIG_MAP"
)

