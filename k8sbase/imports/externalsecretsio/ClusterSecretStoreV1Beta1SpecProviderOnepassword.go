// external-secretsio
package externalsecretsio


// OnePassword configures this store to sync secrets using the 1Password Cloud provider.
type ClusterSecretStoreV1Beta1SpecProviderOnepassword struct {
	// Auth defines the information necessary to authenticate against OnePassword Connect Server.
	Auth *ClusterSecretStoreV1Beta1SpecProviderOnepasswordAuth `field:"required" json:"auth" yaml:"auth"`
	// ConnectHost defines the OnePassword Connect Server to connect to.
	ConnectHost *string `field:"required" json:"connectHost" yaml:"connectHost"`
	// Vaults defines which OnePassword vaults to search in which order.
	Vaults *map[string]*float64 `field:"required" json:"vaults" yaml:"vaults"`
}

