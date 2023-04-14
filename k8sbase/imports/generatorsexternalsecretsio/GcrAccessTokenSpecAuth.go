// generatorsexternal-secretsio
package generatorsexternalsecretsio


// Auth defines the means for authenticating with GCP.
type GcrAccessTokenSpecAuth struct {
	SecretRef *GcrAccessTokenSpecAuthSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	WorkloadIdentity *GcrAccessTokenSpecAuthWorkloadIdentity `field:"optional" json:"workloadIdentity" yaml:"workloadIdentity"`
}

