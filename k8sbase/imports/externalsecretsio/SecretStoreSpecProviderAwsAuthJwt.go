// external-secretsio
package externalsecretsio


// Authenticate against AWS using service account tokens.
type SecretStoreSpecProviderAwsAuthJwt struct {
	// A reference to a ServiceAccount resource.
	ServiceAccountRef *SecretStoreSpecProviderAwsAuthJwtServiceAccountRef `field:"optional" json:"serviceAccountRef" yaml:"serviceAccountRef"`
}

