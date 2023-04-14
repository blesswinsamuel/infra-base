// external-secretsio
package externalsecretsio


// Auth type defines how to authenticate to the keyvault service.
//
// Valid values are: - "ServicePrincipal" (default): Using a service principal (tenantId, clientId, clientSecret) - "ManagedIdentity": Using Managed Identity assigned to the pod (see aad-pod-identity).
type ClusterSecretStoreV1Beta1SpecProviderAzurekvAuthType string

const (
	// ServicePrincipal.
	ClusterSecretStoreV1Beta1SpecProviderAzurekvAuthType_SERVICE_PRINCIPAL ClusterSecretStoreV1Beta1SpecProviderAzurekvAuthType = "SERVICE_PRINCIPAL"
	// ManagedIdentity.
	ClusterSecretStoreV1Beta1SpecProviderAzurekvAuthType_MANAGED_IDENTITY ClusterSecretStoreV1Beta1SpecProviderAzurekvAuthType = "MANAGED_IDENTITY"
	// WorkloadIdentity.
	ClusterSecretStoreV1Beta1SpecProviderAzurekvAuthType_WORKLOAD_IDENTITY ClusterSecretStoreV1Beta1SpecProviderAzurekvAuthType = "WORKLOAD_IDENTITY"
)

