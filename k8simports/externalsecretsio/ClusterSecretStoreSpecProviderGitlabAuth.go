package externalsecretsio


// Auth configures how secret-manager authenticates with a GitLab instance.
type ClusterSecretStoreSpecProviderGitlabAuth struct {
	SecretRef *ClusterSecretStoreSpecProviderGitlabAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

