package generatorsexternalsecretsio


type EcrAuthorizationTokenSpec struct {
	// Region specifies the region to operate in.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Auth defines how to authenticate with AWS.
	Auth *EcrAuthorizationTokenSpecAuth `field:"optional" json:"auth" yaml:"auth"`
	// You can assume a role before making calls to the desired AWS service.
	Role *string `field:"optional" json:"role" yaml:"role"`
}

