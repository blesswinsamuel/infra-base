// external-secretsio
package externalsecretsio


// AlibabaAuth contains a secretRef for credentials.
type SecretStoreV1Beta1SpecProviderAlibabaAuth struct {
	// AlibabaAuthSecretRef holds secret references for Alibaba credentials.
	SecretRef *SecretStoreV1Beta1SpecProviderAlibabaAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

