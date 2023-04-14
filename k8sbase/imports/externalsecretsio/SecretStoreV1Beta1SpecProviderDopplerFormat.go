// external-secretsio
package externalsecretsio


// Format enables the downloading of secrets as a file (string).
type SecretStoreV1Beta1SpecProviderDopplerFormat string

const (
	// json.
	SecretStoreV1Beta1SpecProviderDopplerFormat_JSON SecretStoreV1Beta1SpecProviderDopplerFormat = "JSON"
	// dotnet-json.
	SecretStoreV1Beta1SpecProviderDopplerFormat_DOTNET_JSON SecretStoreV1Beta1SpecProviderDopplerFormat = "DOTNET_JSON"
	// env.
	SecretStoreV1Beta1SpecProviderDopplerFormat_ENV SecretStoreV1Beta1SpecProviderDopplerFormat = "ENV"
	// yaml.
	SecretStoreV1Beta1SpecProviderDopplerFormat_YAML SecretStoreV1Beta1SpecProviderDopplerFormat = "YAML"
	// docker.
	SecretStoreV1Beta1SpecProviderDopplerFormat_DOCKER SecretStoreV1Beta1SpecProviderDopplerFormat = "DOCKER"
)

