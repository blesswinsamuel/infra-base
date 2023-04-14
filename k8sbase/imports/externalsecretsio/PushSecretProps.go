// external-secretsio
package externalsecretsio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type PushSecretProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// PushSecretSpec configures the behavior of the PushSecret.
	Spec *PushSecretSpec `field:"optional" json:"spec" yaml:"spec"`
}

