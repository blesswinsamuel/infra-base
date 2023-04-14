// generatorsexternal-secretsio
package generatorsexternalsecretsio


type AcrAccessTokenSpecAuth struct {
	// ManagedIdentity uses Azure Managed Identity to authenticate with Azure.
	ManagedIdentity *AcrAccessTokenSpecAuthManagedIdentity `field:"optional" json:"managedIdentity" yaml:"managedIdentity"`
	// ServicePrincipal uses Azure Service Principal credentials to authenticate with Azure.
	ServicePrincipal *AcrAccessTokenSpecAuthServicePrincipal `field:"optional" json:"servicePrincipal" yaml:"servicePrincipal"`
	// WorkloadIdentity uses Azure Workload Identity to authenticate with Azure.
	WorkloadIdentity *AcrAccessTokenSpecAuthWorkloadIdentity `field:"optional" json:"workloadIdentity" yaml:"workloadIdentity"`
}

