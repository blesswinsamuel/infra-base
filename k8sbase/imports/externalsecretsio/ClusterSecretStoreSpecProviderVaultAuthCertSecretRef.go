// external-secretsio
package externalsecretsio


// SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method.
type ClusterSecretStoreSpecProviderVaultAuthCertSecretRef struct {
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

