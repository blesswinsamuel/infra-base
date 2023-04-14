// external-secretsio
package externalsecretsio


// Auth type defines how to authenticate to the keyvault service.
//
// Valid values are: - "ServicePrincipal" (default): Using a service principal (tenantId, clientId, clientSecret) - "ManagedIdentity": Using Managed Identity assigned to the pod (see aad-pod-identity).
type SecretStoreSpecProviderAzurekvAuthType string

const (
	// ServicePrincipal.
	SecretStoreSpecProviderAzurekvAuthType_SERVICE_PRINCIPAL SecretStoreSpecProviderAzurekvAuthType = "SERVICE_PRINCIPAL"
	// ManagedIdentity.
	SecretStoreSpecProviderAzurekvAuthType_MANAGED_IDENTITY SecretStoreSpecProviderAzurekvAuthType = "MANAGED_IDENTITY"
	// WorkloadIdentity.
	SecretStoreSpecProviderAzurekvAuthType_WORKLOAD_IDENTITY SecretStoreSpecProviderAzurekvAuthType = "WORKLOAD_IDENTITY"
)

