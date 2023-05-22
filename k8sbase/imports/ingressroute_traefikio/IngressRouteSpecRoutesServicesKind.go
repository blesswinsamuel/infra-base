package ingressroute_traefikio


// Kind defines the kind of the Service.
type IngressRouteSpecRoutesServicesKind string

const (
	// Service.
	IngressRouteSpecRoutesServicesKind_SERVICE IngressRouteSpecRoutesServicesKind = "SERVICE"
	// TraefikService.
	IngressRouteSpecRoutesServicesKind_TRAEFIK_SERVICE IngressRouteSpecRoutesServicesKind = "TRAEFIK_SERVICE"
)

