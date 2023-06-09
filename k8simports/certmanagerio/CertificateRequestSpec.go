package certmanagerio


// Desired state of the CertificateRequest resource.
type CertificateRequestSpec struct {
	// IssuerRef is a reference to the issuer for this CertificateRequest.
	//
	// If the `kind` field is not set, or set to `Issuer`, an Issuer resource with the given name in the same namespace as the CertificateRequest will be used.  If the `kind` field is set to `ClusterIssuer`, a ClusterIssuer with the provided name will be used. The `name` field in this stanza is required at all times. The group field refers to the API group of the issuer which defaults to `cert-manager.io` if empty.
	IssuerRef *CertificateRequestSpecIssuerRef `field:"required" json:"issuerRef" yaml:"issuerRef"`
	// The PEM-encoded x509 certificate signing request to be submitted to the CA for signing.
	Request *string `field:"required" json:"request" yaml:"request"`
	// The requested 'duration' (i.e. lifetime) of the Certificate. This option may be ignored/overridden by some issuer types.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Extra contains extra attributes of the user that created the CertificateRequest.
	//
	// Populated by the cert-manager webhook on creation and immutable.
	Extra *map[string]*[]*string `field:"optional" json:"extra" yaml:"extra"`
	// Groups contains group membership of the user that created the CertificateRequest.
	//
	// Populated by the cert-manager webhook on creation and immutable.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// IsCA will request to mark the certificate as valid for certificate signing when submitting to the issuer.
	//
	// This will automatically add the `cert sign` usage to the list of `usages`.
	IsCa *bool `field:"optional" json:"isCa" yaml:"isCa"`
	// UID contains the uid of the user that created the CertificateRequest.
	//
	// Populated by the cert-manager webhook on creation and immutable.
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
	// Usages is the set of x509 usages that are requested for the certificate.
	//
	// If usages are set they SHOULD be encoded inside the CSR spec Defaults to `digital signature` and `key encipherment` if not specified.
	Usages *[]CertificateRequestSpecUsages `field:"optional" json:"usages" yaml:"usages"`
	// Username contains the name of the user that created the CertificateRequest.
	//
	// Populated by the cert-manager webhook on creation and immutable.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

