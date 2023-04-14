// external-secretsio
package externalsecretsio


// YandexLockbox configures this store to sync secrets using Yandex Lockbox provider.
type SecretStoreSpecProviderYandexlockbox struct {
	// Auth defines the information necessary to authenticate against Yandex Lockbox.
	Auth *SecretStoreSpecProviderYandexlockboxAuth `field:"required" json:"auth" yaml:"auth"`
	// Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443').
	ApiEndpoint *string `field:"optional" json:"apiEndpoint" yaml:"apiEndpoint"`
	// The provider for the CA bundle to use to validate Yandex.Cloud server certificate.
	CaProvider *SecretStoreSpecProviderYandexlockboxCaProvider `field:"optional" json:"caProvider" yaml:"caProvider"`
}

