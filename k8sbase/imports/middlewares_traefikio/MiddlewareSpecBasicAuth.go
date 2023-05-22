package middlewares_traefikio


// BasicAuth holds the basic auth middleware configuration.
//
// This middleware restricts access to your services to known users. More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/basicauth/
type MiddlewareSpecBasicAuth struct {
	// HeaderField defines a header field to store the authenticated user.
	//
	// More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/basicauth/#headerfield
	HeaderField *string `field:"optional" json:"headerField" yaml:"headerField"`
	// Realm allows the protected resources on a server to be partitioned into a set of protection spaces, each with its own authentication scheme.
	//
	// Default: traefik.
	Realm *string `field:"optional" json:"realm" yaml:"realm"`
	// RemoveHeader sets the removeHeader option to true to remove the authorization header before forwarding the request to your service.
	//
	// Default: false.
	RemoveHeader *bool `field:"optional" json:"removeHeader" yaml:"removeHeader"`
	// Secret is the name of the referenced Kubernetes Secret containing user credentials.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}

