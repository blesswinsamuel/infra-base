// middlewares_traefikcontainous
package middlewares_traefikcontainous


// Chain holds the configuration of the chain middleware.
//
// This middleware enables to define reusable combinations of other pieces of middleware. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/chain/
type MiddlewareSpecChain struct {
	// Middlewares is the list of MiddlewareRef which composes the chain.
	Middlewares *[]*MiddlewareSpecChainMiddlewares `field:"optional" json:"middlewares" yaml:"middlewares"`
}

