// ingressroute_traefikcontainous
package ingressroute_traefikcontainous


// Sticky defines the sticky sessions configuration.
//
// More info: https://doc.traefik.io/traefik/v2.9/routing/services/#sticky-sessions
type IngressRouteSpecRoutesServicesSticky struct {
	// Cookie defines the sticky cookie configuration.
	Cookie *IngressRouteSpecRoutesServicesStickyCookie `field:"optional" json:"cookie" yaml:"cookie"`
}

