// middlewares_traefikcontainous
package middlewares_traefikcontainous

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Middleware is the CRD implementation of a Traefik Middleware.
//
// More info: https://doc.traefik.io/traefik/v2.9/middlewares/http/overview/
type MiddlewareProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// MiddlewareSpec defines the desired state of a Middleware.
	Spec *MiddlewareSpec `field:"required" json:"spec" yaml:"spec"`
}

