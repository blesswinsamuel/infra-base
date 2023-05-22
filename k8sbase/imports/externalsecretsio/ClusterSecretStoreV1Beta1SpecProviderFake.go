package externalsecretsio


// Fake configures a store with static key/value pairs.
type ClusterSecretStoreV1Beta1SpecProviderFake struct {
	Data *[]*ClusterSecretStoreV1Beta1SpecProviderFakeData `field:"required" json:"data" yaml:"data"`
}

