// external-secretsio
package externalsecretsio


type PushSecretSpecSecretStoreRefs struct {
	// Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to `SecretStore`.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Optionally, sync to secret stores with label selector.
	LabelSelector *PushSecretSpecSecretStoreRefsLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// Optionally, sync to the SecretStore of the given name.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

