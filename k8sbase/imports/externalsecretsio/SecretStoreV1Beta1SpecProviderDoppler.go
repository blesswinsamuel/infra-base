package externalsecretsio


// Doppler configures this store to sync secrets using the Doppler provider.
type SecretStoreV1Beta1SpecProviderDoppler struct {
	// Auth configures how the Operator authenticates with the Doppler API.
	Auth *SecretStoreV1Beta1SpecProviderDopplerAuth `field:"required" json:"auth" yaml:"auth"`
	// Doppler config (required if not using a Service Token).
	Config *string `field:"optional" json:"config" yaml:"config"`
	// Format enables the downloading of secrets as a file (string).
	Format SecretStoreV1Beta1SpecProviderDopplerFormat `field:"optional" json:"format" yaml:"format"`
	// Environment variable compatible name transforms that change secret names to a different format.
	NameTransformer SecretStoreV1Beta1SpecProviderDopplerNameTransformer `field:"optional" json:"nameTransformer" yaml:"nameTransformer"`
	// Doppler project (required if not using a Service Token).
	Project *string `field:"optional" json:"project" yaml:"project"`
}

