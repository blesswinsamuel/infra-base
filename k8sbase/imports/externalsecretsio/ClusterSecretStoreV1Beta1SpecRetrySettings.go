package externalsecretsio


// Used to configure http retries if failed.
type ClusterSecretStoreV1Beta1SpecRetrySettings struct {
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	RetryInterval *string `field:"optional" json:"retryInterval" yaml:"retryInterval"`
}

