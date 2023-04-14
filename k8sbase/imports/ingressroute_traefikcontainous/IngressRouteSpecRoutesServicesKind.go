// ingressroute_traefikcontainous
package ingressroute_traefikcontainous


// Kind defines the kind of the Service.
type IngressRouteSpecRoutesServicesKind string

const (
	// Service.
	IngressRouteSpecRoutesServicesKind_SERVICE IngressRouteSpecRoutesServicesKind = "SERVICE"
	// TraefikService.
	IngressRouteSpecRoutesServicesKind_TRAEFIK_SERVICE IngressRouteSpecRoutesServicesKind = "TRAEFIK_SERVICE"
)

