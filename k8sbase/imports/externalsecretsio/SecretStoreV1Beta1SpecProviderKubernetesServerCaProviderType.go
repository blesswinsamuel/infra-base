package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type SecretStoreV1Beta1SpecProviderKubernetesServerCaProviderType string

const (
	// Secret.
	SecretStoreV1Beta1SpecProviderKubernetesServerCaProviderType_SECRET SecretStoreV1Beta1SpecProviderKubernetesServerCaProviderType = "SECRET"
	// ConfigMap.
	SecretStoreV1Beta1SpecProviderKubernetesServerCaProviderType_CONFIG_MAP SecretStoreV1Beta1SpecProviderKubernetesServerCaProviderType = "CONFIG_MAP"
)

