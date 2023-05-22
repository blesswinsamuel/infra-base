package externalsecretsio


type ClusterSecretStoreV1Beta1SpecProviderIbmAuthSecretRef struct {
	// The SecretAccessKey is used for authentication.
	SecretApiKeySecretRef *ClusterSecretStoreV1Beta1SpecProviderIbmAuthSecretRefSecretApiKeySecretRef `field:"optional" json:"secretApiKeySecretRef" yaml:"secretApiKeySecretRef"`
}

