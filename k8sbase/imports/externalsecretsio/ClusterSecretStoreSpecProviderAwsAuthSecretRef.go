// external-secretsio
package externalsecretsio


// AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.
type ClusterSecretStoreSpecProviderAwsAuthSecretRef struct {
	// The AccessKeyID is used for authentication.
	AccessKeyIdSecretRef *ClusterSecretStoreSpecProviderAwsAuthSecretRefAccessKeyIdSecretRef `field:"optional" json:"accessKeyIdSecretRef" yaml:"accessKeyIdSecretRef"`
	// The SecretAccessKey is used for authentication.
	SecretAccessKeySecretRef *ClusterSecretStoreSpecProviderAwsAuthSecretRefSecretAccessKeySecretRef `field:"optional" json:"secretAccessKeySecretRef" yaml:"secretAccessKeySecretRef"`
}

