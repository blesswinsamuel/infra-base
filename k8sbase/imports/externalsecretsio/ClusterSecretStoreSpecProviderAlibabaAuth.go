// external-secretsio
package externalsecretsio


// AlibabaAuth contains a secretRef for credentials.
type ClusterSecretStoreSpecProviderAlibabaAuth struct {
	// AlibabaAuthSecretRef holds secret references for Alibaba credentials.
	SecretRef *ClusterSecretStoreSpecProviderAlibabaAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

