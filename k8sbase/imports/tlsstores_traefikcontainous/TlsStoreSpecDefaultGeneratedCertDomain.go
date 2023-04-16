// tlsstores_traefikcontainous
package tlsstores_traefikcontainous


// Domain is the domain definition for the DefaultCertificate.
type TlsStoreSpecDefaultGeneratedCertDomain struct {
	// Main defines the main domain name.
	Main *string `field:"optional" json:"main" yaml:"main"`
	// SANs defines the subject alternative domain names.
	Sans *[]*string `field:"optional" json:"sans" yaml:"sans"`
}

