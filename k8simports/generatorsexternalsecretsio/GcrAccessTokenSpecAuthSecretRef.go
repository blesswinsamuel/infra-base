package generatorsexternalsecretsio


type GcrAccessTokenSpecAuthSecretRef struct {
	// The SecretAccessKey is used for authentication.
	SecretAccessKeySecretRef *GcrAccessTokenSpecAuthSecretRefSecretAccessKeySecretRef `field:"optional" json:"secretAccessKeySecretRef" yaml:"secretAccessKeySecretRef"`
}

