// acmecert-managerio
package acmecertmanagerio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Challenge is a type to represent a Challenge request with an ACME server.
type ChallengeProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	Spec *ChallengeSpec `field:"required" json:"spec" yaml:"spec"`
}

