package externalsecretsio


// Reference to a Secret that contains the details to authenticate with Akeyless.
type ClusterSecretStoreV1Beta1SpecProviderAkeylessAuthSecretRefSecretRef struct {
	// The SecretAccessID is used for authentication.
	AccessId *ClusterSecretStoreV1Beta1SpecProviderAkeylessAuthSecretRefSecretRefAccessId `field:"optional" json:"accessId" yaml:"accessId"`
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	AccessType *ClusterSecretStoreV1Beta1SpecProviderAkeylessAuthSecretRefSecretRefAccessType `field:"optional" json:"accessType" yaml:"accessType"`
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	AccessTypeParam *ClusterSecretStoreV1Beta1SpecProviderAkeylessAuthSecretRefSecretRefAccessTypeParam `field:"optional" json:"accessTypeParam" yaml:"accessTypeParam"`
}

