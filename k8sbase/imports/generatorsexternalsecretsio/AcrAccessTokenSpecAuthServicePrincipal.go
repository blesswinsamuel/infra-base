// generatorsexternal-secretsio
package generatorsexternalsecretsio


// ServicePrincipal uses Azure Service Principal credentials to authenticate with Azure.
type AcrAccessTokenSpecAuthServicePrincipal struct {
	// Configuration used to authenticate with Azure using static credentials stored in a Kind=Secret.
	SecretRef *AcrAccessTokenSpecAuthServicePrincipalSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

