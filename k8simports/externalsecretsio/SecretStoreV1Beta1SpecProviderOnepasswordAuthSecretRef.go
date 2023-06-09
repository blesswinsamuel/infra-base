package externalsecretsio


// OnePasswordAuthSecretRef holds secret references for 1Password credentials.
type SecretStoreV1Beta1SpecProviderOnepasswordAuthSecretRef struct {
	// The ConnectToken is used for authentication to a 1Password Connect Server.
	ConnectTokenSecretRef *SecretStoreV1Beta1SpecProviderOnepasswordAuthSecretRefConnectTokenSecretRef `field:"required" json:"connectTokenSecretRef" yaml:"connectTokenSecretRef"`
}

