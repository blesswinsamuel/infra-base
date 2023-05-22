package externalsecretsio


// AzureKV configures this store to sync secrets using Azure Key Vault provider.
type SecretStoreV1Beta1SpecProviderAzurekv struct {
	// Vault Url from which the secrets to be fetched from.
	VaultUrl *string `field:"required" json:"vaultUrl" yaml:"vaultUrl"`
	// Auth configures how the operator authenticates with Azure.
	//
	// Required for ServicePrincipal auth type.
	AuthSecretRef *SecretStoreV1Beta1SpecProviderAzurekvAuthSecretRef `field:"optional" json:"authSecretRef" yaml:"authSecretRef"`
	// Auth type defines how to authenticate to the keyvault service.
	//
	// Valid values are: - "ServicePrincipal" (default): Using a service principal (tenantId, clientId, clientSecret) - "ManagedIdentity": Using Managed Identity assigned to the pod (see aad-pod-identity).
	AuthType SecretStoreV1Beta1SpecProviderAzurekvAuthType `field:"optional" json:"authType" yaml:"authType"`
	// EnvironmentType specifies the Azure cloud environment endpoints to use for connecting and authenticating with Azure.
	//
	// By default it points to the public cloud AAD endpoint. The following endpoints are available, also see here: https://github.com/Azure/go-autorest/blob/main/autorest/azure/environments.go#L152 PublicCloud, USGovernmentCloud, ChinaCloud, GermanCloud
	EnvironmentType SecretStoreV1Beta1SpecProviderAzurekvEnvironmentType `field:"optional" json:"environmentType" yaml:"environmentType"`
	// If multiple Managed Identity is assigned to the pod, you can select the one to be used.
	IdentityId *string `field:"optional" json:"identityId" yaml:"identityId"`
	// ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.
	ServiceAccountRef *SecretStoreV1Beta1SpecProviderAzurekvServiceAccountRef `field:"optional" json:"serviceAccountRef" yaml:"serviceAccountRef"`
	// TenantID configures the Azure Tenant to send requests to.
	//
	// Required for ServicePrincipal auth type.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

