// cert-managerio
package certmanagerio


// Desired state of the ClusterIssuer resource.
type ClusterIssuerSpec struct {
	// ACME configures this issuer to communicate with a RFC8555 (ACME) server to obtain signed x509 certificates.
	Acme *ClusterIssuerSpecAcme `field:"optional" json:"acme" yaml:"acme"`
	// CA configures this issuer to sign certificates using a signing CA keypair stored in a Secret resource.
	//
	// This is used to build internal PKIs that are managed by cert-manager.
	Ca *ClusterIssuerSpecCa `field:"optional" json:"ca" yaml:"ca"`
	// SelfSigned configures this issuer to 'self sign' certificates using the private key used to create the CertificateRequest object.
	SelfSigned *ClusterIssuerSpecSelfSigned `field:"optional" json:"selfSigned" yaml:"selfSigned"`
	// Vault configures this issuer to sign certificates using a HashiCorp Vault PKI backend.
	Vault *ClusterIssuerSpecVault `field:"optional" json:"vault" yaml:"vault"`
	// Venafi configures this issuer to sign certificates using a Venafi TPP or Venafi Cloud policy zone.
	Venafi *ClusterIssuerSpecVenafi `field:"optional" json:"venafi" yaml:"venafi"`
}

