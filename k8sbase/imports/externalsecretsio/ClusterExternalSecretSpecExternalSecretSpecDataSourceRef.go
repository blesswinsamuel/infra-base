// external-secretsio
package externalsecretsio


// SourceRef allows you to override the source from which the value will pulled from.
type ClusterExternalSecretSpecExternalSecretSpecDataSourceRef struct {
	// GeneratorRef points to a generator custom resource in.
	GeneratorRef *ClusterExternalSecretSpecExternalSecretSpecDataSourceRefGeneratorRef `field:"optional" json:"generatorRef" yaml:"generatorRef"`
	// SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.
	StoreRef *ClusterExternalSecretSpecExternalSecretSpecDataSourceRefStoreRef `field:"optional" json:"storeRef" yaml:"storeRef"`
}

