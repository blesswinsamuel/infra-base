package externalsecretsio


// Auth defines the information necessary to authenticate against GCP.
type SecretStoreV1Beta1SpecProviderGcpsmAuth struct {
	SecretRef *SecretStoreV1Beta1SpecProviderGcpsmAuthSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	WorkloadIdentity *SecretStoreV1Beta1SpecProviderGcpsmAuthWorkloadIdentity `field:"optional" json:"workloadIdentity" yaml:"workloadIdentity"`
}

