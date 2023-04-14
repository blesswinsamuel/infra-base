// external-secretsio
package externalsecretsio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// SecretStore represents a secure external location for storing secrets, which can be referenced as part of `storeRef` fields.
type SecretStoreProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// SecretStoreSpec defines the desired state of SecretStore.
	Spec *SecretStoreSpec `field:"optional" json:"spec" yaml:"spec"`
}

