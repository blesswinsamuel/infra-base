package externalsecretsio


// EnvironmentType specifies the Azure cloud environment endpoints to use for connecting and authenticating with Azure.
//
// By default it points to the public cloud AAD endpoint. The following endpoints are available, also see here: https://github.com/Azure/go-autorest/blob/main/autorest/azure/environments.go#L152 PublicCloud, USGovernmentCloud, ChinaCloud, GermanCloud
type ClusterSecretStoreV1Beta1SpecProviderAzurekvEnvironmentType string

const (
	// PublicCloud.
	ClusterSecretStoreV1Beta1SpecProviderAzurekvEnvironmentType_PUBLIC_CLOUD ClusterSecretStoreV1Beta1SpecProviderAzurekvEnvironmentType = "PUBLIC_CLOUD"
	// USGovernmentCloud.
	ClusterSecretStoreV1Beta1SpecProviderAzurekvEnvironmentType_US_GOVERNMENT_CLOUD ClusterSecretStoreV1Beta1SpecProviderAzurekvEnvironmentType = "US_GOVERNMENT_CLOUD"
	// ChinaCloud.
	ClusterSecretStoreV1Beta1SpecProviderAzurekvEnvironmentType_CHINA_CLOUD ClusterSecretStoreV1Beta1SpecProviderAzurekvEnvironmentType = "CHINA_CLOUD"
	// GermanCloud.
	ClusterSecretStoreV1Beta1SpecProviderAzurekvEnvironmentType_GERMAN_CLOUD ClusterSecretStoreV1Beta1SpecProviderAzurekvEnvironmentType = "GERMAN_CLOUD"
)

