package generatorsexternalsecretsio


// WorkloadIdentity uses Azure Workload Identity to authenticate with Azure.
type AcrAccessTokenSpecAuthWorkloadIdentity struct {
	// ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.
	ServiceAccountRef *AcrAccessTokenSpecAuthWorkloadIdentityServiceAccountRef `field:"optional" json:"serviceAccountRef" yaml:"serviceAccountRef"`
}

