// external-secretsio
package externalsecretsio


// Environment variable compatible name transforms that change secret names to a different format.
type SecretStoreV1Beta1SpecProviderDopplerNameTransformer string

const (
	// upper-camel.
	SecretStoreV1Beta1SpecProviderDopplerNameTransformer_UPPER_CAMEL SecretStoreV1Beta1SpecProviderDopplerNameTransformer = "UPPER_CAMEL"
	// camel.
	SecretStoreV1Beta1SpecProviderDopplerNameTransformer_CAMEL SecretStoreV1Beta1SpecProviderDopplerNameTransformer = "CAMEL"
	// lower-snake.
	SecretStoreV1Beta1SpecProviderDopplerNameTransformer_LOWER_SNAKE SecretStoreV1Beta1SpecProviderDopplerNameTransformer = "LOWER_SNAKE"
	// tf-var.
	SecretStoreV1Beta1SpecProviderDopplerNameTransformer_TF_VAR SecretStoreV1Beta1SpecProviderDopplerNameTransformer = "TF_VAR"
	// dotnet-env.
	SecretStoreV1Beta1SpecProviderDopplerNameTransformer_DOTNET_ENV SecretStoreV1Beta1SpecProviderDopplerNameTransformer = "DOTNET_ENV"
)

