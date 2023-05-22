package externalsecretsio


// KeeperSecurity configures this store to sync secrets using the KeeperSecurity provider.
type SecretStoreV1Beta1SpecProviderKeepersecurity struct {
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	AuthRef *SecretStoreV1Beta1SpecProviderKeepersecurityAuthRef `field:"required" json:"authRef" yaml:"authRef"`
	FolderId *string `field:"required" json:"folderId" yaml:"folderId"`
}

