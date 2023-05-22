package externalsecretsio


// SecretStoreSpec defines the desired state of SecretStore.
type ClusterSecretStoreSpec struct {
	// Used to configure the provider.
	//
	// Only one provider may be set.
	Provider *ClusterSecretStoreSpecProvider `field:"required" json:"provider" yaml:"provider"`
	// Used to select the correct KES controller (think: ingress.ingressClassName) The KES controller is instantiated with a specific controller name and filters ES based on this property.
	Controller *string `field:"optional" json:"controller" yaml:"controller"`
	// Used to configure http retries if failed.
	RetrySettings *ClusterSecretStoreSpecRetrySettings `field:"optional" json:"retrySettings" yaml:"retrySettings"`
}

