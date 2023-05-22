package externalsecretsio


// Auth configures how the operator authenticates with Azure.
//
// Required for ServicePrincipal auth type.
type ClusterSecretStoreSpecProviderAzurekvAuthSecretRef struct {
	// The Azure clientId of the service principle used for authentication.
	ClientId *ClusterSecretStoreSpecProviderAzurekvAuthSecretRefClientId `field:"optional" json:"clientId" yaml:"clientId"`
	// The Azure ClientSecret of the service principle used for authentication.
	ClientSecret *ClusterSecretStoreSpecProviderAzurekvAuthSecretRefClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

