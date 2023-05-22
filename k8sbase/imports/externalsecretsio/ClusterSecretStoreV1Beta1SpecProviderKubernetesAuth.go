package externalsecretsio


// Auth configures how secret-manager authenticates with a Kubernetes instance.
type ClusterSecretStoreV1Beta1SpecProviderKubernetesAuth struct {
	// has both clientCert and clientKey as secretKeySelector.
	Cert *ClusterSecretStoreV1Beta1SpecProviderKubernetesAuthCert `field:"optional" json:"cert" yaml:"cert"`
	// points to a service account that should be used for authentication.
	ServiceAccount *ClusterSecretStoreV1Beta1SpecProviderKubernetesAuthServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// use static token to authenticate with.
	Token *ClusterSecretStoreV1Beta1SpecProviderKubernetesAuthToken `field:"optional" json:"token" yaml:"token"`
}

