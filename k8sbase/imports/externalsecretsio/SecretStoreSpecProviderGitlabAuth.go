// external-secretsio
package externalsecretsio


// Auth configures how secret-manager authenticates with a GitLab instance.
type SecretStoreSpecProviderGitlabAuth struct {
	SecretRef *SecretStoreSpecProviderGitlabAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

