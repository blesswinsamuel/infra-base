// middlewarestcp_traefikcontainous
package middlewarestcp_traefikcontainous

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// MiddlewareTCP is the CRD implementation of a Traefik TCP middleware.
//
// More info: https://doc.traefik.io/traefik/v2.9/middlewares/overview/
type MiddlewareTcpProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// MiddlewareTCPSpec defines the desired state of a MiddlewareTCP.
	Spec *MiddlewareTcpSpec `field:"required" json:"spec" yaml:"spec"`
}

