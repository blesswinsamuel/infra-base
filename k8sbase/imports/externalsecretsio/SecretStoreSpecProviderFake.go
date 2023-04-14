// external-secretsio
package externalsecretsio


// Fake configures a store with static key/value pairs.
type SecretStoreSpecProviderFake struct {
	Data *[]*SecretStoreSpecProviderFakeData `field:"required" json:"data" yaml:"data"`
}

