package externalsecretsio


// use static token to authenticate with.
type ClusterSecretStoreSpecProviderKubernetesAuthToken struct {
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	BearerToken *ClusterSecretStoreSpecProviderKubernetesAuthTokenBearerToken `field:"optional" json:"bearerToken" yaml:"bearerToken"`
}

