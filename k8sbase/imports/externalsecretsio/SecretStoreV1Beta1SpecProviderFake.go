package externalsecretsio


// Fake configures a store with static key/value pairs.
type SecretStoreV1Beta1SpecProviderFake struct {
	Data *[]*SecretStoreV1Beta1SpecProviderFakeData `field:"required" json:"data" yaml:"data"`
}

