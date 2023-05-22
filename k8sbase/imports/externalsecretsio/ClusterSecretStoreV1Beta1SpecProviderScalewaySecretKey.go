package externalsecretsio


// SecretKey is the non-secret part of the api key.
type ClusterSecretStoreV1Beta1SpecProviderScalewaySecretKey struct {
	// SecretRef references a key in a secret that will be used as value.
	SecretRef *ClusterSecretStoreV1Beta1SpecProviderScalewaySecretKeySecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	// Value can be specified directly to set a value without using a secret.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

