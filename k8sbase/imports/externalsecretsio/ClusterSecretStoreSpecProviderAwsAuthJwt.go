// external-secretsio
package externalsecretsio


// Authenticate against AWS using service account tokens.
type ClusterSecretStoreSpecProviderAwsAuthJwt struct {
	// A reference to a ServiceAccount resource.
	ServiceAccountRef *ClusterSecretStoreSpecProviderAwsAuthJwtServiceAccountRef `field:"optional" json:"serviceAccountRef" yaml:"serviceAccountRef"`
}

