package externalsecretsio


// Auth defines the information necessary to authenticate against OnePassword Connect Server.
type ClusterSecretStoreV1Beta1SpecProviderOnepasswordAuth struct {
	// OnePasswordAuthSecretRef holds secret references for 1Password credentials.
	SecretRef *ClusterSecretStoreV1Beta1SpecProviderOnepasswordAuthSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

