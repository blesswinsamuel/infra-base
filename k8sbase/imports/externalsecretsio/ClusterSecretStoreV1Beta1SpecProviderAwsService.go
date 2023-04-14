// external-secretsio
package externalsecretsio


// Service defines which service should be used to fetch the secrets.
type ClusterSecretStoreV1Beta1SpecProviderAwsService string

const (
	// SecretsManager.
	ClusterSecretStoreV1Beta1SpecProviderAwsService_SECRETS_MANAGER ClusterSecretStoreV1Beta1SpecProviderAwsService = "SECRETS_MANAGER"
	// ParameterStore.
	ClusterSecretStoreV1Beta1SpecProviderAwsService_PARAMETER_STORE ClusterSecretStoreV1Beta1SpecProviderAwsService = "PARAMETER_STORE"
)

