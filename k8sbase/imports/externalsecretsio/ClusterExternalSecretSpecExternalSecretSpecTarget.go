// external-secretsio
package externalsecretsio


// ExternalSecretTarget defines the Kubernetes Secret to be created There can be only one target per ExternalSecret.
type ClusterExternalSecretSpecExternalSecretSpecTarget struct {
	// CreationPolicy defines rules on how to create the resulting Secret Defaults to 'Owner'.
	CreationPolicy ClusterExternalSecretSpecExternalSecretSpecTargetCreationPolicy `field:"optional" json:"creationPolicy" yaml:"creationPolicy"`
	// DeletionPolicy defines rules on how to delete the resulting Secret Defaults to 'Retain'.
	DeletionPolicy ClusterExternalSecretSpecExternalSecretSpecTargetDeletionPolicy `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Immutable defines if the final secret will be immutable.
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
	// Name defines the name of the Secret resource to be managed This field is immutable Defaults to the .metadata.name of the ExternalSecret resource.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Template defines a blueprint for the created Secret resource.
	Template *ClusterExternalSecretSpecExternalSecretSpecTargetTemplate `field:"optional" json:"template" yaml:"template"`
}

