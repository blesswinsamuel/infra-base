// external-secretsio
package externalsecretsio


// Auth configures how secret-manager authenticates with the IBM secrets manager.
type SecretStoreSpecProviderIbmAuth struct {
	SecretRef *SecretStoreSpecProviderIbmAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

