package ingressrouteudp_traefikio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// IngressRouteUDP is a CRD implementation of a Traefik UDP Router.
type IngressRouteUdpProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// IngressRouteUDPSpec defines the desired state of a IngressRouteUDP.
	Spec *IngressRouteUdpSpec `field:"required" json:"spec" yaml:"spec"`
}

