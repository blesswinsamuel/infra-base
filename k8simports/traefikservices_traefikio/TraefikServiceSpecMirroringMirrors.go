package traefikservices_traefikio


// MirrorService holds the mirror configuration.
type TraefikServiceSpecMirroringMirrors struct {
	// Name defines the name of the referenced Kubernetes Service or TraefikService.
	//
	// The differentiation between the two is specified in the Kind field.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Kind defines the kind of the Service.
	Kind TraefikServiceSpecMirroringMirrorsKind `field:"optional" json:"kind" yaml:"kind"`
	// Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP.
	//
	// The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.
	NativeLb *bool `field:"optional" json:"nativeLb" yaml:"nativeLb"`
	// PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service.
	//
	// By default, passHostHeader is true.
	PassHostHeader *bool `field:"optional" json:"passHostHeader" yaml:"passHostHeader"`
	// Percent defines the part of the traffic to mirror.
	//
	// Supported values: 0 to 100.
	Percent *float64 `field:"optional" json:"percent" yaml:"percent"`
	// Port defines the port of a Kubernetes Service.
	//
	// This can be a reference to a named port.
	Port TraefikServiceSpecMirroringMirrorsPort `field:"optional" json:"port" yaml:"port"`
	// ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.
	ResponseForwarding *TraefikServiceSpecMirroringMirrorsResponseForwarding `field:"optional" json:"responseForwarding" yaml:"responseForwarding"`
	// Scheme defines the scheme to use for the request to the upstream Kubernetes Service.
	//
	// It defaults to https when Kubernetes Service port is 443, http otherwise.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// ServersTransport defines the name of ServersTransport resource to use.
	//
	// It allows to configure the transport between Traefik and your servers. Can only be used on a Kubernetes Service.
	ServersTransport *string `field:"optional" json:"serversTransport" yaml:"serversTransport"`
	// Sticky defines the sticky sessions configuration.
	//
	// More info: https://doc.traefik.io/traefik/v2.10/routing/services/#sticky-sessions
	Sticky *TraefikServiceSpecMirroringMirrorsSticky `field:"optional" json:"sticky" yaml:"sticky"`
	// Strategy defines the load balancing strategy between the servers.
	//
	// RoundRobin is the only supported value at the moment.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
	// Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

