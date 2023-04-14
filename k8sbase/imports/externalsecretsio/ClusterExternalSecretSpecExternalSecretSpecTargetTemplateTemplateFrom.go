// external-secretsio
package externalsecretsio


type ClusterExternalSecretSpecExternalSecretSpecTargetTemplateTemplateFrom struct {
	ConfigMap *ClusterExternalSecretSpecExternalSecretSpecTargetTemplateTemplateFromConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	Literal *string `field:"optional" json:"literal" yaml:"literal"`
	Secret *ClusterExternalSecretSpecExternalSecretSpecTargetTemplateTemplateFromSecret `field:"optional" json:"secret" yaml:"secret"`
	Target *string `field:"optional" json:"target" yaml:"target"`
}

