package middlewares_traefikio


// ReplacePathRegex holds the replace path regex middleware configuration.
//
// This middleware replaces the path of a URL using regex matching and replacement. More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/replacepathregex/
type MiddlewareSpecReplacePathRegex struct {
	// Regex defines the regular expression used to match and capture the path from the request URL.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Replacement defines the replacement path format, which can include captured variables.
	Replacement *string `field:"optional" json:"replacement" yaml:"replacement"`
}

