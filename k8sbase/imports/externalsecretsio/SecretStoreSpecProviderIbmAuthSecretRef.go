package externalsecretsio


type SecretStoreSpecProviderIbmAuthSecretRef struct {
	// The SecretAccessKey is used for authentication.
	SecretApiKeySecretRef *SecretStoreSpecProviderIbmAuthSecretRefSecretApiKeySecretRef `field:"optional" json:"secretApiKeySecretRef" yaml:"secretApiKeySecretRef"`
}

