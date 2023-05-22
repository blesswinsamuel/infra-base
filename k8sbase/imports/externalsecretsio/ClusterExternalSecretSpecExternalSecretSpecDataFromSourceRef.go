package externalsecretsio


// SourceRef points to a store or generator which contains secret values ready to use.
//
// Use this in combination with Extract or Find pull values out of a specific SecretStore. When sourceRef points to a generator Extract or Find is not supported. The generator returns a static map of values
type ClusterExternalSecretSpecExternalSecretSpecDataFromSourceRef struct {
	// GeneratorRef points to a generator custom resource in.
	GeneratorRef *ClusterExternalSecretSpecExternalSecretSpecDataFromSourceRefGeneratorRef `field:"optional" json:"generatorRef" yaml:"generatorRef"`
	// SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.
	StoreRef *ClusterExternalSecretSpecExternalSecretSpecDataFromSourceRefStoreRef `field:"optional" json:"storeRef" yaml:"storeRef"`
}

