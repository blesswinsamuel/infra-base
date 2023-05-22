package externalsecretsio


// Auth configures how secret-manager authenticates with the IBM secrets manager.
type ClusterSecretStoreV1Beta1SpecProviderIbmAuth struct {
	// IBM Container-based auth with IAM Trusted Profile.
	ContainerAuth *ClusterSecretStoreV1Beta1SpecProviderIbmAuthContainerAuth `field:"optional" json:"containerAuth" yaml:"containerAuth"`
	SecretRef *ClusterSecretStoreV1Beta1SpecProviderIbmAuthSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

