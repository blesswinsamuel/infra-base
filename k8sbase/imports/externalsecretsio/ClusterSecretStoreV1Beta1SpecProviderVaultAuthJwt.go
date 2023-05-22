package externalsecretsio


// Jwt authenticates with Vault by passing role and JWT token using the JWT/OIDC authentication method.
type ClusterSecretStoreV1Beta1SpecProviderVaultAuthJwt struct {
	// Path where the JWT authentication backend is mounted in Vault, e.g: "jwt".
	Path *string `field:"required" json:"path" yaml:"path"`
	// Optional ServiceAccountToken specifies the Kubernetes service account for which to request a token for with the `TokenRequest` API.
	KubernetesServiceAccountToken *ClusterSecretStoreV1Beta1SpecProviderVaultAuthJwtKubernetesServiceAccountToken `field:"optional" json:"kubernetesServiceAccountToken" yaml:"kubernetesServiceAccountToken"`
	// Role is a JWT role to authenticate using the JWT/OIDC Vault authentication method.
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Vault using the JWT/OIDC authentication method.
	SecretRef *ClusterSecretStoreV1Beta1SpecProviderVaultAuthJwtSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

