package externalsecretsio


// Auth configures how the operator authenticates with Akeyless.
type SecretStoreV1Beta1SpecProviderAkeylessAuthSecretRef struct {
	// Kubernetes authenticates with Akeyless by passing the ServiceAccount token stored in the named Secret resource.
	KubernetesAuth *SecretStoreV1Beta1SpecProviderAkeylessAuthSecretRefKubernetesAuth `field:"optional" json:"kubernetesAuth" yaml:"kubernetesAuth"`
	// Reference to a Secret that contains the details to authenticate with Akeyless.
	SecretRef *SecretStoreV1Beta1SpecProviderAkeylessAuthSecretRefSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

