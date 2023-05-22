package externalsecretsio


// Senhasegura configures this store to sync secrets using senhasegura provider.
type SecretStoreV1Beta1SpecProviderSenhasegura struct {
	// Auth defines parameters to authenticate in senhasegura.
	Auth *SecretStoreV1Beta1SpecProviderSenhaseguraAuth `field:"required" json:"auth" yaml:"auth"`
	// Module defines which senhasegura module should be used to get secrets.
	Module *string `field:"required" json:"module" yaml:"module"`
	// URL of senhasegura.
	Url *string `field:"required" json:"url" yaml:"url"`
	// IgnoreSslCertificate defines if SSL certificate must be ignored.
	IgnoreSslCertificate *bool `field:"optional" json:"ignoreSslCertificate" yaml:"ignoreSslCertificate"`
}

