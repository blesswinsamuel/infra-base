package externalsecretsio


// see: https://external-secrets.io/v0.4.1/spec/#external-secrets.io/v1alpha1.CAProvider.
type SecretStoreV1Beta1SpecProviderKubernetesServerCaProvider struct {
	// The name of the object located at the provider type.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of provider to use such as "Secret", or "ConfigMap".
	Type SecretStoreV1Beta1SpecProviderKubernetesServerCaProviderType `field:"required" json:"type" yaml:"type"`
	// The key where the CA certificate can be found in the Secret or ConfigMap.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The namespace the Provider type is in.
	//
	// Can only be defined when used in a ClusterSecretStore.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

