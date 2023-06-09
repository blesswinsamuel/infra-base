package middlewares_traefikio


// ContentType holds the content-type middleware configuration.
//
// This middleware exists to enable the correct behavior until at least the default one can be changed in a future version.
type MiddlewareSpecContentType struct {
	// AutoDetect specifies whether to let the `Content-Type` header, if it has not been set by the backend, be automatically set to a value derived from the contents of the response.
	//
	// As a proxy, the default behavior should be to leave the header alone, regardless of what the backend did with it. However, the historic default was to always auto-detect and set the header if it was nil, and it is going to be kept that way in order to support users currently relying on it.
	AutoDetect *bool `field:"optional" json:"autoDetect" yaml:"autoDetect"`
}

