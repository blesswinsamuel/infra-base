package middlewares_traefikio


// ReplacePath holds the replace path middleware configuration.
//
// This middleware replaces the path of the request URL and store the original path in an X-Replaced-Path header. More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/replacepath/
type MiddlewareSpecReplacePath struct {
	// Path defines the path to use as replacement in the request URL.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

