package middlewares_traefikio


// IPWhiteList holds the IP whitelist middleware configuration.
//
// This middleware accepts / refuses requests based on the client IP. More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/ipwhitelist/
type MiddlewareSpecIpWhiteList struct {
	// IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP.
	//
	// More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/ipwhitelist/#ipstrategy
	IpStrategy *MiddlewareSpecIpWhiteListIpStrategy `field:"optional" json:"ipStrategy" yaml:"ipStrategy"`
	// SourceRange defines the set of allowed IPs (or ranges of allowed IPs by using CIDR notation).
	SourceRange *[]*string `field:"optional" json:"sourceRange" yaml:"sourceRange"`
}

