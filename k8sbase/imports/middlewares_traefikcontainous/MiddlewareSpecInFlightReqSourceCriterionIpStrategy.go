// middlewares_traefikcontainous
package middlewares_traefikcontainous


// IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP.
//
// More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/ipwhitelist/#ipstrategy
type MiddlewareSpecInFlightReqSourceCriterionIpStrategy struct {
	// Depth tells Traefik to use the X-Forwarded-For header and take the IP located at the depth position (starting from the right).
	Depth *float64 `field:"optional" json:"depth" yaml:"depth"`
	// ExcludedIPs configures Traefik to scan the X-Forwarded-For header and select the first IP not in the list.
	ExcludedIPs *[]*string `field:"optional" json:"excludedIPs" yaml:"excludedIPs"`
}

