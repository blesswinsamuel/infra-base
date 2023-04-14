// cert-managerio
package certmanagerio


// SelfSigned configures this issuer to 'self sign' certificates using the private key used to create the CertificateRequest object.
type ClusterIssuerSpecSelfSigned struct {
	// The CRL distribution points is an X.509 v3 certificate extension which identifies the location of the CRL from which the revocation of this certificate can be checked. If not set certificate will be issued without CDP. Values are strings.
	CrlDistributionPoints *[]*string `field:"optional" json:"crlDistributionPoints" yaml:"crlDistributionPoints"`
}

