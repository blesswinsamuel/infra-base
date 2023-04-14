// middlewares_traefikcontainous
package middlewares_traefikcontainous


// MiddlewareRef is a reference to a Middleware resource.
type MiddlewareSpecChainMiddlewares struct {
	// Name defines the name of the referenced Middleware resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace defines the namespace of the referenced Middleware resource.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

