// traefikservices_traefikcontainous
package traefikservices_traefikcontainous


// TraefikServiceSpec defines the desired state of a TraefikService.
type TraefikServiceSpec struct {
	// Mirroring defines the Mirroring service configuration.
	Mirroring *TraefikServiceSpecMirroring `field:"optional" json:"mirroring" yaml:"mirroring"`
	// Weighted defines the Weighted Round Robin configuration.
	Weighted *TraefikServiceSpecWeighted `field:"optional" json:"weighted" yaml:"weighted"`
}

