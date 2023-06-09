package middlewares_traefikio


// RedirectScheme holds the redirect scheme middleware configuration.
//
// This middleware redirects requests from a scheme/port to another. More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/redirectscheme/
type MiddlewareSpecRedirectScheme struct {
	// Permanent defines whether the redirection is permanent (301).
	Permanent *bool `field:"optional" json:"permanent" yaml:"permanent"`
	// Port defines the port of the new URL.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// Scheme defines the scheme of the new URL.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

