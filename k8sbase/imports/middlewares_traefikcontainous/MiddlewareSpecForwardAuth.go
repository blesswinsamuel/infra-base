// middlewares_traefikcontainous
package middlewares_traefikcontainous


// ForwardAuth holds the forward auth middleware configuration.
//
// This middleware delegates the request authentication to a Service. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/forwardauth/
type MiddlewareSpecForwardAuth struct {
	// Address defines the authentication server address.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// AuthRequestHeaders defines the list of the headers to copy from the request to the authentication server.
	//
	// If not set or empty then all request headers are passed.
	AuthRequestHeaders *[]*string `field:"optional" json:"authRequestHeaders" yaml:"authRequestHeaders"`
	// AuthResponseHeaders defines the list of headers to copy from the authentication server response and set on forwarded request, replacing any existing conflicting headers.
	AuthResponseHeaders *[]*string `field:"optional" json:"authResponseHeaders" yaml:"authResponseHeaders"`
	// AuthResponseHeadersRegex defines the regex to match headers to copy from the authentication server response and set on forwarded request, after stripping all headers that match the regex.
	//
	// More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/forwardauth/#authresponseheadersregex
	AuthResponseHeadersRegex *string `field:"optional" json:"authResponseHeadersRegex" yaml:"authResponseHeadersRegex"`
	// TLS defines the configuration used to secure the connection to the authentication server.
	Tls *MiddlewareSpecForwardAuthTls `field:"optional" json:"tls" yaml:"tls"`
	// TrustForwardHeader defines whether to trust (ie: forward) all X-Forwarded-* headers.
	TrustForwardHeader *bool `field:"optional" json:"trustForwardHeader" yaml:"trustForwardHeader"`
}

