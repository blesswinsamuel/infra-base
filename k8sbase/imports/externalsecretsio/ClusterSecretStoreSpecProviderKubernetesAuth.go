package externalsecretsio


// Auth configures how secret-manager authenticates with a Kubernetes instance.
type ClusterSecretStoreSpecProviderKubernetesAuth struct {
	// has both clientCert and clientKey as secretKeySelector.
	Cert *ClusterSecretStoreSpecProviderKubernetesAuthCert `field:"optional" json:"cert" yaml:"cert"`
	// points to a service account that should be used for authentication.
	ServiceAccount *ClusterSecretStoreSpecProviderKubernetesAuthServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// use static token to authenticate with.
	Token *ClusterSecretStoreSpecProviderKubernetesAuthToken `field:"optional" json:"token" yaml:"token"`
}

