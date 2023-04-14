// external-secretsio
package externalsecretsio


// Auth configures how secret-manager authenticates with a GitLab instance.
type SecretStoreV1Beta1SpecProviderGitlabAuth struct {
	SecretRef *SecretStoreV1Beta1SpecProviderGitlabAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

