package certmanagerio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A ClusterIssuer represents a certificate issuing authority which can be referenced as part of `issuerRef` fields.
//
// It is similar to an Issuer, however it is cluster-scoped and therefore can be referenced by resources that exist in *any* namespace, not just the same namespace as the referent.
type ClusterIssuerProps struct {
	// Desired state of the ClusterIssuer resource.
	Spec *ClusterIssuerSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

