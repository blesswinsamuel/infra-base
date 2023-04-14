// external-secretsio
package externalsecretsio


// PushSecretSpec configures the behavior of the PushSecret.
type PushSecretSpec struct {
	SecretStoreRefs *[]*PushSecretSpecSecretStoreRefs `field:"required" json:"secretStoreRefs" yaml:"secretStoreRefs"`
	// The Secret Selector (k8s source) for the Push Secret.
	Selector *PushSecretSpecSelector `field:"required" json:"selector" yaml:"selector"`
	// Secret Data that should be pushed to providers.
	Data *[]*PushSecretSpecData `field:"optional" json:"data" yaml:"data"`
	// Deletion Policy to handle Secrets in the provider.
	//
	// Possible Values: "Delete/None". Defaults to "None".
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The Interval to which External Secrets will try to push a secret definition.
	RefreshInterval *string `field:"optional" json:"refreshInterval" yaml:"refreshInterval"`
}

