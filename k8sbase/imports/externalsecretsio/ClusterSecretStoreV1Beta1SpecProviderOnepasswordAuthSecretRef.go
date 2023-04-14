// external-secretsio
package externalsecretsio


// OnePasswordAuthSecretRef holds secret references for 1Password credentials.
type ClusterSecretStoreV1Beta1SpecProviderOnepasswordAuthSecretRef struct {
	// The ConnectToken is used for authentication to a 1Password Connect Server.
	ConnectTokenSecretRef *ClusterSecretStoreV1Beta1SpecProviderOnepasswordAuthSecretRefConnectTokenSecretRef `field:"required" json:"connectTokenSecretRef" yaml:"connectTokenSecretRef"`
}

