// traefikservices_traefikcontainous
package traefikservices_traefikcontainous


// Sticky defines the sticky sessions configuration.
//
// More info: https://doc.traefik.io/traefik/v2.9/routing/services/#sticky-sessions
type TraefikServiceSpecWeightedServicesSticky struct {
	// Cookie defines the sticky cookie configuration.
	Cookie *TraefikServiceSpecWeightedServicesStickyCookie `field:"optional" json:"cookie" yaml:"cookie"`
}

