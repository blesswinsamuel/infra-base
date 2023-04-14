// generatorsexternal-secretsio
package generatorsexternalsecretsio


// Version is the Vault KV secret engine version.
//
// This can be either "v1" or "v2". Version defaults to "v2".
type VaultDynamicSecretSpecProviderVersion string

const (
	// v1.
	VaultDynamicSecretSpecProviderVersion_V1 VaultDynamicSecretSpecProviderVersion = "V1"
	// v2.
	VaultDynamicSecretSpecProviderVersion_V2 VaultDynamicSecretSpecProviderVersion = "V2"
)

