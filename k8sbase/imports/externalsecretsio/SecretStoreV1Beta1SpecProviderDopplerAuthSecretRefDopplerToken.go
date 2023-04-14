// external-secretsio
package externalsecretsio


// The DopplerToken is used for authentication.
//
// See https://docs.doppler.com/reference/api#authentication for auth token types. The Key attribute defaults to dopplerToken if not specified.
type SecretStoreV1Beta1SpecProviderDopplerAuthSecretRefDopplerToken struct {
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

