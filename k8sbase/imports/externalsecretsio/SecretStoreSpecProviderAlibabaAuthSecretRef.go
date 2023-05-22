package externalsecretsio


// AlibabaAuthSecretRef holds secret references for Alibaba credentials.
type SecretStoreSpecProviderAlibabaAuthSecretRef struct {
	// The AccessKeyID is used for authentication.
	AccessKeyIdSecretRef *SecretStoreSpecProviderAlibabaAuthSecretRefAccessKeyIdSecretRef `field:"required" json:"accessKeyIdSecretRef" yaml:"accessKeyIdSecretRef"`
	// The AccessKeySecret is used for authentication.
	AccessKeySecretSecretRef *SecretStoreSpecProviderAlibabaAuthSecretRefAccessKeySecretSecretRef `field:"required" json:"accessKeySecretSecretRef" yaml:"accessKeySecretSecretRef"`
}

