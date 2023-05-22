package middlewares_traefikio


// SourceCriterion defines what criterion is used to group requests as originating from a common source.
//
// If several strategies are defined at the same time, an error will be raised. If none are set, the default is to use the request's remote address field (as an ipStrategy).
type MiddlewareSpecRateLimitSourceCriterion struct {
	// IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP.
	//
	// More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/ipwhitelist/#ipstrategy
	IpStrategy *MiddlewareSpecRateLimitSourceCriterionIpStrategy `field:"optional" json:"ipStrategy" yaml:"ipStrategy"`
	// RequestHeaderName defines the name of the header used to group incoming requests.
	RequestHeaderName *string `field:"optional" json:"requestHeaderName" yaml:"requestHeaderName"`
	// RequestHost defines whether to consider the request Host as the source.
	RequestHost *bool `field:"optional" json:"requestHost" yaml:"requestHost"`
}

