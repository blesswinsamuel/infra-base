package ingressroute_traefikio


// Sticky defines the sticky sessions configuration.
//
// More info: https://doc.traefik.io/traefik/v2.10/routing/services/#sticky-sessions
type IngressRouteSpecRoutesServicesSticky struct {
	// Cookie defines the sticky cookie configuration.
	Cookie *IngressRouteSpecRoutesServicesStickyCookie `field:"optional" json:"cookie" yaml:"cookie"`
}

