package ingressrouteudp_traefikio


// RouteUDP holds the UDP route configuration.
type IngressRouteUdpSpecRoutes struct {
	// Services defines the list of UDP services.
	Services *[]*IngressRouteUdpSpecRoutesServices `field:"optional" json:"services" yaml:"services"`
}

