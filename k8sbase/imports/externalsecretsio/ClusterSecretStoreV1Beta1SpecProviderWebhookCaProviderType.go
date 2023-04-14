// external-secretsio
package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type ClusterSecretStoreV1Beta1SpecProviderWebhookCaProviderType string

const (
	// Secret.
	ClusterSecretStoreV1Beta1SpecProviderWebhookCaProviderType_SECRET ClusterSecretStoreV1Beta1SpecProviderWebhookCaProviderType = "SECRET"
	// ConfigMap.
	ClusterSecretStoreV1Beta1SpecProviderWebhookCaProviderType_CONFIG_MAP ClusterSecretStoreV1Beta1SpecProviderWebhookCaProviderType = "CONFIG_MAP"
)

