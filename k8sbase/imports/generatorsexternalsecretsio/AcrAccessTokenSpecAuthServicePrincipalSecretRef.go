// generatorsexternal-secretsio
package generatorsexternalsecretsio


// Configuration used to authenticate with Azure using static credentials stored in a Kind=Secret.
type AcrAccessTokenSpecAuthServicePrincipalSecretRef struct {
	// The Azure clientId of the service principle used for authentication.
	ClientId *AcrAccessTokenSpecAuthServicePrincipalSecretRefClientId `field:"optional" json:"clientId" yaml:"clientId"`
	// The Azure ClientSecret of the service principle used for authentication.
	ClientSecret *AcrAccessTokenSpecAuthServicePrincipalSecretRefClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

