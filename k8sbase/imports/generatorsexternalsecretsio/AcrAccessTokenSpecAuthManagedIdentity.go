// generatorsexternal-secretsio
package generatorsexternalsecretsio


// ManagedIdentity uses Azure Managed Identity to authenticate with Azure.
type AcrAccessTokenSpecAuthManagedIdentity struct {
	// If multiple Managed Identity is assigned to the pod, you can select the one to be used.
	IdentityId *string `field:"optional" json:"identityId" yaml:"identityId"`
}

