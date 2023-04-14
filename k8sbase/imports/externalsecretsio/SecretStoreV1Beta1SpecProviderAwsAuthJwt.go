// external-secretsio
package externalsecretsio


// Authenticate against AWS using service account tokens.
type SecretStoreV1Beta1SpecProviderAwsAuthJwt struct {
	// A reference to a ServiceAccount resource.
	ServiceAccountRef *SecretStoreV1Beta1SpecProviderAwsAuthJwtServiceAccountRef `field:"optional" json:"serviceAccountRef" yaml:"serviceAccountRef"`
}

