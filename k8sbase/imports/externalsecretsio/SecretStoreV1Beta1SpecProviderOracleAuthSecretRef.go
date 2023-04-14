// external-secretsio
package externalsecretsio


// SecretRef to pass through sensitive information.
type SecretStoreV1Beta1SpecProviderOracleAuthSecretRef struct {
	// Fingerprint is the fingerprint of the API private key.
	Fingerprint *SecretStoreV1Beta1SpecProviderOracleAuthSecretRefFingerprint `field:"required" json:"fingerprint" yaml:"fingerprint"`
	// PrivateKey is the user's API Signing Key in PEM format, used for authentication.
	Privatekey *SecretStoreV1Beta1SpecProviderOracleAuthSecretRefPrivatekey `field:"required" json:"privatekey" yaml:"privatekey"`
}

