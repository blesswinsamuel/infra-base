// middlewares_traefikcontainous
package middlewares_traefikcontainous


// Retry holds the retry middleware configuration.
//
// This middleware reissues requests a given number of times to a backend server if that server does not reply. As soon as the server answers, the middleware stops retrying, regardless of the response status. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/retry/
type MiddlewareSpecRetry struct {
	// Attempts defines how many times the request should be retried.
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
	// InitialInterval defines the first wait time in the exponential backoff series.
	//
	// The maximum interval is calculated as twice the initialInterval. If unspecified, requests will be retried immediately. The value of initialInterval should be provided in seconds or as a valid duration format, see https://pkg.go.dev/time#ParseDuration.
	InitialInterval MiddlewareSpecRetryInitialInterval `field:"optional" json:"initialInterval" yaml:"initialInterval"`
}

