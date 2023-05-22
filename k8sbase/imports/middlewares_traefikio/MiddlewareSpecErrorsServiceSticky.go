package middlewares_traefikio


// Sticky defines the sticky sessions configuration.
//
// More info: https://doc.traefik.io/traefik/v2.10/routing/services/#sticky-sessions
type MiddlewareSpecErrorsServiceSticky struct {
	// Cookie defines the sticky cookie configuration.
	Cookie *MiddlewareSpecErrorsServiceStickyCookie `field:"optional" json:"cookie" yaml:"cookie"`
}

