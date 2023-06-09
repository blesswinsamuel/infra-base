package generatorsexternalsecretsio


// Authenticate against AWS using service account tokens.
type EcrAuthorizationTokenSpecAuthJwt struct {
	// A reference to a ServiceAccount resource.
	ServiceAccountRef *EcrAuthorizationTokenSpecAuthJwtServiceAccountRef `field:"optional" json:"serviceAccountRef" yaml:"serviceAccountRef"`
}

