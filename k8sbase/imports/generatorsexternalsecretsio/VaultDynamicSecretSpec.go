// generatorsexternal-secretsio
package generatorsexternalsecretsio


type VaultDynamicSecretSpec struct {
	// Vault path to obtain the dynamic secret from.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Vault provider common spec.
	Provider *VaultDynamicSecretSpecProvider `field:"required" json:"provider" yaml:"provider"`
	// Vault API method to use (GET/POST/other).
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Parameters to pass to Vault write (for non-GET methods).
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

