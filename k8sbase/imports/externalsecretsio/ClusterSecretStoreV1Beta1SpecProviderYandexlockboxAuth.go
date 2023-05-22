package externalsecretsio


// Auth defines the information necessary to authenticate against Yandex Lockbox.
type ClusterSecretStoreV1Beta1SpecProviderYandexlockboxAuth struct {
	// The authorized key used for authentication.
	AuthorizedKeySecretRef *ClusterSecretStoreV1Beta1SpecProviderYandexlockboxAuthAuthorizedKeySecretRef `field:"optional" json:"authorizedKeySecretRef" yaml:"authorizedKeySecretRef"`
}

