// external-secretsio
package externalsecretsio


// Fake configures a store with static key/value pairs.
type ClusterSecretStoreSpecProviderFake struct {
	Data *[]*ClusterSecretStoreSpecProviderFakeData `field:"required" json:"data" yaml:"data"`
}

