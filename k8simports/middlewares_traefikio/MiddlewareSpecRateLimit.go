package middlewares_traefikio


// RateLimit holds the rate limit configuration.
//
// This middleware ensures that services will receive a fair amount of requests, and allows one to define what fair is. More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/ratelimit/
type MiddlewareSpecRateLimit struct {
	// Average is the maximum rate, by default in requests/s, allowed for the given source.
	//
	// It defaults to 0, which means no rate limiting. The rate is actually defined by dividing Average by Period. So for a rate below 1req/s, one needs to define a Period larger than a second.
	Average *float64 `field:"optional" json:"average" yaml:"average"`
	// Burst is the maximum number of requests allowed to arrive in the same arbitrarily small period of time.
	//
	// It defaults to 1.
	Burst *float64 `field:"optional" json:"burst" yaml:"burst"`
	// Period, in combination with Average, defines the actual maximum rate, such as: r = Average / Period.
	//
	// It defaults to a second.
	Period MiddlewareSpecRateLimitPeriod `field:"optional" json:"period" yaml:"period"`
	// SourceCriterion defines what criterion is used to group requests as originating from a common source.
	//
	// If several strategies are defined at the same time, an error will be raised. If none are set, the default is to use the request's remote address field (as an ipStrategy).
	SourceCriterion *MiddlewareSpecRateLimitSourceCriterion `field:"optional" json:"sourceCriterion" yaml:"sourceCriterion"`
}

