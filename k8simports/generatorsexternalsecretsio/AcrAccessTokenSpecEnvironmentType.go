package generatorsexternalsecretsio


// EnvironmentType specifies the Azure cloud environment endpoints to use for connecting and authenticating with Azure.
//
// By default it points to the public cloud AAD endpoint. The following endpoints are available, also see here: https://github.com/Azure/go-autorest/blob/main/autorest/azure/environments.go#L152 PublicCloud, USGovernmentCloud, ChinaCloud, GermanCloud
type AcrAccessTokenSpecEnvironmentType string

const (
	// PublicCloud.
	AcrAccessTokenSpecEnvironmentType_PUBLIC_CLOUD AcrAccessTokenSpecEnvironmentType = "PUBLIC_CLOUD"
	// USGovernmentCloud.
	AcrAccessTokenSpecEnvironmentType_US_GOVERNMENT_CLOUD AcrAccessTokenSpecEnvironmentType = "US_GOVERNMENT_CLOUD"
	// ChinaCloud.
	AcrAccessTokenSpecEnvironmentType_CHINA_CLOUD AcrAccessTokenSpecEnvironmentType = "CHINA_CLOUD"
	// GermanCloud.
	AcrAccessTokenSpecEnvironmentType_GERMAN_CLOUD AcrAccessTokenSpecEnvironmentType = "GERMAN_CLOUD"
)

