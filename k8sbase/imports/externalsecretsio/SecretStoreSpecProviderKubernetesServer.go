package externalsecretsio


// configures the Kubernetes server Address.
type SecretStoreSpecProviderKubernetesServer struct {
	// CABundle is a base64-encoded CA certificate.
	CaBundle *string `field:"optional" json:"caBundle" yaml:"caBundle"`
	// see: https://external-secrets.io/v0.4.1/spec/#external-secrets.io/v1alpha1.CAProvider.
	CaProvider *SecretStoreSpecProviderKubernetesServerCaProvider `field:"optional" json:"caProvider" yaml:"caProvider"`
	// configures the Kubernetes server Address.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

