package externalsecretsio


type SecretStoreSpecProviderGcpsmAuthSecretRef struct {
	// The SecretAccessKey is used for authentication.
	SecretAccessKeySecretRef *SecretStoreSpecProviderGcpsmAuthSecretRefSecretAccessKeySecretRef `field:"optional" json:"secretAccessKeySecretRef" yaml:"secretAccessKeySecretRef"`
}

