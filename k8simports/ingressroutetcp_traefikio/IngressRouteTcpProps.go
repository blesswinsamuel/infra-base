package ingressroutetcp_traefikio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// IngressRouteTCP is the CRD implementation of a Traefik TCP Router.
type IngressRouteTcpProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// IngressRouteTCPSpec defines the desired state of IngressRouteTCP.
	Spec *IngressRouteTcpSpec `field:"required" json:"spec" yaml:"spec"`
}

