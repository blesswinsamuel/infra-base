// cert-managerio
package certmanagerio


// Auth configures how cert-manager authenticates with the Vault server.
type ClusterIssuerSpecVaultAuth struct {
	// AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.
	AppRole *ClusterIssuerSpecVaultAuthAppRole `field:"optional" json:"appRole" yaml:"appRole"`
	// Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.
	Kubernetes *ClusterIssuerSpecVaultAuthKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// TokenSecretRef authenticates with Vault by presenting a token.
	TokenSecretRef *ClusterIssuerSpecVaultAuthTokenSecretRef `field:"optional" json:"tokenSecretRef" yaml:"tokenSecretRef"`
}

