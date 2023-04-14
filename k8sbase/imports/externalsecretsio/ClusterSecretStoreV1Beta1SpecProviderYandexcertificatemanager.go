// external-secretsio
package externalsecretsio


// YandexCertificateManager configures this store to sync secrets using Yandex Certificate Manager provider.
type ClusterSecretStoreV1Beta1SpecProviderYandexcertificatemanager struct {
	// Auth defines the information necessary to authenticate against Yandex Certificate Manager.
	Auth *ClusterSecretStoreV1Beta1SpecProviderYandexcertificatemanagerAuth `field:"required" json:"auth" yaml:"auth"`
	// Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443').
	ApiEndpoint *string `field:"optional" json:"apiEndpoint" yaml:"apiEndpoint"`
	// The provider for the CA bundle to use to validate Yandex.Cloud server certificate.
	CaProvider *ClusterSecretStoreV1Beta1SpecProviderYandexcertificatemanagerCaProvider `field:"optional" json:"caProvider" yaml:"caProvider"`
}

