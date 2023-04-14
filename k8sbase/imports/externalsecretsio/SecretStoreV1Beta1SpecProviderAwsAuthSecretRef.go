// external-secretsio
package externalsecretsio


// AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.
type SecretStoreV1Beta1SpecProviderAwsAuthSecretRef struct {
	// The AccessKeyID is used for authentication.
	AccessKeyIdSecretRef *SecretStoreV1Beta1SpecProviderAwsAuthSecretRefAccessKeyIdSecretRef `field:"optional" json:"accessKeyIdSecretRef" yaml:"accessKeyIdSecretRef"`
	// The SecretAccessKey is used for authentication.
	SecretAccessKeySecretRef *SecretStoreV1Beta1SpecProviderAwsAuthSecretRefSecretAccessKeySecretRef `field:"optional" json:"secretAccessKeySecretRef" yaml:"secretAccessKeySecretRef"`
	// The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html.
	SessionTokenSecretRef *SecretStoreV1Beta1SpecProviderAwsAuthSecretRefSessionTokenSecretRef `field:"optional" json:"sessionTokenSecretRef" yaml:"sessionTokenSecretRef"`
}

