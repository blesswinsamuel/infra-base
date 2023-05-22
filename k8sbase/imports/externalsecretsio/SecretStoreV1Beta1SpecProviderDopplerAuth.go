package externalsecretsio


// Auth configures how the Operator authenticates with the Doppler API.
type SecretStoreV1Beta1SpecProviderDopplerAuth struct {
	SecretRef *SecretStoreV1Beta1SpecProviderDopplerAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

