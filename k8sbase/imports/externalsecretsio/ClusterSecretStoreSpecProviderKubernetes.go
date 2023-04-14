// external-secretsio
package externalsecretsio


// Kubernetes configures this store to sync secrets using a Kubernetes cluster provider.
type ClusterSecretStoreSpecProviderKubernetes struct {
	// Auth configures how secret-manager authenticates with a Kubernetes instance.
	Auth *ClusterSecretStoreSpecProviderKubernetesAuth `field:"required" json:"auth" yaml:"auth"`
	// Remote namespace to fetch the secrets from.
	RemoteNamespace *string `field:"optional" json:"remoteNamespace" yaml:"remoteNamespace"`
	// configures the Kubernetes server Address.
	Server *ClusterSecretStoreSpecProviderKubernetesServer `field:"optional" json:"server" yaml:"server"`
}

