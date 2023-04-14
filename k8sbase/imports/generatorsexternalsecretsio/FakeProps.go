// generatorsexternal-secretsio
package generatorsexternalsecretsio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Fake generator is used for testing.
//
// It lets you define a static set of credentials that is always returned.
type FakeProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// FakeSpec contains the static data.
	Spec *FakeSpec `field:"optional" json:"spec" yaml:"spec"`
}

