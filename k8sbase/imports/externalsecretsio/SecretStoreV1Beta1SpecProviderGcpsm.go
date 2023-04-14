// external-secretsio
package externalsecretsio


// GCPSM configures this store to sync secrets using Google Cloud Platform Secret Manager provider.
type SecretStoreV1Beta1SpecProviderGcpsm struct {
	// Auth defines the information necessary to authenticate against GCP.
	Auth *SecretStoreV1Beta1SpecProviderGcpsmAuth `field:"optional" json:"auth" yaml:"auth"`
	// ProjectID project where secret is located.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

