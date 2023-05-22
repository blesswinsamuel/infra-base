package externalsecretsio


type SecretStoreV1Beta1SpecProviderIbmAuthSecretRef struct {
	// The SecretAccessKey is used for authentication.
	SecretApiKeySecretRef *SecretStoreV1Beta1SpecProviderIbmAuthSecretRefSecretApiKeySecretRef `field:"optional" json:"secretApiKeySecretRef" yaml:"secretApiKeySecretRef"`
}

