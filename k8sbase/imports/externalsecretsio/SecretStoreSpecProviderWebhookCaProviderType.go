// external-secretsio
package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type SecretStoreSpecProviderWebhookCaProviderType string

const (
	// Secret.
	SecretStoreSpecProviderWebhookCaProviderType_SECRET SecretStoreSpecProviderWebhookCaProviderType = "SECRET"
	// ConfigMap.
	SecretStoreSpecProviderWebhookCaProviderType_CONFIG_MAP SecretStoreSpecProviderWebhookCaProviderType = "CONFIG_MAP"
)

