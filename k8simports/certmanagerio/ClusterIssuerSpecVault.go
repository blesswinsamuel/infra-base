package certmanagerio


// Vault configures this issuer to sign certificates using a HashiCorp Vault PKI backend.
type ClusterIssuerSpecVault struct {
	// Auth configures how cert-manager authenticates with the Vault server.
	Auth *ClusterIssuerSpecVaultAuth `field:"required" json:"auth" yaml:"auth"`
	// Path is the mount path of the Vault PKI backend's `sign` endpoint, e.g: "my_pki_mount/sign/my-role-name".
	Path *string `field:"required" json:"path" yaml:"path"`
	// Server is the connection address for the Vault server, e.g: "https://vault.example.com:8200".
	Server *string `field:"required" json:"server" yaml:"server"`
	// Base64-encoded bundle of PEM CAs which will be used to validate the certificate chain presented by Vault.
	//
	// Only used if using HTTPS to connect to Vault and ignored for HTTP connections. Mutually exclusive with CABundleSecretRef. If neither CABundle nor CABundleSecretRef are defined, the certificate bundle in the cert-manager controller container is used to validate the TLS connection.
	CaBundle *string `field:"optional" json:"caBundle" yaml:"caBundle"`
	// Reference to a Secret containing a bundle of PEM-encoded CAs to use when verifying the certificate chain presented by Vault when using HTTPS.
	//
	// Mutually exclusive with CABundle. If neither CABundle nor CABundleSecretRef are defined, the certificate bundle in the cert-manager controller container is used to validate the TLS connection. If no key for the Secret is specified, cert-manager will default to 'ca.crt'.
	CaBundleSecretRef *ClusterIssuerSpecVaultCaBundleSecretRef `field:"optional" json:"caBundleSecretRef" yaml:"caBundleSecretRef"`
	// Name of the vault namespace.
	//
	// Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: "ns1" More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

