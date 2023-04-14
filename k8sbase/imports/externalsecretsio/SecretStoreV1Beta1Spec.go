// external-secretsio
package externalsecretsio


// SecretStoreSpec defines the desired state of SecretStore.
type SecretStoreV1Beta1Spec struct {
	// Used to configure the provider.
	//
	// Only one provider may be set.
	Provider *SecretStoreV1Beta1SpecProvider `field:"required" json:"provider" yaml:"provider"`
	// Used to constraint a ClusterSecretStore to specific namespaces.
	//
	// Relevant only to ClusterSecretStore.
	Conditions *[]*SecretStoreV1Beta1SpecConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// Used to select the correct KES controller (think: ingress.ingressClassName) The KES controller is instantiated with a specific controller name and filters ES based on this property.
	Controller *string `field:"optional" json:"controller" yaml:"controller"`
	// Used to configure store refresh interval in seconds.
	//
	// Empty or 0 will default to the controller config.
	RefreshInterval *float64 `field:"optional" json:"refreshInterval" yaml:"refreshInterval"`
	// Used to configure http retries if failed.
	RetrySettings *SecretStoreV1Beta1SpecRetrySettings `field:"optional" json:"retrySettings" yaml:"retrySettings"`
}

