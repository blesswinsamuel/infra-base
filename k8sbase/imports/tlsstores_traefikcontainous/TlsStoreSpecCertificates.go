// tlsstores_traefikcontainous
package tlsstores_traefikcontainous


// Certificate holds a secret name for the TLSStore resource.
type TlsStoreSpecCertificates struct {
	// SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

