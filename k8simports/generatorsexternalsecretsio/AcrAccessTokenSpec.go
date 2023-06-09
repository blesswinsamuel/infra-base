package generatorsexternalsecretsio


// ACRAccessTokenSpec defines how to generate the access token e.g. how to authenticate and which registry to use. see: https://github.com/Azure/acr/blob/main/docs/AAD-OAuth.md#overview.
type AcrAccessTokenSpec struct {
	Auth *AcrAccessTokenSpecAuth `field:"required" json:"auth" yaml:"auth"`
	// the domain name of the ACR registry e.g. foobarexample.azurecr.io.
	Registry *string `field:"required" json:"registry" yaml:"registry"`
	// EnvironmentType specifies the Azure cloud environment endpoints to use for connecting and authenticating with Azure.
	//
	// By default it points to the public cloud AAD endpoint. The following endpoints are available, also see here: https://github.com/Azure/go-autorest/blob/main/autorest/azure/environments.go#L152 PublicCloud, USGovernmentCloud, ChinaCloud, GermanCloud
	EnvironmentType AcrAccessTokenSpecEnvironmentType `field:"optional" json:"environmentType" yaml:"environmentType"`
	// Define the scope for the access token, e.g. pull/push access for a repository. if not provided it will return a refresh token that has full scope. Note: you need to pin it down to the repository level, there is no wildcard available. examples: repository:my-repository:pull,push repository:my-repository:pull see docs for details: https://docs.docker.com/registry/spec/auth/scope/.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// TenantID configures the Azure Tenant to send requests to.
	//
	// Required for ServicePrincipal auth type.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

