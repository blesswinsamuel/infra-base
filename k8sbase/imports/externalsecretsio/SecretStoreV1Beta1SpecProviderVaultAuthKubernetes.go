// external-secretsio
package externalsecretsio


// Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.
type SecretStoreV1Beta1SpecProviderVaultAuthKubernetes struct {
	// Path where the Kubernetes authentication backend is mounted in Vault, e.g: "kubernetes".
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// A required field containing the Vault Role to assume.
	//
	// A Role binds a Kubernetes ServiceAccount with a set of Vault policies.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault.
	//
	// If a name is specified without a key, `token` is the default. If one is not specified, the one bound to the controller will be used.
	SecretRef *SecretStoreV1Beta1SpecProviderVaultAuthKubernetesSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	// Optional service account field containing the name of a kubernetes ServiceAccount.
	//
	// If the service account is specified, the service account secret token JWT will be used for authenticating with Vault. If the service account selector is not supplied, the secretRef will be used instead.
	ServiceAccountRef *SecretStoreV1Beta1SpecProviderVaultAuthKubernetesServiceAccountRef `field:"optional" json:"serviceAccountRef" yaml:"serviceAccountRef"`
}

