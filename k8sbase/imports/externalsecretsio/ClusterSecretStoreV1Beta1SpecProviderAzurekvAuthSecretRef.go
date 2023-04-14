// external-secretsio
package externalsecretsio


// Auth configures how the operator authenticates with Azure.
//
// Required for ServicePrincipal auth type.
type ClusterSecretStoreV1Beta1SpecProviderAzurekvAuthSecretRef struct {
	// The Azure clientId of the service principle used for authentication.
	ClientId *ClusterSecretStoreV1Beta1SpecProviderAzurekvAuthSecretRefClientId `field:"optional" json:"clientId" yaml:"clientId"`
	// The Azure ClientSecret of the service principle used for authentication.
	ClientSecret *ClusterSecretStoreV1Beta1SpecProviderAzurekvAuthSecretRefClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

