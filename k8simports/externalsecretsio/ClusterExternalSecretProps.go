package externalsecretsio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// ClusterExternalSecret is the Schema for the clusterexternalsecrets API.
type ClusterExternalSecretProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// ClusterExternalSecretSpec defines the desired state of ClusterExternalSecret.
	Spec *ClusterExternalSecretSpec `field:"optional" json:"spec" yaml:"spec"`
}

