// external-secretsio
package externalsecretsio


// Auth defines the information necessary to authenticate against GCP.
type SecretStoreSpecProviderGcpsmAuth struct {
	SecretRef *SecretStoreSpecProviderGcpsmAuthSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	WorkloadIdentity *SecretStoreSpecProviderGcpsmAuthWorkloadIdentity `field:"optional" json:"workloadIdentity" yaml:"workloadIdentity"`
}

