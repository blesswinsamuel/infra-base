package externalsecretsio


// see: https://external-secrets.io/v0.4.1/spec/#external-secrets.io/v1alpha1.CAProvider.
type SecretStoreSpecProviderKubernetesServerCaProvider struct {
	// The name of the object located at the provider type.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of provider to use such as "Secret", or "ConfigMap".
	Type SecretStoreSpecProviderKubernetesServerCaProviderType `field:"required" json:"type" yaml:"type"`
	// The key the value inside of the provider type to use, only used with "Secret" type.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The namespace the Provider type is in.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

