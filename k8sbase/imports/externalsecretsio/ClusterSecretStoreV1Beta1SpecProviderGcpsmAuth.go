package externalsecretsio


// Auth defines the information necessary to authenticate against GCP.
type ClusterSecretStoreV1Beta1SpecProviderGcpsmAuth struct {
	SecretRef *ClusterSecretStoreV1Beta1SpecProviderGcpsmAuthSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	WorkloadIdentity *ClusterSecretStoreV1Beta1SpecProviderGcpsmAuthWorkloadIdentity `field:"optional" json:"workloadIdentity" yaml:"workloadIdentity"`
}

