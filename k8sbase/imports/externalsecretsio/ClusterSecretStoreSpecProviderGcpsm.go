// external-secretsio
package externalsecretsio


// GCPSM configures this store to sync secrets using Google Cloud Platform Secret Manager provider.
type ClusterSecretStoreSpecProviderGcpsm struct {
	// Auth defines the information necessary to authenticate against GCP.
	Auth *ClusterSecretStoreSpecProviderGcpsmAuth `field:"optional" json:"auth" yaml:"auth"`
	// ProjectID project where secret is located.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

