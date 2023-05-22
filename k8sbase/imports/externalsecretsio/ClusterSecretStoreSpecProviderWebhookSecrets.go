package externalsecretsio


type ClusterSecretStoreSpecProviderWebhookSecrets struct {
	// Name of this secret in templates.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Secret ref to fill in credentials.
	SecretRef *ClusterSecretStoreSpecProviderWebhookSecretsSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

