package externalsecretsio


// Reference to a Secret that contains the details to authenticate with Akeyless.
type ClusterSecretStoreSpecProviderAkeylessAuthSecretRefSecretRef struct {
	// The SecretAccessID is used for authentication.
	AccessId *ClusterSecretStoreSpecProviderAkeylessAuthSecretRefSecretRefAccessId `field:"optional" json:"accessId" yaml:"accessId"`
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	AccessType *ClusterSecretStoreSpecProviderAkeylessAuthSecretRefSecretRefAccessType `field:"optional" json:"accessType" yaml:"accessType"`
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	AccessTypeParam *ClusterSecretStoreSpecProviderAkeylessAuthSecretRefSecretRefAccessTypeParam `field:"optional" json:"accessTypeParam" yaml:"accessTypeParam"`
}

