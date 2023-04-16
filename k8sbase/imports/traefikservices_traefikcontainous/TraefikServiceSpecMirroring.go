// traefikservices_traefikcontainous
package traefikservices_traefikcontainous


// Mirroring defines the Mirroring service configuration.
type TraefikServiceSpecMirroring struct {
	// Name defines the name of the referenced Kubernetes Service or TraefikService.
	//
	// The differentiation between the two is specified in the Kind field.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Kind defines the kind of the Service.
	Kind TraefikServiceSpecMirroringKind `field:"optional" json:"kind" yaml:"kind"`
	// MaxBodySize defines the maximum size allowed for the body of the request.
	//
	// If the body is larger, the request is not mirrored. Default value is -1, which means unlimited size.
	MaxBodySize *float64 `field:"optional" json:"maxBodySize" yaml:"maxBodySize"`
	// Mirrors defines the list of mirrors where Traefik will duplicate the traffic.
	Mirrors *[]*TraefikServiceSpecMirroringMirrors `field:"optional" json:"mirrors" yaml:"mirrors"`
	// Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service.
	//
	// By default, passHostHeader is true.
	PassHostHeader *bool `field:"optional" json:"passHostHeader" yaml:"passHostHeader"`
	// Port defines the port of a Kubernetes Service.
	//
	// This can be a reference to a named port.
	Port TraefikServiceSpecMirroringPort `field:"optional" json:"port" yaml:"port"`
	// ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.
	ResponseForwarding *TraefikServiceSpecMirroringResponseForwarding `field:"optional" json:"responseForwarding" yaml:"responseForwarding"`
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
	// More info: https://doc.traefik.io/traefik/v2.9/routing/services/#sticky-sessions
	Sticky *TraefikServiceSpecMirroringSticky `field:"optional" json:"sticky" yaml:"sticky"`
	// Strategy defines the load balancing strategy between the servers.
	//
	// RoundRobin is the only supported value at the moment.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
	// Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

