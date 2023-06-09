package middlewares_traefikio


// TLS defines the configuration used to secure the connection to the authentication server.
type MiddlewareSpecForwardAuthTls struct {
	CaOptional *bool `field:"optional" json:"caOptional" yaml:"caOptional"`
	// CASecret is the name of the referenced Kubernetes Secret containing the CA to validate the server certificate.
	//
	// The CA certificate is extracted from key `tls.ca` or `ca.crt`.
	CaSecret *string `field:"optional" json:"caSecret" yaml:"caSecret"`
	// CertSecret is the name of the referenced Kubernetes Secret containing the client certificate.
	//
	// The client certificate is extracted from the keys `tls.crt` and `tls.key`.
	CertSecret *string `field:"optional" json:"certSecret" yaml:"certSecret"`
	// InsecureSkipVerify defines whether the server certificates should be validated.
	InsecureSkipVerify *bool `field:"optional" json:"insecureSkipVerify" yaml:"insecureSkipVerify"`
}

