// external-secretsio
package externalsecretsio


// AlibabaAuthSecretRef holds secret references for Alibaba credentials.
type ClusterSecretStoreV1Beta1SpecProviderAlibabaAuthSecretRef struct {
	// The AccessKeyID is used for authentication.
	AccessKeyIdSecretRef *ClusterSecretStoreV1Beta1SpecProviderAlibabaAuthSecretRefAccessKeyIdSecretRef `field:"required" json:"accessKeyIdSecretRef" yaml:"accessKeyIdSecretRef"`
	// The AccessKeySecret is used for authentication.
	AccessKeySecretSecretRef *ClusterSecretStoreV1Beta1SpecProviderAlibabaAuthSecretRefAccessKeySecretSecretRef `field:"required" json:"accessKeySecretSecretRef" yaml:"accessKeySecretSecretRef"`
}

