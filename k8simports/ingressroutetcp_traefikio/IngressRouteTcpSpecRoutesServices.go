package ingressroutetcp_traefikio


// ServiceTCP defines an upstream TCP service to proxy traffic to.
type IngressRouteTcpSpecRoutesServices struct {
	// Name defines the name of the referenced Kubernetes Service.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Port defines the port of a Kubernetes Service.
	//
	// This can be a reference to a named port.
	Port IngressRouteTcpSpecRoutesServicesPort `field:"required" json:"port" yaml:"port"`
	// Namespace defines the namespace of the referenced Kubernetes Service.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP.
	//
	// The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.
	NativeLb *bool `field:"optional" json:"nativeLb" yaml:"nativeLb"`
	// ProxyProtocol defines the PROXY protocol configuration.
	//
	// More info: https://doc.traefik.io/traefik/v2.10/routing/services/#proxy-protocol
	ProxyProtocol *IngressRouteTcpSpecRoutesServicesProxyProtocol `field:"optional" json:"proxyProtocol" yaml:"proxyProtocol"`
	// TerminationDelay defines the deadline that the proxy sets, after one of its connected peers indicates it has closed the writing capability of its connection, to close the reading capability as well, hence fully terminating the connection.
	//
	// It is a duration in milliseconds, defaulting to 100. A negative value means an infinite deadline (i.e. the reading capability is never closed).
	TerminationDelay *float64 `field:"optional" json:"terminationDelay" yaml:"terminationDelay"`
	// Weight defines the weight used when balancing requests between multiple Kubernetes Service.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

