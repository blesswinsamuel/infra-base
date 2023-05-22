package externalsecretsio


// Authenticate against AWS using service account tokens.
type ClusterSecretStoreV1Beta1SpecProviderAwsAuthJwt struct {
	// A reference to a ServiceAccount resource.
	ServiceAccountRef *ClusterSecretStoreV1Beta1SpecProviderAwsAuthJwtServiceAccountRef `field:"optional" json:"serviceAccountRef" yaml:"serviceAccountRef"`
}

