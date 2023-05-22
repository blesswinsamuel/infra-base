package certmanagerio


// ACME configures this issuer to communicate with a RFC8555 (ACME) server to obtain signed x509 certificates.
type ClusterIssuerSpecAcme struct {
	// PrivateKey is the name of a Kubernetes Secret resource that will be used to store the automatically generated ACME account private key.
	//
	// Optionally, a `key` may be specified to select a specific entry within the named Secret resource. If `key` is not specified, a default of `tls.key` will be used.
	PrivateKeySecretRef *ClusterIssuerSpecAcmePrivateKeySecretRef `field:"required" json:"privateKeySecretRef" yaml:"privateKeySecretRef"`
	// Server is the URL used to access the ACME server's 'directory' endpoint.
	//
	// For example, for Let's Encrypt's staging endpoint, you would use: "https://acme-staging-v02.api.letsencrypt.org/directory". Only ACME v2 endpoints (i.e. RFC 8555) are supported.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Base64-encoded bundle of PEM CAs which can be used to validate the certificate chain presented by the ACME server.
	//
	// Mutually exclusive with SkipTLSVerify; prefer using CABundle to prevent various kinds of security vulnerabilities. If CABundle and SkipTLSVerify are unset, the system certificate bundle inside the container is used to validate the TLS connection.
	CaBundle *string `field:"optional" json:"caBundle" yaml:"caBundle"`
	// Enables or disables generating a new ACME account key.
	//
	// If true, the Issuer resource will *not* request a new account but will expect the account key to be supplied via an existing secret. If false, the cert-manager system will generate a new ACME account key for the Issuer. Defaults to false.
	DisableAccountKeyGeneration *bool `field:"optional" json:"disableAccountKeyGeneration" yaml:"disableAccountKeyGeneration"`
	// Email is the email address to be associated with the ACME account.
	//
	// This field is optional, but it is strongly recommended to be set. It will be used to contact you in case of issues with your account or certificates, including expiry notification emails. This field may be updated after the account is initially registered.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Enables requesting a Not After date on certificates that matches the duration of the certificate.
	//
	// This is not supported by all ACME servers like Let's Encrypt. If set to true when the ACME server does not support it it will create an error on the Order. Defaults to false.
	EnableDurationFeature *bool `field:"optional" json:"enableDurationFeature" yaml:"enableDurationFeature"`
	// ExternalAccountBinding is a reference to a CA external account of the ACME server.
	//
	// If set, upon registration cert-manager will attempt to associate the given external account credentials with the registered ACME account.
	ExternalAccountBinding *ClusterIssuerSpecAcmeExternalAccountBinding `field:"optional" json:"externalAccountBinding" yaml:"externalAccountBinding"`
	// PreferredChain is the chain to use if the ACME server outputs multiple.
	//
	// PreferredChain is no guarantee that this one gets delivered by the ACME endpoint. For example, for Let's Encrypt's DST crosssign you would use: "DST Root CA X3" or "ISRG Root X1" for the newer Let's Encrypt root CA. This value picks the first certificate bundle in the ACME alternative chains that has a certificate with this value as its issuer's CN
	PreferredChain *string `field:"optional" json:"preferredChain" yaml:"preferredChain"`
	// INSECURE: Enables or disables validation of the ACME server TLS certificate.
	//
	// If true, requests to the ACME server will not have the TLS certificate chain validated. Mutually exclusive with CABundle; prefer using CABundle to prevent various kinds of security vulnerabilities. Only enable this option in development environments. If CABundle and SkipTLSVerify are unset, the system certificate bundle inside the container is used to validate the TLS connection. Defaults to false.
	SkipTlsVerify *bool `field:"optional" json:"skipTlsVerify" yaml:"skipTlsVerify"`
	// Solvers is a list of challenge solvers that will be used to solve ACME challenges for the matching domains.
	//
	// Solver configurations must be provided in order to obtain certificates from an ACME server. For more information, see: https://cert-manager.io/docs/configuration/acme/
	Solvers *[]*ClusterIssuerSpecAcmeSolvers `field:"optional" json:"solvers" yaml:"solvers"`
}

