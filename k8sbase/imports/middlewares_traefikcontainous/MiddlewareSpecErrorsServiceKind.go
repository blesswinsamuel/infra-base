// middlewares_traefikcontainous
package middlewares_traefikcontainous


// Kind defines the kind of the Service.
type MiddlewareSpecErrorsServiceKind string

const (
	// Service.
	MiddlewareSpecErrorsServiceKind_SERVICE MiddlewareSpecErrorsServiceKind = "SERVICE"
	// TraefikService.
	MiddlewareSpecErrorsServiceKind_TRAEFIK_SERVICE MiddlewareSpecErrorsServiceKind = "TRAEFIK_SERVICE"
)

