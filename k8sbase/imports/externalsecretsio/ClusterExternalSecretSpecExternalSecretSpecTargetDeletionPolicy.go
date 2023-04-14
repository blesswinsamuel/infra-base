// external-secretsio
package externalsecretsio


// DeletionPolicy defines rules on how to delete the resulting Secret Defaults to 'Retain'.
type ClusterExternalSecretSpecExternalSecretSpecTargetDeletionPolicy string

const (
	// Delete.
	ClusterExternalSecretSpecExternalSecretSpecTargetDeletionPolicy_DELETE ClusterExternalSecretSpecExternalSecretSpecTargetDeletionPolicy = "DELETE"
	// Merge.
	ClusterExternalSecretSpecExternalSecretSpecTargetDeletionPolicy_MERGE ClusterExternalSecretSpecExternalSecretSpecTargetDeletionPolicy = "MERGE"
	// Retain.
	ClusterExternalSecretSpecExternalSecretSpecTargetDeletionPolicy_RETAIN ClusterExternalSecretSpecExternalSecretSpecTargetDeletionPolicy = "RETAIN"
)

