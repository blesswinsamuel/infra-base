package ingressroutetcp_traefikio


// IngressRouteTCPSpec defines the desired state of IngressRouteTCP.
type IngressRouteTcpSpec struct {
	// Routes defines the list of routes.
	Routes *[]*IngressRouteTcpSpecRoutes `field:"required" json:"routes" yaml:"routes"`
	// EntryPoints defines the list of entry point names to bind to.
	//
	// Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v2.10/routing/entrypoints/ Default: all.
	EntryPoints *[]*string `field:"optional" json:"entryPoints" yaml:"entryPoints"`
	// TLS defines the TLS configuration on a layer 4 / TCP Route.
	//
	// More info: https://doc.traefik.io/traefik/v2.10/routing/routers/#tls_1
	Tls *IngressRouteTcpSpecTls `field:"optional" json:"tls" yaml:"tls"`
}

