package externalsecretsio


// AlibabaAuth contains a secretRef for credentials.
type ClusterSecretStoreV1Beta1SpecProviderAlibabaAuth struct {
	// AlibabaAuthSecretRef holds secret references for Alibaba credentials.
	SecretRef *ClusterSecretStoreV1Beta1SpecProviderAlibabaAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

