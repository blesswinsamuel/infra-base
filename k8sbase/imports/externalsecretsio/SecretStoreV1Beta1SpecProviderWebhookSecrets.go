package externalsecretsio


type SecretStoreV1Beta1SpecProviderWebhookSecrets struct {
	// Name of this secret in templates.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Secret ref to fill in credentials.
	SecretRef *SecretStoreV1Beta1SpecProviderWebhookSecretsSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

