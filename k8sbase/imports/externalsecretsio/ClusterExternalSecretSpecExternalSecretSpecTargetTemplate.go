package externalsecretsio


// Template defines a blueprint for the created Secret resource.
type ClusterExternalSecretSpecExternalSecretSpecTargetTemplate struct {
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// ExternalSecretTemplateMetadata defines metadata fields for the Secret blueprint.
	Metadata *ClusterExternalSecretSpecExternalSecretSpecTargetTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	TemplateFrom *[]*ClusterExternalSecretSpecExternalSecretSpecTargetTemplateTemplateFrom `field:"optional" json:"templateFrom" yaml:"templateFrom"`
	Type *string `field:"optional" json:"type" yaml:"type"`
}

