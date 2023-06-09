package externalsecretsio


// use static token to authenticate with.
type ClusterSecretStoreV1Beta1SpecProviderKubernetesAuthToken struct {
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	BearerToken *ClusterSecretStoreV1Beta1SpecProviderKubernetesAuthTokenBearerToken `field:"optional" json:"bearerToken" yaml:"bearerToken"`
}

