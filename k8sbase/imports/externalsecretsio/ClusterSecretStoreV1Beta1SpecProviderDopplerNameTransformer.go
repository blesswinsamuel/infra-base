package externalsecretsio


// Environment variable compatible name transforms that change secret names to a different format.
type ClusterSecretStoreV1Beta1SpecProviderDopplerNameTransformer string

const (
	// upper-camel.
	ClusterSecretStoreV1Beta1SpecProviderDopplerNameTransformer_UPPER_CAMEL ClusterSecretStoreV1Beta1SpecProviderDopplerNameTransformer = "UPPER_CAMEL"
	// camel.
	ClusterSecretStoreV1Beta1SpecProviderDopplerNameTransformer_CAMEL ClusterSecretStoreV1Beta1SpecProviderDopplerNameTransformer = "CAMEL"
	// lower-snake.
	ClusterSecretStoreV1Beta1SpecProviderDopplerNameTransformer_LOWER_SNAKE ClusterSecretStoreV1Beta1SpecProviderDopplerNameTransformer = "LOWER_SNAKE"
	// tf-var.
	ClusterSecretStoreV1Beta1SpecProviderDopplerNameTransformer_TF_VAR ClusterSecretStoreV1Beta1SpecProviderDopplerNameTransformer = "TF_VAR"
	// dotnet-env.
	ClusterSecretStoreV1Beta1SpecProviderDopplerNameTransformer_DOTNET_ENV ClusterSecretStoreV1Beta1SpecProviderDopplerNameTransformer = "DOTNET_ENV"
)

