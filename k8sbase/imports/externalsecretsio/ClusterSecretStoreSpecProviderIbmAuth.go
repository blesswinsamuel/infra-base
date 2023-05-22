package externalsecretsio


// Auth configures how secret-manager authenticates with the IBM secrets manager.
type ClusterSecretStoreSpecProviderIbmAuth struct {
	SecretRef *ClusterSecretStoreSpecProviderIbmAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

