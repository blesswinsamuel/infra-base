// external-secretsio
package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type SecretStoreV1Beta1SpecProviderAkeylessCaProviderType string

const (
	// Secret.
	SecretStoreV1Beta1SpecProviderAkeylessCaProviderType_SECRET SecretStoreV1Beta1SpecProviderAkeylessCaProviderType = "SECRET"
	// ConfigMap.
	SecretStoreV1Beta1SpecProviderAkeylessCaProviderType_CONFIG_MAP SecretStoreV1Beta1SpecProviderAkeylessCaProviderType = "CONFIG_MAP"
)

