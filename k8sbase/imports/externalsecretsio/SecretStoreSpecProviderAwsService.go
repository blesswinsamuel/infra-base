package externalsecretsio


// Service defines which service should be used to fetch the secrets.
type SecretStoreSpecProviderAwsService string

const (
	// SecretsManager.
	SecretStoreSpecProviderAwsService_SECRETS_MANAGER SecretStoreSpecProviderAwsService = "SECRETS_MANAGER"
	// ParameterStore.
	SecretStoreSpecProviderAwsService_PARAMETER_STORE SecretStoreSpecProviderAwsService = "PARAMETER_STORE"
)

