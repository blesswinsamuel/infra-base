// external-secretsio
package externalsecretsio


// AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.
type SecretStoreSpecProviderVaultAuthAppRole struct {
	// Path where the App Role authentication backend is mounted in Vault, e.g: "approle".
	Path *string `field:"required" json:"path" yaml:"path"`
	// RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault.
	//
	// The `key` field must be specified and denotes which entry within the Secret resource is used as the app role secret.
	SecretRef *SecretStoreSpecProviderVaultAuthAppRoleSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

