// external-secretsio
package externalsecretsio


// Kubernetes configures this store to sync secrets using a Kubernetes cluster provider.
type SecretStoreV1Beta1SpecProviderKubernetes struct {
	// Auth configures how secret-manager authenticates with a Kubernetes instance.
	Auth *SecretStoreV1Beta1SpecProviderKubernetesAuth `field:"required" json:"auth" yaml:"auth"`
	// Remote namespace to fetch the secrets from.
	RemoteNamespace *string `field:"optional" json:"remoteNamespace" yaml:"remoteNamespace"`
	// configures the Kubernetes server Address.
	Server *SecretStoreV1Beta1SpecProviderKubernetesServer `field:"optional" json:"server" yaml:"server"`
}

