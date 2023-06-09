package certmanagerio


// CA configures this issuer to sign certificates using a signing CA keypair stored in a Secret resource.
//
// This is used to build internal PKIs that are managed by cert-manager.
type IssuerSpecCa struct {
	// SecretName is the name of the secret used to sign Certificates issued by this Issuer.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// The CRL distribution points is an X.509 v3 certificate extension which identifies the location of the CRL from which the revocation of this certificate can be checked. If not set, certificates will be issued without distribution points set.
	CrlDistributionPoints *[]*string `field:"optional" json:"crlDistributionPoints" yaml:"crlDistributionPoints"`
	// The OCSP server list is an X.509 v3 extension that defines a list of URLs of OCSP responders. The OCSP responders can be queried for the revocation status of an issued certificate. If not set, the certificate will be issued with no OCSP servers set. For example, an OCSP server URL could be "http://ocsp.int-x3.letsencrypt.org".
	OcspServers *[]*string `field:"optional" json:"ocspServers" yaml:"ocspServers"`
}

