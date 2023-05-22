package ingressrouteudp_traefikio


// IngressRouteUDPSpec defines the desired state of a IngressRouteUDP.
type IngressRouteUdpSpec struct {
	// Routes defines the list of routes.
	Routes *[]*IngressRouteUdpSpecRoutes `field:"required" json:"routes" yaml:"routes"`
	// EntryPoints defines the list of entry point names to bind to.
	//
	// Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v2.10/routing/entrypoints/ Default: all.
	EntryPoints *[]*string `field:"optional" json:"entryPoints" yaml:"entryPoints"`
}

