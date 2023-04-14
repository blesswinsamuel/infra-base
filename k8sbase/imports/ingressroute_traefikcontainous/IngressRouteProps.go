// ingressroute_traefikcontainous
package ingressroute_traefikcontainous

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// IngressRoute is the CRD implementation of a Traefik HTTP Router.
type IngressRouteProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// IngressRouteSpec defines the desired state of IngressRoute.
	Spec *IngressRouteSpec `field:"required" json:"spec" yaml:"spec"`
}

