package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type SecretStoreV1Beta1SpecProviderWebhookCaProviderType string

const (
	// Secret.
	SecretStoreV1Beta1SpecProviderWebhookCaProviderType_SECRET SecretStoreV1Beta1SpecProviderWebhookCaProviderType = "SECRET"
	// ConfigMap.
	SecretStoreV1Beta1SpecProviderWebhookCaProviderType_CONFIG_MAP SecretStoreV1Beta1SpecProviderWebhookCaProviderType = "CONFIG_MAP"
)

