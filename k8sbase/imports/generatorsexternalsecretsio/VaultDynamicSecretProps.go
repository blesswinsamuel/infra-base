// generatorsexternal-secretsio
package generatorsexternalsecretsio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type VaultDynamicSecretProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	Spec *VaultDynamicSecretSpec `field:"optional" json:"spec" yaml:"spec"`
}

