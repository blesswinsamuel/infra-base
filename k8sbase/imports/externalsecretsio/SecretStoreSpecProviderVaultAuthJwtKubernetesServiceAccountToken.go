// external-secretsio
package externalsecretsio


// Optional ServiceAccountToken specifies the Kubernetes service account for which to request a token for with the `TokenRequest` API.
type SecretStoreSpecProviderVaultAuthJwtKubernetesServiceAccountToken struct {
	// Service account field containing the name of a kubernetes ServiceAccount.
	ServiceAccountRef *SecretStoreSpecProviderVaultAuthJwtKubernetesServiceAccountTokenServiceAccountRef `field:"required" json:"serviceAccountRef" yaml:"serviceAccountRef"`
	// Optional audiences field that will be used to request a temporary Kubernetes service account token for the service account referenced by `serviceAccountRef`.
	//
	// Defaults to a single audience `vault` it not specified.
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
	// Optional expiration time in seconds that will be used to request a temporary Kubernetes service account token for the service account referenced by `serviceAccountRef`.
	//
	// Defaults to 10 minutes.
	ExpirationSeconds *float64 `field:"optional" json:"expirationSeconds" yaml:"expirationSeconds"`
}

