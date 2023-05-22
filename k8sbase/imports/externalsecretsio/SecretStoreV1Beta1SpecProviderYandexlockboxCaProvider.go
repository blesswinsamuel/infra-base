package externalsecretsio


// The provider for the CA bundle to use to validate Yandex.Cloud server certificate.
type SecretStoreV1Beta1SpecProviderYandexlockboxCaProvider struct {
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	CertSecretRef *SecretStoreV1Beta1SpecProviderYandexlockboxCaProviderCertSecretRef `field:"optional" json:"certSecretRef" yaml:"certSecretRef"`
}

