// ingressroutetcp_traefikcontainous
package ingressroutetcp_traefikcontainous


// Store defines the reference to the TLSStore, that will be used to store certificates.
//
// Please note that only `default` TLSStore can be used.
type IngressRouteTcpSpecTlsStore struct {
	// Name defines the name of the referenced Traefik resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace defines the namespace of the referenced Traefik resource.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

