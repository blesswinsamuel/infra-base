package externalsecretsio


// AlibabaAuthSecretRef holds secret references for Alibaba credentials.
type ClusterSecretStoreSpecProviderAlibabaAuthSecretRef struct {
	// The AccessKeyID is used for authentication.
	AccessKeyIdSecretRef *ClusterSecretStoreSpecProviderAlibabaAuthSecretRefAccessKeyIdSecretRef `field:"required" json:"accessKeyIdSecretRef" yaml:"accessKeyIdSecretRef"`
	// The AccessKeySecret is used for authentication.
	AccessKeySecretSecretRef *ClusterSecretStoreSpecProviderAlibabaAuthSecretRefAccessKeySecretSecretRef `field:"required" json:"accessKeySecretSecretRef" yaml:"accessKeySecretSecretRef"`
}

