package middlewares_traefikio


// InFlightReq holds the in-flight request middleware configuration.
//
// This middleware limits the number of requests being processed and served concurrently. More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/inflightreq/
type MiddlewareSpecInFlightReq struct {
	// Amount defines the maximum amount of allowed simultaneous in-flight request.
	//
	// The middleware responds with HTTP 429 Too Many Requests if there are already amount requests in progress (based on the same sourceCriterion strategy).
	Amount *float64 `field:"optional" json:"amount" yaml:"amount"`
	// SourceCriterion defines what criterion is used to group requests as originating from a common source.
	//
	// If several strategies are defined at the same time, an error will be raised. If none are set, the default is to use the requestHost. More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/inflightreq/#sourcecriterion
	SourceCriterion *MiddlewareSpecInFlightReqSourceCriterion `field:"optional" json:"sourceCriterion" yaml:"sourceCriterion"`
}

