package externalsecretsio


// Auth defines the information necessary to authenticate against Yandex Lockbox.
type SecretStoreSpecProviderYandexlockboxAuth struct {
	// The authorized key used for authentication.
	AuthorizedKeySecretRef *SecretStoreSpecProviderYandexlockboxAuthAuthorizedKeySecretRef `field:"optional" json:"authorizedKeySecretRef" yaml:"authorizedKeySecretRef"`
}

