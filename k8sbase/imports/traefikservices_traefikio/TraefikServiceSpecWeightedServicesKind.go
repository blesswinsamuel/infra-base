package traefikservices_traefikio


// Kind defines the kind of the Service.
type TraefikServiceSpecWeightedServicesKind string

const (
	// Service.
	TraefikServiceSpecWeightedServicesKind_SERVICE TraefikServiceSpecWeightedServicesKind = "SERVICE"
	// TraefikService.
	TraefikServiceSpecWeightedServicesKind_TRAEFIK_SERVICE TraefikServiceSpecWeightedServicesKind = "TRAEFIK_SERVICE"
)

