package externalsecretsio


// Auth defines the information necessary to authenticate against AWS if not set aws sdk will infer credentials from your environment see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials.
type SecretStoreSpecProviderAwsAuth struct {
	// Authenticate against AWS using service account tokens.
	Jwt *SecretStoreSpecProviderAwsAuthJwt `field:"optional" json:"jwt" yaml:"jwt"`
	// AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.
	SecretRef *SecretStoreSpecProviderAwsAuthSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

