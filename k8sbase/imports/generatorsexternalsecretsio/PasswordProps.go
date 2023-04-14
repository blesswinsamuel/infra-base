// generatorsexternal-secretsio
package generatorsexternalsecretsio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Password generates a random password based on the configuration parameters in spec.
//
// You can specify the length, characterset and other attributes.
type PasswordProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// PasswordSpec controls the behavior of the password generator.
	Spec *PasswordSpec `field:"optional" json:"spec" yaml:"spec"`
}

