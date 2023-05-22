package externalsecretsio


// Auth defines the information necessary to authenticate against Yandex Lockbox.
type ClusterSecretStoreSpecProviderYandexlockboxAuth struct {
	// The authorized key used for authentication.
	AuthorizedKeySecretRef *ClusterSecretStoreSpecProviderYandexlockboxAuthAuthorizedKeySecretRef `field:"optional" json:"authorizedKeySecretRef" yaml:"authorizedKeySecretRef"`
}

