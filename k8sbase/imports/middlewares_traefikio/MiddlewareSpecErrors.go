package middlewares_traefikio


// ErrorPage holds the custom error middleware configuration.
//
// This middleware returns a custom page in lieu of the default, according to configured ranges of HTTP Status codes. More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/errorpages/
type MiddlewareSpecErrors struct {
	// Query defines the URL for the error page (hosted by service).
	//
	// The {status} variable can be used in order to insert the status code in the URL.
	Query *string `field:"optional" json:"query" yaml:"query"`
	// Service defines the reference to a Kubernetes Service that will serve the error page.
	//
	// More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/errorpages/#service
	Service *MiddlewareSpecErrorsService `field:"optional" json:"service" yaml:"service"`
	// Status defines which status or range of statuses should result in an error page.
	//
	// It can be either a status code as a number (500), as multiple comma-separated numbers (500,502), as ranges by separating two codes with a dash (500-599), or a combination of the two (404,418,500-599).
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
}

