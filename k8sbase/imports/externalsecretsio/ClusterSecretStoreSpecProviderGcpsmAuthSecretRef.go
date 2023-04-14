// external-secretsio
package externalsecretsio


type ClusterSecretStoreSpecProviderGcpsmAuthSecretRef struct {
	// The SecretAccessKey is used for authentication.
	SecretAccessKeySecretRef *ClusterSecretStoreSpecProviderGcpsmAuthSecretRefSecretAccessKeySecretRef `field:"optional" json:"secretAccessKeySecretRef" yaml:"secretAccessKeySecretRef"`
}

