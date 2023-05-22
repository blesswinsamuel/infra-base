package serverstransports_traefikio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// ServersTransport is the CRD implementation of a ServersTransport.
//
// If no serversTransport is specified, the default@internal will be used. The default@internal serversTransport is created from the static configuration. More info: https://doc.traefik.io/traefik/v2.10/routing/services/#serverstransport_1
type ServersTransportProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// ServersTransportSpec defines the desired state of a ServersTransport.
	Spec *ServersTransportSpec `field:"required" json:"spec" yaml:"spec"`
}

