package ingressrouteudp_traefikio


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
	// NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP.
	//
	// The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.
	NativeLb *bool `field:"optional" json:"nativeLb" yaml:"nativeLb"`
	// Weight defines the weight used when balancing requests between multiple Kubernetes Service.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

