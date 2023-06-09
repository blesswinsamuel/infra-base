package middlewares_traefikio


// StripPrefixRegex holds the strip prefix regex middleware configuration.
//
// This middleware removes the matching prefixes from the URL path. More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/stripprefixregex/
type MiddlewareSpecStripPrefixRegex struct {
	// Regex defines the regular expression to match the path prefix from the request URL.
	Regex *[]*string `field:"optional" json:"regex" yaml:"regex"`
}

