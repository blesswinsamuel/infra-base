// middlewares_traefikcontainous
package middlewares_traefikcontainous


// RedirectRegex holds the redirect regex middleware configuration.
//
// This middleware redirects a request using regex matching and replacement. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/redirectregex/#regex
type MiddlewareSpecRedirectRegex struct {
	// Permanent defines whether the redirection is permanent (301).
	Permanent *bool `field:"optional" json:"permanent" yaml:"permanent"`
	// Regex defines the regex used to match and capture elements from the request URL.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Replacement defines how to modify the URL to have the new target URL.
	Replacement *string `field:"optional" json:"replacement" yaml:"replacement"`
}

