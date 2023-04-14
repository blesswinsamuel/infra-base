// middlewares_traefikcontainous
package middlewares_traefikcontainous


// MiddlewareSpec defines the desired state of a Middleware.
type MiddlewareSpec struct {
	// AddPrefix holds the add prefix middleware configuration.
	//
	// This middleware updates the path of a request before forwarding it. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/addprefix/
	AddPrefix *MiddlewareSpecAddPrefix `field:"optional" json:"addPrefix" yaml:"addPrefix"`
	// BasicAuth holds the basic auth middleware configuration.
	//
	// This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/basicauth/
	BasicAuth *MiddlewareSpecBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// Buffering holds the buffering middleware configuration.
	//
	// This middleware retries or limits the size of requests that can be forwarded to backends. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/buffering/#maxrequestbodybytes
	Buffering *MiddlewareSpecBuffering `field:"optional" json:"buffering" yaml:"buffering"`
	// Chain holds the configuration of the chain middleware.
	//
	// This middleware enables to define reusable combinations of other pieces of middleware. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/chain/
	Chain *MiddlewareSpecChain `field:"optional" json:"chain" yaml:"chain"`
	// CircuitBreaker holds the circuit breaker configuration.
	CircuitBreaker *MiddlewareSpecCircuitBreaker `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// Compress holds the compress middleware configuration.
	//
	// This middleware compresses responses before sending them to the client, using gzip compression. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/compress/
	Compress *MiddlewareSpecCompress `field:"optional" json:"compress" yaml:"compress"`
	// ContentType holds the content-type middleware configuration.
	//
	// This middleware exists to enable the correct behavior until at least the default one can be changed in a future version.
	ContentType *MiddlewareSpecContentType `field:"optional" json:"contentType" yaml:"contentType"`
	// DigestAuth holds the digest auth middleware configuration.
	//
	// This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/digestauth/
	DigestAuth *MiddlewareSpecDigestAuth `field:"optional" json:"digestAuth" yaml:"digestAuth"`
	// ErrorPage holds the custom error middleware configuration.
	//
	// This middleware returns a custom page in lieu of the default, according to configured ranges of HTTP Status codes. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/errorpages/
	Errors *MiddlewareSpecErrors `field:"optional" json:"errors" yaml:"errors"`
	// ForwardAuth holds the forward auth middleware configuration.
	//
	// This middleware delegates the request authentication to a Service. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/forwardauth/
	ForwardAuth *MiddlewareSpecForwardAuth `field:"optional" json:"forwardAuth" yaml:"forwardAuth"`
	// Headers holds the headers middleware configuration.
	//
	// This middleware manages the requests and responses headers. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/headers/#customrequestheaders
	Headers *MiddlewareSpecHeaders `field:"optional" json:"headers" yaml:"headers"`
	// InFlightReq holds the in-flight request middleware configuration.
	//
	// This middleware limits the number of requests being processed and served concurrently. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/inflightreq/
	InFlightReq *MiddlewareSpecInFlightReq `field:"optional" json:"inFlightReq" yaml:"inFlightReq"`
	// IPWhiteList holds the IP whitelist middleware configuration.
	//
	// This middleware accepts / refuses requests based on the client IP. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/ipwhitelist/
	IpWhiteList *MiddlewareSpecIpWhiteList `field:"optional" json:"ipWhiteList" yaml:"ipWhiteList"`
	// PassTLSClientCert holds the pass TLS client cert middleware configuration.
	//
	// This middleware adds the selected data from the passed client TLS certificate to a header. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/passtlsclientcert/
	PassTlsClientCert *MiddlewareSpecPassTlsClientCert `field:"optional" json:"passTlsClientCert" yaml:"passTlsClientCert"`
	// Plugin defines the middleware plugin configuration.
	//
	// More info: https://doc.traefik.io/traefik/plugins/
	Plugin *map[string]interface{} `field:"optional" json:"plugin" yaml:"plugin"`
	// RateLimit holds the rate limit configuration.
	//
	// This middleware ensures that services will receive a fair amount of requests, and allows one to define what fair is. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/ratelimit/
	RateLimit *MiddlewareSpecRateLimit `field:"optional" json:"rateLimit" yaml:"rateLimit"`
	// RedirectRegex holds the redirect regex middleware configuration.
	//
	// This middleware redirects a request using regex matching and replacement. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/redirectregex/#regex
	RedirectRegex *MiddlewareSpecRedirectRegex `field:"optional" json:"redirectRegex" yaml:"redirectRegex"`
	// RedirectScheme holds the redirect scheme middleware configuration.
	//
	// This middleware redirects requests from a scheme/port to another. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/redirectscheme/
	RedirectScheme *MiddlewareSpecRedirectScheme `field:"optional" json:"redirectScheme" yaml:"redirectScheme"`
	// ReplacePath holds the replace path middleware configuration.
	//
	// This middleware replaces the path of the request URL and store the original path in an X-Replaced-Path header. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/replacepath/
	ReplacePath *MiddlewareSpecReplacePath `field:"optional" json:"replacePath" yaml:"replacePath"`
	// ReplacePathRegex holds the replace path regex middleware configuration.
	//
	// This middleware replaces the path of a URL using regex matching and replacement. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/replacepathregex/
	ReplacePathRegex *MiddlewareSpecReplacePathRegex `field:"optional" json:"replacePathRegex" yaml:"replacePathRegex"`
	// Retry holds the retry middleware configuration.
	//
	// This middleware reissues requests a given number of times to a backend server if that server does not reply. As soon as the server answers, the middleware stops retrying, regardless of the response status. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/retry/
	Retry *MiddlewareSpecRetry `field:"optional" json:"retry" yaml:"retry"`
	// StripPrefix holds the strip prefix middleware configuration.
	//
	// This middleware removes the specified prefixes from the URL path. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/stripprefix/
	StripPrefix *MiddlewareSpecStripPrefix `field:"optional" json:"stripPrefix" yaml:"stripPrefix"`
	// StripPrefixRegex holds the strip prefix regex middleware configuration.
	//
	// This middleware removes the matching prefixes from the URL path. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/stripprefixregex/
	StripPrefixRegex *MiddlewareSpecStripPrefixRegex `field:"optional" json:"stripPrefixRegex" yaml:"stripPrefixRegex"`
}

