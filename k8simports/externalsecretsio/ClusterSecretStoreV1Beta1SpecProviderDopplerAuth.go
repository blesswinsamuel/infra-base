package externalsecretsio


// Auth configures how the Operator authenticates with the Doppler API.
type ClusterSecretStoreV1Beta1SpecProviderDopplerAuth struct {
	SecretRef *ClusterSecretStoreV1Beta1SpecProviderDopplerAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

