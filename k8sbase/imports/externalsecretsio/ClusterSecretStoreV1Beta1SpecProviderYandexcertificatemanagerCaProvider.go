package externalsecretsio


// The provider for the CA bundle to use to validate Yandex.Cloud server certificate.
type ClusterSecretStoreV1Beta1SpecProviderYandexcertificatemanagerCaProvider struct {
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	CertSecretRef *ClusterSecretStoreV1Beta1SpecProviderYandexcertificatemanagerCaProviderCertSecretRef `field:"optional" json:"certSecretRef" yaml:"certSecretRef"`
}

