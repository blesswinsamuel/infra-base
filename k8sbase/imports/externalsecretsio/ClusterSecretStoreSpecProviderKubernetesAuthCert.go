// external-secretsio
package externalsecretsio


// has both clientCert and clientKey as secretKeySelector.
type ClusterSecretStoreSpecProviderKubernetesAuthCert struct {
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	ClientCert *ClusterSecretStoreSpecProviderKubernetesAuthCertClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	ClientKey *ClusterSecretStoreSpecProviderKubernetesAuthCertClientKey `field:"optional" json:"clientKey" yaml:"clientKey"`
}

