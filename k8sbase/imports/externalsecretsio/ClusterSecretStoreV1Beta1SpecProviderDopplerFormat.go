// external-secretsio
package externalsecretsio


// Format enables the downloading of secrets as a file (string).
type ClusterSecretStoreV1Beta1SpecProviderDopplerFormat string

const (
	// json.
	ClusterSecretStoreV1Beta1SpecProviderDopplerFormat_JSON ClusterSecretStoreV1Beta1SpecProviderDopplerFormat = "JSON"
	// dotnet-json.
	ClusterSecretStoreV1Beta1SpecProviderDopplerFormat_DOTNET_JSON ClusterSecretStoreV1Beta1SpecProviderDopplerFormat = "DOTNET_JSON"
	// env.
	ClusterSecretStoreV1Beta1SpecProviderDopplerFormat_ENV ClusterSecretStoreV1Beta1SpecProviderDopplerFormat = "ENV"
	// yaml.
	ClusterSecretStoreV1Beta1SpecProviderDopplerFormat_YAML ClusterSecretStoreV1Beta1SpecProviderDopplerFormat = "YAML"
	// docker.
	ClusterSecretStoreV1Beta1SpecProviderDopplerFormat_DOCKER ClusterSecretStoreV1Beta1SpecProviderDopplerFormat = "DOCKER"
)

