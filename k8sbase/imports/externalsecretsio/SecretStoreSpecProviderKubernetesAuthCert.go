package externalsecretsio


// has both clientCert and clientKey as secretKeySelector.
type SecretStoreSpecProviderKubernetesAuthCert struct {
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	ClientCert *SecretStoreSpecProviderKubernetesAuthCertClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	ClientKey *SecretStoreSpecProviderKubernetesAuthCertClientKey `field:"optional" json:"clientKey" yaml:"clientKey"`
}

