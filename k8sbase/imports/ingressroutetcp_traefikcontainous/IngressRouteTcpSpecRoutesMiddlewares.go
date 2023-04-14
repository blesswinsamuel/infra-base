// ingressroutetcp_traefikcontainous
package ingressroutetcp_traefikcontainous


// ObjectReference is a generic reference to a Traefik resource.
type IngressRouteTcpSpecRoutesMiddlewares struct {
	// Name defines the name of the referenced Traefik resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace defines the namespace of the referenced Traefik resource.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

