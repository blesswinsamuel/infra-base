package externalsecretsio


// Auth configures how the operator authenticates with Akeyless.
type ClusterSecretStoreSpecProviderAkeylessAuthSecretRef struct {
	// Kubernetes authenticates with Akeyless by passing the ServiceAccount token stored in the named Secret resource.
	KubernetesAuth *ClusterSecretStoreSpecProviderAkeylessAuthSecretRefKubernetesAuth `field:"optional" json:"kubernetesAuth" yaml:"kubernetesAuth"`
	// Reference to a Secret that contains the details to authenticate with Akeyless.
	SecretRef *ClusterSecretStoreSpecProviderAkeylessAuthSecretRefSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

