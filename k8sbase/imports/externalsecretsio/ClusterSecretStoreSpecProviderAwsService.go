package externalsecretsio


// Service defines which service should be used to fetch the secrets.
type ClusterSecretStoreSpecProviderAwsService string

const (
	// SecretsManager.
	ClusterSecretStoreSpecProviderAwsService_SECRETS_MANAGER ClusterSecretStoreSpecProviderAwsService = "SECRETS_MANAGER"
	// ParameterStore.
	ClusterSecretStoreSpecProviderAwsService_PARAMETER_STORE ClusterSecretStoreSpecProviderAwsService = "PARAMETER_STORE"
)

