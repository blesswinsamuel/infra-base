package generatorsexternalsecretsio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// ECRAuthorizationTokenSpec uses the GetAuthorizationToken API to retrieve an authorization token.
//
// The authorization token is valid for 12 hours. The authorizationToken returned is a base64 encoded string that can be decoded and used in a docker login command to authenticate to a registry. For more information, see Registry authentication (https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_auth) in the Amazon Elastic Container Registry User Guide.
type EcrAuthorizationTokenProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	Spec *EcrAuthorizationTokenSpec `field:"optional" json:"spec" yaml:"spec"`
}

