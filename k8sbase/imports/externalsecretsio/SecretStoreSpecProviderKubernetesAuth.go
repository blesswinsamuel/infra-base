// external-secretsio
package externalsecretsio


// Auth configures how secret-manager authenticates with a Kubernetes instance.
type SecretStoreSpecProviderKubernetesAuth struct {
	// has both clientCert and clientKey as secretKeySelector.
	Cert *SecretStoreSpecProviderKubernetesAuthCert `field:"optional" json:"cert" yaml:"cert"`
	// points to a service account that should be used for authentication.
	ServiceAccount *SecretStoreSpecProviderKubernetesAuthServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// use static token to authenticate with.
	Token *SecretStoreSpecProviderKubernetesAuthToken `field:"optional" json:"token" yaml:"token"`
}

