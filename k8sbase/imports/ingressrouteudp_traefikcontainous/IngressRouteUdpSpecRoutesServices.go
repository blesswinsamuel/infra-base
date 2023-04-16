// ingressrouteudp_traefikcontainous
package ingressrouteudp_traefikcontainous


// ServiceUDP defines an upstream UDP service to proxy traffic to.
type IngressRouteUdpSpecRoutesServices struct {
	// Name defines the name of the referenced Kubernetes Service.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Port defines the port of a Kubernetes Service.
	//
	// This can be a reference to a named port.
	Port IngressRouteUdpSpecRoutesServicesPort `field:"required" json:"port" yaml:"port"`
	// Namespace defines the namespace of the referenced Kubernetes Service.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Weight defines the weight used when balancing requests between multiple Kubernetes Service.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

