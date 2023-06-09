package middlewares_traefikio


// AddPrefix holds the add prefix middleware configuration.
//
// This middleware updates the path of a request before forwarding it. More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/addprefix/
type MiddlewareSpecAddPrefix struct {
	// Prefix is the string to add before the current path in the requested URL.
	//
	// It should include a leading slash (/).
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

