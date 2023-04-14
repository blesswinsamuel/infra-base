// external-secretsio
package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type ClusterSecretStoreSpecProviderWebhookCaProviderType string

const (
	// Secret.
	ClusterSecretStoreSpecProviderWebhookCaProviderType_SECRET ClusterSecretStoreSpecProviderWebhookCaProviderType = "SECRET"
	// ConfigMap.
	ClusterSecretStoreSpecProviderWebhookCaProviderType_CONFIG_MAP ClusterSecretStoreSpecProviderWebhookCaProviderType = "CONFIG_MAP"
)

