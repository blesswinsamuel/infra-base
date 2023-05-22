package externalsecretsio


// Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method.
type SecretStoreV1Beta1SpecProviderVaultAuthLdap struct {
	// Path where the LDAP authentication backend is mounted in Vault, e.g: "ldap".
	Path *string `field:"required" json:"path" yaml:"path"`
	// Username is a LDAP user name used to authenticate using the LDAP Vault authentication method.
	Username *string `field:"required" json:"username" yaml:"username"`
	// SecretRef to a key in a Secret resource containing password for the LDAP user used to authenticate with Vault using the LDAP authentication method.
	SecretRef *SecretStoreV1Beta1SpecProviderVaultAuthLdapSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

