// traefikservices_traefikcontainous
package traefikservices_traefikcontainous


// Sticky defines whether sticky sessions are enabled.
//
// More info: https://doc.traefik.io/traefik/v2.9/routing/providers/kubernetes-crd/#stickiness-and-load-balancing
type TraefikServiceSpecWeightedSticky struct {
	// Cookie defines the sticky cookie configuration.
	Cookie *TraefikServiceSpecWeightedStickyCookie `field:"optional" json:"cookie" yaml:"cookie"`
}

