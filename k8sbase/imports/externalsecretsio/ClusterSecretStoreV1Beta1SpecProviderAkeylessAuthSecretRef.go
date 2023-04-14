// external-secretsio
package externalsecretsio


// Auth configures how the operator authenticates with Akeyless.
type ClusterSecretStoreV1Beta1SpecProviderAkeylessAuthSecretRef struct {
	// Kubernetes authenticates with Akeyless by passing the ServiceAccount token stored in the named Secret resource.
	KubernetesAuth *ClusterSecretStoreV1Beta1SpecProviderAkeylessAuthSecretRefKubernetesAuth `field:"optional" json:"kubernetesAuth" yaml:"kubernetesAuth"`
	// Reference to a Secret that contains the details to authenticate with Akeyless.
	SecretRef *ClusterSecretStoreV1Beta1SpecProviderAkeylessAuthSecretRefSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

