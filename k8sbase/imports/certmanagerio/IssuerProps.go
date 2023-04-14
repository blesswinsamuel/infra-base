// cert-managerio
package certmanagerio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// An Issuer represents a certificate issuing authority which can be referenced as part of `issuerRef` fields.
//
// It is scoped to a single namespace and can therefore only be referenced by resources within the same namespace.
type IssuerProps struct {
	// Desired state of the Issuer resource.
	Spec *IssuerSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

