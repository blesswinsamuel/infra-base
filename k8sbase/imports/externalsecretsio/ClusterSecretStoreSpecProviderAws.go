// external-secretsio
package externalsecretsio


// AWS configures this store to sync secrets using AWS Secret Manager provider.
type ClusterSecretStoreSpecProviderAws struct {
	// AWS Region to be used for the provider.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Service defines which service should be used to fetch the secrets.
	Service ClusterSecretStoreSpecProviderAwsService `field:"required" json:"service" yaml:"service"`
	// Auth defines the information necessary to authenticate against AWS if not set aws sdk will infer credentials from your environment see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials.
	Auth *ClusterSecretStoreSpecProviderAwsAuth `field:"optional" json:"auth" yaml:"auth"`
	// Role is a Role ARN which the SecretManager provider will assume.
	Role *string `field:"optional" json:"role" yaml:"role"`
}

