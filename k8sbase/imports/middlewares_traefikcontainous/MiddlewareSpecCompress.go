// middlewares_traefikcontainous
package middlewares_traefikcontainous


// Compress holds the compress middleware configuration.
//
// This middleware compresses responses before sending them to the client, using gzip compression. More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/compress/
type MiddlewareSpecCompress struct {
	// ExcludedContentTypes defines the list of content types to compare the Content-Type header of the incoming requests and responses before compressing.
	ExcludedContentTypes *[]*string `field:"optional" json:"excludedContentTypes" yaml:"excludedContentTypes"`
	// MinResponseBodyBytes defines the minimum amount of bytes a response body must have to be compressed.
	//
	// Default: 1024.
	MinResponseBodyBytes *float64 `field:"optional" json:"minResponseBodyBytes" yaml:"minResponseBodyBytes"`
}

