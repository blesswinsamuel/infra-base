// middlewares_traefikcontainous
package middlewares_traefikcontainous


// Buffering holds the buffering middleware configuration.
//
// This middleware retries or limits the size of requests that can be forwarded to backends. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/buffering/#maxrequestbodybytes
type MiddlewareSpecBuffering struct {
	// MaxRequestBodyBytes defines the maximum allowed body size for the request (in bytes).
	//
	// If the request exceeds the allowed size, it is not forwarded to the service, and the client gets a 413 (Request Entity Too Large) response. Default: 0 (no maximum).
	MaxRequestBodyBytes *float64 `field:"optional" json:"maxRequestBodyBytes" yaml:"maxRequestBodyBytes"`
	// MaxResponseBodyBytes defines the maximum allowed response size from the service (in bytes).
	//
	// If the response exceeds the allowed size, it is not forwarded to the client. The client gets a 500 (Internal Server Error) response instead. Default: 0 (no maximum).
	MaxResponseBodyBytes *float64 `field:"optional" json:"maxResponseBodyBytes" yaml:"maxResponseBodyBytes"`
	// MemRequestBodyBytes defines the threshold (in bytes) from which the request will be buffered on disk instead of in memory.
	//
	// Default: 1048576 (1Mi).
	MemRequestBodyBytes *float64 `field:"optional" json:"memRequestBodyBytes" yaml:"memRequestBodyBytes"`
	// MemResponseBodyBytes defines the threshold (in bytes) from which the response will be buffered on disk instead of in memory.
	//
	// Default: 1048576 (1Mi).
	MemResponseBodyBytes *float64 `field:"optional" json:"memResponseBodyBytes" yaml:"memResponseBodyBytes"`
	// RetryExpression defines the retry conditions.
	//
	// It is a logical combination of functions with operators AND (&&) and OR (||). More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/buffering/#retryexpression
	RetryExpression *string `field:"optional" json:"retryExpression" yaml:"retryExpression"`
}

