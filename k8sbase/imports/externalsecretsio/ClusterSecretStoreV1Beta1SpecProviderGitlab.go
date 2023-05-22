package externalsecretsio


// Gitlab configures this store to sync secrets using Gitlab Variables provider.
type ClusterSecretStoreV1Beta1SpecProviderGitlab struct {
	// Auth configures how secret-manager authenticates with a GitLab instance.
	Auth *ClusterSecretStoreV1Beta1SpecProviderGitlabAuth `field:"required" json:"auth" yaml:"auth"`
	// Environment environment_scope of gitlab CI/CD variables (Please see https://docs.gitlab.com/ee/ci/environments/#create-a-static-environment on how to create environments).
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// GroupIDs specify, which gitlab groups to pull secrets from.
	//
	// Group secrets are read from left to right followed by the project variables.
	GroupIDs *[]*string `field:"optional" json:"groupIDs" yaml:"groupIDs"`
	// InheritFromGroups specifies whether parent groups should be discovered and checked for secrets.
	InheritFromGroups *bool `field:"optional" json:"inheritFromGroups" yaml:"inheritFromGroups"`
	// ProjectID specifies a project where secrets are located.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// URL configures the GitLab instance URL.
	//
	// Defaults to https://gitlab.com/.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

