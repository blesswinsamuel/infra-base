package externalsecretsio


// The spec for the ExternalSecrets to be created.
type ClusterExternalSecretSpecExternalSecretSpec struct {
	// Data defines the connection between the Kubernetes Secret keys and the Provider data.
	Data *[]*ClusterExternalSecretSpecExternalSecretSpecData `field:"optional" json:"data" yaml:"data"`
	// DataFrom is used to fetch all properties from a specific Provider data If multiple entries are specified, the Secret keys are merged in the specified order.
	DataFrom *[]*ClusterExternalSecretSpecExternalSecretSpecDataFrom `field:"optional" json:"dataFrom" yaml:"dataFrom"`
	// RefreshInterval is the amount of time before the values are read again from the SecretStore provider Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h" May be set to zero to fetch and create it once.
	//
	// Defaults to 1h.
	RefreshInterval *string `field:"optional" json:"refreshInterval" yaml:"refreshInterval"`
	// SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.
	SecretStoreRef *ClusterExternalSecretSpecExternalSecretSpecSecretStoreRef `field:"optional" json:"secretStoreRef" yaml:"secretStoreRef"`
	// ExternalSecretTarget defines the Kubernetes Secret to be created There can be only one target per ExternalSecret.
	Target *ClusterExternalSecretSpecExternalSecretSpecTarget `field:"optional" json:"target" yaml:"target"`
}

