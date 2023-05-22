package externalsecretsio


// Oracle configures this store to sync secrets using Oracle Vault provider.
type SecretStoreV1Beta1SpecProviderOracle struct {
	// Region is the region where vault is located.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Vault is the vault's OCID of the specific vault where secret is located.
	Vault *string `field:"required" json:"vault" yaml:"vault"`
	// Auth configures how secret-manager authenticates with the Oracle Vault.
	//
	// If empty, use the instance principal, otherwise the user credentials specified in Auth.
	Auth *SecretStoreV1Beta1SpecProviderOracleAuth `field:"optional" json:"auth" yaml:"auth"`
}

