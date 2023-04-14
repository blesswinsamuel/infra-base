// external-secretsio
package externalsecretsio


// Auth configures how the operator authenticates with Azure.
//
// Required for ServicePrincipal auth type.
type SecretStoreSpecProviderAzurekvAuthSecretRef struct {
	// The Azure clientId of the service principle used for authentication.
	ClientId *SecretStoreSpecProviderAzurekvAuthSecretRefClientId `field:"optional" json:"clientId" yaml:"clientId"`
	// The Azure ClientSecret of the service principle used for authentication.
	ClientSecret *SecretStoreSpecProviderAzurekvAuthSecretRefClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

