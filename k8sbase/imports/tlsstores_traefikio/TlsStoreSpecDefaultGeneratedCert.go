package tlsstores_traefikio


// DefaultGeneratedCert defines the default generated certificate configuration.
type TlsStoreSpecDefaultGeneratedCert struct {
	// Domain is the domain definition for the DefaultCertificate.
	Domain *TlsStoreSpecDefaultGeneratedCertDomain `field:"optional" json:"domain" yaml:"domain"`
	// Resolver is the name of the resolver that will be used to issue the DefaultCertificate.
	Resolver *string `field:"optional" json:"resolver" yaml:"resolver"`
}

