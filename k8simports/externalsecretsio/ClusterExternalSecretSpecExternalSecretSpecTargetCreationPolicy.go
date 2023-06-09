package externalsecretsio


// CreationPolicy defines rules on how to create the resulting Secret Defaults to 'Owner'.
type ClusterExternalSecretSpecExternalSecretSpecTargetCreationPolicy string

const (
	// Owner.
	ClusterExternalSecretSpecExternalSecretSpecTargetCreationPolicy_OWNER ClusterExternalSecretSpecExternalSecretSpecTargetCreationPolicy = "OWNER"
	// Orphan.
	ClusterExternalSecretSpecExternalSecretSpecTargetCreationPolicy_ORPHAN ClusterExternalSecretSpecExternalSecretSpecTargetCreationPolicy = "ORPHAN"
	// Merge.
	ClusterExternalSecretSpecExternalSecretSpecTargetCreationPolicy_MERGE ClusterExternalSecretSpecExternalSecretSpecTargetCreationPolicy = "MERGE"
	// None.
	ClusterExternalSecretSpecExternalSecretSpecTargetCreationPolicy_NONE ClusterExternalSecretSpecExternalSecretSpecTargetCreationPolicy = "NONE"
)

