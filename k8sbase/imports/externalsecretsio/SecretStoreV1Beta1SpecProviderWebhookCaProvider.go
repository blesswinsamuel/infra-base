// external-secretsio
package externalsecretsio


// The provider for the CA bundle to use to validate webhook server certificate.
type SecretStoreV1Beta1SpecProviderWebhookCaProvider struct {
	// The name of the object located at the provider type.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of provider to use such as "Secret", or "ConfigMap".
	Type SecretStoreV1Beta1SpecProviderWebhookCaProviderType `field:"required" json:"type" yaml:"type"`
	// The key the value inside of the provider type to use, only used with "Secret" type.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The namespace the Provider type is in.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

