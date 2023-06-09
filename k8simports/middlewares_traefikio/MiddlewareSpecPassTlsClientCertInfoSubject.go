package middlewares_traefikio


// Subject defines the client certificate subject details to add to the X-Forwarded-Tls-Client-Cert-Info header.
type MiddlewareSpecPassTlsClientCertInfoSubject struct {
	// CommonName defines whether to add the organizationalUnit information into the subject.
	CommonName *bool `field:"optional" json:"commonName" yaml:"commonName"`
	// Country defines whether to add the country information into the subject.
	Country *bool `field:"optional" json:"country" yaml:"country"`
	// DomainComponent defines whether to add the domainComponent information into the subject.
	DomainComponent *bool `field:"optional" json:"domainComponent" yaml:"domainComponent"`
	// Locality defines whether to add the locality information into the subject.
	Locality *bool `field:"optional" json:"locality" yaml:"locality"`
	// Organization defines whether to add the organization information into the subject.
	Organization *bool `field:"optional" json:"organization" yaml:"organization"`
	// OrganizationalUnit defines whether to add the organizationalUnit information into the subject.
	OrganizationalUnit *bool `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// Province defines whether to add the province information into the subject.
	Province *bool `field:"optional" json:"province" yaml:"province"`
	// SerialNumber defines whether to add the serialNumber information into the subject.
	SerialNumber *bool `field:"optional" json:"serialNumber" yaml:"serialNumber"`
}

