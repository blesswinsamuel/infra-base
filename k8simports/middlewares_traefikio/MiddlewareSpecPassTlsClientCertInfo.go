package middlewares_traefikio


// Info selects the specific client certificate details you want to add to the X-Forwarded-Tls-Client-Cert-Info header.
type MiddlewareSpecPassTlsClientCertInfo struct {
	// Issuer defines the client certificate issuer details to add to the X-Forwarded-Tls-Client-Cert-Info header.
	Issuer *MiddlewareSpecPassTlsClientCertInfoIssuer `field:"optional" json:"issuer" yaml:"issuer"`
	// NotAfter defines whether to add the Not After information from the Validity part.
	NotAfter *bool `field:"optional" json:"notAfter" yaml:"notAfter"`
	// NotBefore defines whether to add the Not Before information from the Validity part.
	NotBefore *bool `field:"optional" json:"notBefore" yaml:"notBefore"`
	// Sans defines whether to add the Subject Alternative Name information from the Subject Alternative Name part.
	Sans *bool `field:"optional" json:"sans" yaml:"sans"`
	// SerialNumber defines whether to add the client serialNumber information.
	SerialNumber *bool `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// Subject defines the client certificate subject details to add to the X-Forwarded-Tls-Client-Cert-Info header.
	Subject *MiddlewareSpecPassTlsClientCertInfoSubject `field:"optional" json:"subject" yaml:"subject"`
}

