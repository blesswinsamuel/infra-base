package ingressroute_traefikio


// MiddlewareRef is a reference to a Middleware resource.
type IngressRouteSpecRoutesMiddlewares struct {
	// Name defines the name of the referenced Middleware resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace defines the namespace of the referenced Middleware resource.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

