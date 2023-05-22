package externalsecretsio


// IBM configures this store to sync secrets using IBM Cloud provider.
type SecretStoreSpecProviderIbm struct {
	// Auth configures how secret-manager authenticates with the IBM secrets manager.
	Auth *SecretStoreSpecProviderIbmAuth `field:"required" json:"auth" yaml:"auth"`
	// ServiceURL is the Endpoint URL that is specific to the Secrets Manager service instance.
	ServiceUrl *string `field:"optional" json:"serviceUrl" yaml:"serviceUrl"`
}

