// generatorsexternal-secretsio
package generatorsexternalsecretsio


// AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.
type EcrAuthorizationTokenSpecAuthSecretRef struct {
	// The AccessKeyID is used for authentication.
	AccessKeyIdSecretRef *EcrAuthorizationTokenSpecAuthSecretRefAccessKeyIdSecretRef `field:"optional" json:"accessKeyIdSecretRef" yaml:"accessKeyIdSecretRef"`
	// The SecretAccessKey is used for authentication.
	SecretAccessKeySecretRef *EcrAuthorizationTokenSpecAuthSecretRefSecretAccessKeySecretRef `field:"optional" json:"secretAccessKeySecretRef" yaml:"secretAccessKeySecretRef"`
	// The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html.
	SessionTokenSecretRef *EcrAuthorizationTokenSpecAuthSecretRefSessionTokenSecretRef `field:"optional" json:"sessionTokenSecretRef" yaml:"sessionTokenSecretRef"`
}

