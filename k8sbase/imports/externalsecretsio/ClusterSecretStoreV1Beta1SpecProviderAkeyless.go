package externalsecretsio


// Akeyless configures this store to sync secrets using Akeyless Vault provider.
type ClusterSecretStoreV1Beta1SpecProviderAkeyless struct {
	// Akeyless GW API Url from which the secrets to be fetched from.
	AkeylessGwApiUrl *string `field:"required" json:"akeylessGwApiUrl" yaml:"akeylessGwApiUrl"`
	// Auth configures how the operator authenticates with Akeyless.
	AuthSecretRef *ClusterSecretStoreV1Beta1SpecProviderAkeylessAuthSecretRef `field:"required" json:"authSecretRef" yaml:"authSecretRef"`
	// PEM/base64 encoded CA bundle used to validate Akeyless Gateway certificate.
	//
	// Only used if the AkeylessGWApiURL URL is using HTTPS protocol. If not set the system root certificates are used to validate the TLS connection.
	CaBundle *string `field:"optional" json:"caBundle" yaml:"caBundle"`
	// The provider for the CA bundle to use to validate Akeyless Gateway certificate.
	CaProvider *ClusterSecretStoreV1Beta1SpecProviderAkeylessCaProvider `field:"optional" json:"caProvider" yaml:"caProvider"`
}

