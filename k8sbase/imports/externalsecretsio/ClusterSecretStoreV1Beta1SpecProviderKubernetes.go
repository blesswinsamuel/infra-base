package externalsecretsio


// Kubernetes configures this store to sync secrets using a Kubernetes cluster provider.
type ClusterSecretStoreV1Beta1SpecProviderKubernetes struct {
	// Auth configures how secret-manager authenticates with a Kubernetes instance.
	Auth *ClusterSecretStoreV1Beta1SpecProviderKubernetesAuth `field:"required" json:"auth" yaml:"auth"`
	// Remote namespace to fetch the secrets from.
	RemoteNamespace *string `field:"optional" json:"remoteNamespace" yaml:"remoteNamespace"`
	// configures the Kubernetes server Address.
	Server *ClusterSecretStoreV1Beta1SpecProviderKubernetesServer `field:"optional" json:"server" yaml:"server"`
}

