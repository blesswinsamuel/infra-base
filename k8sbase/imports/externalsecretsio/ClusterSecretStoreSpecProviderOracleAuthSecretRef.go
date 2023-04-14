// external-secretsio
package externalsecretsio


// SecretRef to pass through sensitive information.
type ClusterSecretStoreSpecProviderOracleAuthSecretRef struct {
	// Fingerprint is the fingerprint of the API private key.
	Fingerprint *ClusterSecretStoreSpecProviderOracleAuthSecretRefFingerprint `field:"required" json:"fingerprint" yaml:"fingerprint"`
	// PrivateKey is the user's API Signing Key in PEM format, used for authentication.
	Privatekey *ClusterSecretStoreSpecProviderOracleAuthSecretRefPrivatekey `field:"required" json:"privatekey" yaml:"privatekey"`
}

