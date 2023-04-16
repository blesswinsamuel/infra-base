// traefikservices_traefikcontainous
package traefikservices_traefikcontainous

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// TraefikService is the CRD implementation of a Traefik Service.
//
// TraefikService object allows to: - Apply weight to Services on load-balancing - Mirror traffic on services More info: https://doc.traefik.io/traefik/v2.9/routing/providers/kubernetes-crd/#kind-traefikservice
type TraefikServiceProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// TraefikServiceSpec defines the desired state of a TraefikService.
	Spec *TraefikServiceSpec `field:"required" json:"spec" yaml:"spec"`
}

