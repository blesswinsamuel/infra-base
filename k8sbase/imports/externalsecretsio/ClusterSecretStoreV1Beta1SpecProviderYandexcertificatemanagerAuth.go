// external-secretsio
package externalsecretsio


// Auth defines the information necessary to authenticate against Yandex Certificate Manager.
type ClusterSecretStoreV1Beta1SpecProviderYandexcertificatemanagerAuth struct {
	// The authorized key used for authentication.
	AuthorizedKeySecretRef *ClusterSecretStoreV1Beta1SpecProviderYandexcertificatemanagerAuthAuthorizedKeySecretRef `field:"optional" json:"authorizedKeySecretRef" yaml:"authorizedKeySecretRef"`
}

