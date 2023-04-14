// external-secretsio
package externalsecretsio


// Used to configure http retries if failed.
type SecretStoreSpecRetrySettings struct {
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	RetryInterval *string `field:"optional" json:"retryInterval" yaml:"retryInterval"`
}

