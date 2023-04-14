// external-secretsio
package externalsecretsio


// Auth configures how secret-manager authenticates with the Vault server.
type ClusterSecretStoreV1Beta1SpecProviderVaultAuth struct {
	// AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.
	AppRole *ClusterSecretStoreV1Beta1SpecProviderVaultAuthAppRole `field:"optional" json:"appRole" yaml:"appRole"`
	// Cert authenticates with TLS Certificates by passing client certificate, private key and ca certificate Cert authentication method.
	Cert *ClusterSecretStoreV1Beta1SpecProviderVaultAuthCert `field:"optional" json:"cert" yaml:"cert"`
	// Jwt authenticates with Vault by passing role and JWT token using the JWT/OIDC authentication method.
	Jwt *ClusterSecretStoreV1Beta1SpecProviderVaultAuthJwt `field:"optional" json:"jwt" yaml:"jwt"`
	// Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.
	Kubernetes *ClusterSecretStoreV1Beta1SpecProviderVaultAuthKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method.
	Ldap *ClusterSecretStoreV1Beta1SpecProviderVaultAuthLdap `field:"optional" json:"ldap" yaml:"ldap"`
	// TokenSecretRef authenticates with Vault by presenting a token.
	TokenSecretRef *ClusterSecretStoreV1Beta1SpecProviderVaultAuthTokenSecretRef `field:"optional" json:"tokenSecretRef" yaml:"tokenSecretRef"`
}

