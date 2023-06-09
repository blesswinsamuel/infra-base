package middlewares_traefikio


// Issuer defines the client certificate issuer details to add to the X-Forwarded-Tls-Client-Cert-Info header.
type MiddlewareSpecPassTlsClientCertInfoIssuer struct {
	// CommonName defines whether to add the organizationalUnit information into the issuer.
	CommonName *bool `field:"optional" json:"commonName" yaml:"commonName"`
	// Country defines whether to add the country information into the issuer.
	Country *bool `field:"optional" json:"country" yaml:"country"`
	// DomainComponent defines whether to add the domainComponent information into the issuer.
	DomainComponent *bool `field:"optional" json:"domainComponent" yaml:"domainComponent"`
	// Locality defines whether to add the locality information into the issuer.
	Locality *bool `field:"optional" json:"locality" yaml:"locality"`
	// Organization defines whether to add the organization information into the issuer.
	Organization *bool `field:"optional" json:"organization" yaml:"organization"`
	// Province defines whether to add the province information into the issuer.
	Province *bool `field:"optional" json:"province" yaml:"province"`
	// SerialNumber defines whether to add the serialNumber information into the issuer.
	SerialNumber *bool `field:"optional" json:"serialNumber" yaml:"serialNumber"`
}

