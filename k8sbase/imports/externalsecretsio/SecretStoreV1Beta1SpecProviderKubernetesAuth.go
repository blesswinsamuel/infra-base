package externalsecretsio


// Auth configures how secret-manager authenticates with a Kubernetes instance.
type SecretStoreV1Beta1SpecProviderKubernetesAuth struct {
	// has both clientCert and clientKey as secretKeySelector.
	Cert *SecretStoreV1Beta1SpecProviderKubernetesAuthCert `field:"optional" json:"cert" yaml:"cert"`
	// points to a service account that should be used for authentication.
	ServiceAccount *SecretStoreV1Beta1SpecProviderKubernetesAuthServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// use static token to authenticate with.
	Token *SecretStoreV1Beta1SpecProviderKubernetesAuthToken `field:"optional" json:"token" yaml:"token"`
}

