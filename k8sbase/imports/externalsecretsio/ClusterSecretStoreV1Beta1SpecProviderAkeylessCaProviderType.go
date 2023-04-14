// external-secretsio
package externalsecretsio


// The type of provider to use such as "Secret", or "ConfigMap".
type ClusterSecretStoreV1Beta1SpecProviderAkeylessCaProviderType string

const (
	// Secret.
	ClusterSecretStoreV1Beta1SpecProviderAkeylessCaProviderType_SECRET ClusterSecretStoreV1Beta1SpecProviderAkeylessCaProviderType = "SECRET"
	// ConfigMap.
	ClusterSecretStoreV1Beta1SpecProviderAkeylessCaProviderType_CONFIG_MAP ClusterSecretStoreV1Beta1SpecProviderAkeylessCaProviderType = "CONFIG_MAP"
)

