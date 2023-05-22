package externalsecretsio


// Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault.
//
// The `key` field must be specified and denotes which entry within the Secret resource is used as the app role secret.
type SecretStoreSpecProviderVaultAuthAppRoleSecretRef struct {
	// The key of the entry in the Secret resource's `data` field to be used.
	//
	// Some instances of this field may be defaulted, in others it may be required.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The name of the Secret resource being referred to.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Namespace of the resource being referred to.
	//
	// Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

