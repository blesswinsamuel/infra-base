// external-secretsio
package externalsecretsio


// Auth configures how secret-manager authenticates with the Oracle Vault.
//
// If empty, use the instance principal, otherwise the user credentials specified in Auth.
type SecretStoreSpecProviderOracleAuth struct {
	// SecretRef to pass through sensitive information.
	SecretRef *SecretStoreSpecProviderOracleAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
	// Tenancy is the tenancy OCID where user is located.
	Tenancy *string `field:"required" json:"tenancy" yaml:"tenancy"`
	// User is an access OCID specific to the account.
	User *string `field:"required" json:"user" yaml:"user"`
}

