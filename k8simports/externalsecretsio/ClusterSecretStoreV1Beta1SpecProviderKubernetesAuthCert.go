package externalsecretsio


// has both clientCert and clientKey as secretKeySelector.
type ClusterSecretStoreV1Beta1SpecProviderKubernetesAuthCert struct {
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	ClientCert *ClusterSecretStoreV1Beta1SpecProviderKubernetesAuthCertClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	ClientKey *ClusterSecretStoreV1Beta1SpecProviderKubernetesAuthCertClientKey `field:"optional" json:"clientKey" yaml:"clientKey"`
}

