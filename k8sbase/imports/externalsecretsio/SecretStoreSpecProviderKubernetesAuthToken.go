package externalsecretsio


// use static token to authenticate with.
type SecretStoreSpecProviderKubernetesAuthToken struct {
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	BearerToken *SecretStoreSpecProviderKubernetesAuthTokenBearerToken `field:"optional" json:"bearerToken" yaml:"bearerToken"`
}

