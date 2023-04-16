// traefikservices_traefikcontainous
package traefikservices_traefikcontainous


// Weighted defines the Weighted Round Robin configuration.
type TraefikServiceSpecWeighted struct {
	// Services defines the list of Kubernetes Service and/or TraefikService to load-balance, with weight.
	Services *[]*TraefikServiceSpecWeightedServices `field:"optional" json:"services" yaml:"services"`
	// Sticky defines whether sticky sessions are enabled.
	//
	// More info: https://doc.traefik.io/traefik/v2.9/routing/providers/kubernetes-crd/#stickiness-and-load-balancing
	Sticky *TraefikServiceSpecWeightedSticky `field:"optional" json:"sticky" yaml:"sticky"`
}

