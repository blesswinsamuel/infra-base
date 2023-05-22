package externalsecretsio


// Auth configures how secret-manager authenticates with a GitLab instance.
type ClusterSecretStoreV1Beta1SpecProviderGitlabAuth struct {
	SecretRef *ClusterSecretStoreV1Beta1SpecProviderGitlabAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

