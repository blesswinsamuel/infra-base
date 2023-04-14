// external-secretsio
package externalsecretsio


// Gitlab configures this store to sync secrets using Gitlab Variables provider.
type SecretStoreSpecProviderGitlab struct {
	// Auth configures how secret-manager authenticates with a GitLab instance.
	Auth *SecretStoreSpecProviderGitlabAuth `field:"required" json:"auth" yaml:"auth"`
	// ProjectID specifies a project where secrets are located.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// URL configures the GitLab instance URL.
	//
	// Defaults to https://gitlab.com/.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

