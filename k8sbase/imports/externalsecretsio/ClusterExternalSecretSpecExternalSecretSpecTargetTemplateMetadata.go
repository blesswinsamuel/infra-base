package externalsecretsio


// ExternalSecretTemplateMetadata defines metadata fields for the Secret blueprint.
type ClusterExternalSecretSpecExternalSecretSpecTargetTemplateMetadata struct {
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

