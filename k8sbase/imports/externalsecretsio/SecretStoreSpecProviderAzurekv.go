package externalsecretsio


// AzureKV configures this store to sync secrets using Azure Key Vault provider.
type SecretStoreSpecProviderAzurekv struct {
	// Vault Url from which the secrets to be fetched from.
	VaultUrl *string `field:"required" json:"vaultUrl" yaml:"vaultUrl"`
	// Auth configures how the operator authenticates with Azure.
	//
	// Required for ServicePrincipal auth type.
	AuthSecretRef *SecretStoreSpecProviderAzurekvAuthSecretRef `field:"optional" json:"authSecretRef" yaml:"authSecretRef"`
	// Auth type defines how to authenticate to the keyvault service.
	//
	// Valid values are: - "ServicePrincipal" (default): Using a service principal (tenantId, clientId, clientSecret) - "ManagedIdentity": Using Managed Identity assigned to the pod (see aad-pod-identity).
	AuthType SecretStoreSpecProviderAzurekvAuthType `field:"optional" json:"authType" yaml:"authType"`
	// If multiple Managed Identity is assigned to the pod, you can select the one to be used.
	IdentityId *string `field:"optional" json:"identityId" yaml:"identityId"`
	// ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.
	ServiceAccountRef *SecretStoreSpecProviderAzurekvServiceAccountRef `field:"optional" json:"serviceAccountRef" yaml:"serviceAccountRef"`
	// TenantID configures the Azure Tenant to send requests to.
	//
	// Required for ServicePrincipal auth type.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

