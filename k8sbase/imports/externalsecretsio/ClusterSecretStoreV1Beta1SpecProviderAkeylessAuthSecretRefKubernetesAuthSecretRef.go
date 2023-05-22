package externalsecretsio


// Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Akeyless.
//
// If a name is specified without a key, `token` is the default. If one is not specified, the one bound to the controller will be used.
type ClusterSecretStoreV1Beta1SpecProviderAkeylessAuthSecretRefKubernetesAuthSecretRef struct {
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

