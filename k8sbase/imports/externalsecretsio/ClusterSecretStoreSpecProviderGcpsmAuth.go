// external-secretsio
package externalsecretsio


// Auth defines the information necessary to authenticate against GCP.
type ClusterSecretStoreSpecProviderGcpsmAuth struct {
	SecretRef *ClusterSecretStoreSpecProviderGcpsmAuthSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	WorkloadIdentity *ClusterSecretStoreSpecProviderGcpsmAuthWorkloadIdentity `field:"optional" json:"workloadIdentity" yaml:"workloadIdentity"`
}

