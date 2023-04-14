// external-secretsio
package externalsecretsio


// Cert authenticates with TLS Certificates by passing client certificate, private key and ca certificate Cert authentication method.
type SecretStoreSpecProviderVaultAuthCert struct {
	// ClientCert is a certificate to authenticate using the Cert Vault authentication method.
	ClientCert *SecretStoreSpecProviderVaultAuthCertClientCert `field:"optional" json:"clientCert" yaml:"clientCert"`
	// SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method.
	SecretRef *SecretStoreSpecProviderVaultAuthCertSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

