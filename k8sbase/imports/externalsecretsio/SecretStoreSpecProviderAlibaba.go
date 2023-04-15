// external-secretsio
package externalsecretsio


// Alibaba configures this store to sync secrets using Alibaba Cloud provider.
type SecretStoreSpecProviderAlibaba struct {
	// AlibabaAuth contains a secretRef for credentials.
	Auth *SecretStoreSpecProviderAlibabaAuth `field:"required" json:"auth" yaml:"auth"`
	// Alibaba Region to be used for the provider.
	RegionId *string `field:"required" json:"regionId" yaml:"regionId"`
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}
