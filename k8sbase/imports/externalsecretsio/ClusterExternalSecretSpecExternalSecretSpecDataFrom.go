// external-secretsio
package externalsecretsio


type ClusterExternalSecretSpecExternalSecretSpecDataFrom struct {
	// Used to extract multiple key/value pairs from one secret Note: Extract does not support sourceRef.Generator or sourceRef.GeneratorRef.
	Extract *ClusterExternalSecretSpecExternalSecretSpecDataFromExtract `field:"optional" json:"extract" yaml:"extract"`
	// Used to find secrets based on tags or regular expressions Note: Find does not support sourceRef.Generator or sourceRef.GeneratorRef.
	Find *ClusterExternalSecretSpecExternalSecretSpecDataFromFind `field:"optional" json:"find" yaml:"find"`
	// Used to rewrite secret Keys after getting them from the secret Provider Multiple Rewrite operations can be provided.
	//
	// They are applied in a layered order (first to last).
	Rewrite *[]*ClusterExternalSecretSpecExternalSecretSpecDataFromRewrite `field:"optional" json:"rewrite" yaml:"rewrite"`
	// SourceRef points to a store or generator which contains secret values ready to use.
	//
	// Use this in combination with Extract or Find pull values out of a specific SecretStore. When sourceRef points to a generator Extract or Find is not supported. The generator returns a static map of values
	SourceRef *ClusterExternalSecretSpecExternalSecretSpecDataFromSourceRef `field:"optional" json:"sourceRef" yaml:"sourceRef"`
}

