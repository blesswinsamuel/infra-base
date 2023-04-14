// external-secretsio
package externalsecretsio


// has both clientCert and clientKey as secretKeySelector.
type SecretStoreV1Beta1SpecProviderKubernetesAuthCert struct {
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	ClientCert *SecretStoreV1Beta1SpecProviderKubernetesAuthCertClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	ClientKey *SecretStoreV1Beta1SpecProviderKubernetesAuthCertClientKey `field:"optional" json:"clientKey" yaml:"clientKey"`
}

