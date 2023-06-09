package ingressroute_traefikio


// IngressRouteSpec defines the desired state of IngressRoute.
type IngressRouteSpec struct {
	// Routes defines the list of routes.
	Routes *[]*IngressRouteSpecRoutes `field:"required" json:"routes" yaml:"routes"`
	// EntryPoints defines the list of entry point names to bind to.
	//
	// Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v2.10/routing/entrypoints/ Default: all.
	EntryPoints *[]*string `field:"optional" json:"entryPoints" yaml:"entryPoints"`
	// TLS defines the TLS configuration.
	//
	// More info: https://doc.traefik.io/traefik/v2.10/routing/routers/#tls
	Tls *IngressRouteSpecTls `field:"optional" json:"tls" yaml:"tls"`
}

