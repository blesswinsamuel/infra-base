// traefikservices_traefikcontainous
package traefikservices_traefikcontainous


// ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.
type TraefikServiceSpecMirroringResponseForwarding struct {
	// FlushInterval defines the interval, in milliseconds, in between flushes to the client while copying the response body.
	//
	// A negative value means to flush immediately after each write to the client. This configuration is ignored when ReverseProxy recognizes a response as a streaming response; for such responses, writes are flushed to the client immediately. Default: 100ms
	FlushInterval *string `field:"optional" json:"flushInterval" yaml:"flushInterval"`
}

