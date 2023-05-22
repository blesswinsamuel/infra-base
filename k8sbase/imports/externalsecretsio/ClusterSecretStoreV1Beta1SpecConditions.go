package externalsecretsio


// ClusterSecretStoreCondition describes a condition by which to choose namespaces to process ExternalSecrets in for a ClusterSecretStore instance.
type ClusterSecretStoreV1Beta1SpecConditions struct {
	// Choose namespaces by name.
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// Choose namespace using a labelSelector.
	NamespaceSelector *ClusterSecretStoreV1Beta1SpecConditionsNamespaceSelector `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
}

