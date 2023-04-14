// cert-managerio
package certmanagerio


// Use the AWS Route53 API to manage DNS01 challenge records.
type ClusterIssuerSpecAcmeSolversDns01Route53 struct {
	// Always set the region when using AccessKeyID and SecretAccessKey.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The AccessKeyID is used for authentication.
	//
	// Cannot be set when SecretAccessKeyID is set. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// The SecretAccessKey is used for authentication.
	//
	// If set, pull the AWS access key ID from a key within a Kubernetes Secret. Cannot be set when AccessKeyID is set. If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials
	AccessKeyIdSecretRef *ClusterIssuerSpecAcmeSolversDns01Route53AccessKeyIdSecretRef `field:"optional" json:"accessKeyIdSecretRef" yaml:"accessKeyIdSecretRef"`
	// If set, the provider will manage only this zone in Route53 and will not do an lookup using the route53:ListHostedZonesByName api call.
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// Role is a Role ARN which the Route53 provider will assume using either the explicit credentials AccessKeyID/SecretAccessKey or the inferred credentials from environment variables, shared credentials file or AWS Instance metadata.
	Role *string `field:"optional" json:"role" yaml:"role"`
	// The SecretAccessKey is used for authentication.
	//
	// If neither the Access Key nor Key ID are set, we fall-back to using env vars, shared credentials file or AWS Instance metadata, see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials
	SecretAccessKeySecretRef *ClusterIssuerSpecAcmeSolversDns01Route53SecretAccessKeySecretRef `field:"optional" json:"secretAccessKeySecretRef" yaml:"secretAccessKeySecretRef"`
}

