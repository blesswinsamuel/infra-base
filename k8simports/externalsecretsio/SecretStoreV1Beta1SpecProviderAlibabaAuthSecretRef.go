package externalsecretsio


// AlibabaAuthSecretRef holds secret references for Alibaba credentials.
type SecretStoreV1Beta1SpecProviderAlibabaAuthSecretRef struct {
	// The AccessKeyID is used for authentication.
	AccessKeyIdSecretRef *SecretStoreV1Beta1SpecProviderAlibabaAuthSecretRefAccessKeyIdSecretRef `field:"required" json:"accessKeyIdSecretRef" yaml:"accessKeyIdSecretRef"`
	// The AccessKeySecret is used for authentication.
	AccessKeySecretSecretRef *SecretStoreV1Beta1SpecProviderAlibabaAuthSecretRefAccessKeySecretSecretRef `field:"required" json:"accessKeySecretSecretRef" yaml:"accessKeySecretSecretRef"`
}

