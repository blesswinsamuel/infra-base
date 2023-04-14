// generatorsexternal-secretsio
package generatorsexternalsecretsio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// GCRAccessToken generates an GCP access token that can be used to authenticate with GCR.
type GcrAccessTokenProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	Spec *GcrAccessTokenSpec `field:"optional" json:"spec" yaml:"spec"`
}

