package generatorsexternalsecretsio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// ACRAccessToken returns a Azure Container Registry token that can be used for pushing/pulling images.
//
// Note: by default it will return an ACR Refresh Token with full access (depending on the identity). This can be scoped down to the repository level using .spec.scope. In case scope is defined it will return an ACR Access Token.
// See docs: https://github.com/Azure/acr/blob/main/docs/AAD-OAuth.md
type AcrAccessTokenProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// ACRAccessTokenSpec defines how to generate the access token e.g. how to authenticate and which registry to use. see: https://github.com/Azure/acr/blob/main/docs/AAD-OAuth.md#overview.
	Spec *AcrAccessTokenSpec `field:"optional" json:"spec" yaml:"spec"`
}

