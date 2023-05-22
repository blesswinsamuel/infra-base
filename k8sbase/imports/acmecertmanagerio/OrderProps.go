package acmecertmanagerio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Order is a type to represent an Order with an ACME server.
type OrderProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	Spec *OrderSpec `field:"required" json:"spec" yaml:"spec"`
}

