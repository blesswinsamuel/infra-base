// external-secretsio
package externalsecretsio


// Service defines which service should be used to fetch the secrets.
type SecretStoreV1Beta1SpecProviderAwsService string

const (
	// SecretsManager.
	SecretStoreV1Beta1SpecProviderAwsService_SECRETS_MANAGER SecretStoreV1Beta1SpecProviderAwsService = "SECRETS_MANAGER"
	// ParameterStore.
	SecretStoreV1Beta1SpecProviderAwsService_PARAMETER_STORE SecretStoreV1Beta1SpecProviderAwsService = "PARAMETER_STORE"
)

