package externalsecretsio


// Alibaba configures this store to sync secrets using Alibaba Cloud provider.
type SecretStoreV1Beta1SpecProviderAlibaba struct {
	// AlibabaAuth contains a secretRef for credentials.
	Auth *SecretStoreV1Beta1SpecProviderAlibabaAuth `field:"required" json:"auth" yaml:"auth"`
	// Alibaba Region to be used for the provider.
	RegionId *string `field:"required" json:"regionId" yaml:"regionId"`
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

