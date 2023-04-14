// external-secretsio
package externalsecretsio


// Auth defines the information necessary to authenticate against AWS if not set aws sdk will infer credentials from your environment see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials.
type SecretStoreV1Beta1SpecProviderAwsAuth struct {
	// Authenticate against AWS using service account tokens.
	Jwt *SecretStoreV1Beta1SpecProviderAwsAuthJwt `field:"optional" json:"jwt" yaml:"jwt"`
	// AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.
	SecretRef *SecretStoreV1Beta1SpecProviderAwsAuthSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

