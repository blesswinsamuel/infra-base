package externalsecretsio


// Reference to a Secret that contains the details to authenticate with Akeyless.
type SecretStoreSpecProviderAkeylessAuthSecretRefSecretRef struct {
	// The SecretAccessID is used for authentication.
	AccessId *SecretStoreSpecProviderAkeylessAuthSecretRefSecretRefAccessId `field:"optional" json:"accessId" yaml:"accessId"`
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	AccessType *SecretStoreSpecProviderAkeylessAuthSecretRefSecretRefAccessType `field:"optional" json:"accessType" yaml:"accessType"`
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	AccessTypeParam *SecretStoreSpecProviderAkeylessAuthSecretRefSecretRefAccessTypeParam `field:"optional" json:"accessTypeParam" yaml:"accessTypeParam"`
}

