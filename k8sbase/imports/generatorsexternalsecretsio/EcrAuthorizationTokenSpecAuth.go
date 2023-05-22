package generatorsexternalsecretsio


// Auth defines how to authenticate with AWS.
type EcrAuthorizationTokenSpecAuth struct {
	// Authenticate against AWS using service account tokens.
	Jwt *EcrAuthorizationTokenSpecAuthJwt `field:"optional" json:"jwt" yaml:"jwt"`
	// AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.
	SecretRef *EcrAuthorizationTokenSpecAuthSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

