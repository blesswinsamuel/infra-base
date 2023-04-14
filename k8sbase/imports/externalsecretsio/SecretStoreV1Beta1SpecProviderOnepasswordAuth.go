// external-secretsio
package externalsecretsio


// Auth defines the information necessary to authenticate against OnePassword Connect Server.
type SecretStoreV1Beta1SpecProviderOnepasswordAuth struct {
	// OnePasswordAuthSecretRef holds secret references for 1Password credentials.
	SecretRef *SecretStoreV1Beta1SpecProviderOnepasswordAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

