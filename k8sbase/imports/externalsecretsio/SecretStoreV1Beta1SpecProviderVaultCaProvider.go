package externalsecretsio


// The provider for the CA bundle to use to validate Vault server certificate.
type SecretStoreV1Beta1SpecProviderVaultCaProvider struct {
	// The name of the object located at the provider type.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of provider to use such as "Secret", or "ConfigMap".
	Type SecretStoreV1Beta1SpecProviderVaultCaProviderType `field:"required" json:"type" yaml:"type"`
	// The key where the CA certificate can be found in the Secret or ConfigMap.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The namespace the Provider type is in.
	//
	// Can only be defined when used in a ClusterSecretStore.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

