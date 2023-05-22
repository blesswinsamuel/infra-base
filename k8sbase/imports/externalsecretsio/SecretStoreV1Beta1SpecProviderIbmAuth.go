package externalsecretsio


// Auth configures how secret-manager authenticates with the IBM secrets manager.
type SecretStoreV1Beta1SpecProviderIbmAuth struct {
	// IBM Container-based auth with IAM Trusted Profile.
	ContainerAuth *SecretStoreV1Beta1SpecProviderIbmAuthContainerAuth `field:"optional" json:"containerAuth" yaml:"containerAuth"`
	SecretRef *SecretStoreV1Beta1SpecProviderIbmAuthSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

