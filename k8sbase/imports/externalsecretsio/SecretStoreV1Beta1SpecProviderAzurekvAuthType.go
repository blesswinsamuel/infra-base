// external-secretsio
package externalsecretsio


// Auth type defines how to authenticate to the keyvault service.
//
// Valid values are: - "ServicePrincipal" (default): Using a service principal (tenantId, clientId, clientSecret) - "ManagedIdentity": Using Managed Identity assigned to the pod (see aad-pod-identity).
type SecretStoreV1Beta1SpecProviderAzurekvAuthType string

const (
	// ServicePrincipal.
	SecretStoreV1Beta1SpecProviderAzurekvAuthType_SERVICE_PRINCIPAL SecretStoreV1Beta1SpecProviderAzurekvAuthType = "SERVICE_PRINCIPAL"
	// ManagedIdentity.
	SecretStoreV1Beta1SpecProviderAzurekvAuthType_MANAGED_IDENTITY SecretStoreV1Beta1SpecProviderAzurekvAuthType = "MANAGED_IDENTITY"
	// WorkloadIdentity.
	SecretStoreV1Beta1SpecProviderAzurekvAuthType_WORKLOAD_IDENTITY SecretStoreV1Beta1SpecProviderAzurekvAuthType = "WORKLOAD_IDENTITY"
)

