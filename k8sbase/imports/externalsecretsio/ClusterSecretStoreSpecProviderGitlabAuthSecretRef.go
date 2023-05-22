package externalsecretsio


type ClusterSecretStoreSpecProviderGitlabAuthSecretRef struct {
	// AccessToken is used for authentication.
	AccessToken *ClusterSecretStoreSpecProviderGitlabAuthSecretRefAccessToken `field:"optional" json:"accessToken" yaml:"accessToken"`
}

