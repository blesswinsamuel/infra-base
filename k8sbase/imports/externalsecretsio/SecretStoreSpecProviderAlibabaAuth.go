// external-secretsio
package externalsecretsio


// AlibabaAuth contains a secretRef for credentials.
type SecretStoreSpecProviderAlibabaAuth struct {
	// AlibabaAuthSecretRef holds secret references for Alibaba credentials.
	SecretRef *SecretStoreSpecProviderAlibabaAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

