package tlsstores_traefikio


// DefaultCertificate defines the default certificate configuration.
type TlsStoreSpecDefaultCertificate struct {
	// SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

