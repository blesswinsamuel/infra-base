package externalsecretsio


// Auth type defines how to authenticate to the keyvault service.
//
// Valid values are: - "ServicePrincipal" (default): Using a service principal (tenantId, clientId, clientSecret) - "ManagedIdentity": Using Managed Identity assigned to the pod (see aad-pod-identity).
type ClusterSecretStoreSpecProviderAzurekvAuthType string

const (
	// ServicePrincipal.
	ClusterSecretStoreSpecProviderAzurekvAuthType_SERVICE_PRINCIPAL ClusterSecretStoreSpecProviderAzurekvAuthType = "SERVICE_PRINCIPAL"
	// ManagedIdentity.
	ClusterSecretStoreSpecProviderAzurekvAuthType_MANAGED_IDENTITY ClusterSecretStoreSpecProviderAzurekvAuthType = "MANAGED_IDENTITY"
	// WorkloadIdentity.
	ClusterSecretStoreSpecProviderAzurekvAuthType_WORKLOAD_IDENTITY ClusterSecretStoreSpecProviderAzurekvAuthType = "WORKLOAD_IDENTITY"
)

