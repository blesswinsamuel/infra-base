// tlsstores_traefikcontainous
package tlsstores_traefikcontainous


// TLSStoreSpec defines the desired state of a TLSStore.
type TlsStoreSpec struct {
	// Certificates is a list of secret names, each secret holding a key/certificate pair to add to the store.
	Certificates *[]*TlsStoreSpecCertificates `field:"optional" json:"certificates" yaml:"certificates"`
	// DefaultCertificate defines the default certificate configuration.
	DefaultCertificate *TlsStoreSpecDefaultCertificate `field:"optional" json:"defaultCertificate" yaml:"defaultCertificate"`
	// DefaultGeneratedCert defines the default generated certificate configuration.
	DefaultGeneratedCert *TlsStoreSpecDefaultGeneratedCert `field:"optional" json:"defaultGeneratedCert" yaml:"defaultGeneratedCert"`
}

