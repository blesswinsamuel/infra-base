// middlewares_traefikcontainous
package middlewares_traefikcontainous


// StripPrefixRegex holds the strip prefix regex middleware configuration.
//
// This middleware removes the matching prefixes from the URL path. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/stripprefixregex/
type MiddlewareSpecStripPrefixRegex struct {
	// Regex defines the regular expression to match the path prefix from the request URL.
	Regex *[]*string `field:"optional" json:"regex" yaml:"regex"`
}

