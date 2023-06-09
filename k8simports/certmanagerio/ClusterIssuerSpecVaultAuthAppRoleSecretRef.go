package certmanagerio


// Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault.
//
// The `key` field must be specified and denotes which entry within the Secret resource is used as the app role secret.
type ClusterIssuerSpecVaultAuthAppRoleSecretRef struct {
	// Name of the resource being referred to.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `field:"required" json:"name" yaml:"name"`
	// The key of the entry in the Secret resource's `data` field to be used.
	//
	// Some instances of this field may be defaulted, in others it may be required.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

