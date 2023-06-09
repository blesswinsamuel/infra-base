package externalsecretsio


// ExternalSecretSpec defines the desired state of ExternalSecret.
type ExternalSecretSpec struct {
	// SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.
	SecretStoreRef *ExternalSecretSpecSecretStoreRef `field:"required" json:"secretStoreRef" yaml:"secretStoreRef"`
	// ExternalSecretTarget defines the Kubernetes Secret to be created There can be only one target per ExternalSecret.
	Target *ExternalSecretSpecTarget `field:"required" json:"target" yaml:"target"`
	// Data defines the connection between the Kubernetes Secret keys and the Provider data.
	Data *[]*ExternalSecretSpecData `field:"optional" json:"data" yaml:"data"`
	// DataFrom is used to fetch all properties from a specific Provider data If multiple entries are specified, the Secret keys are merged in the specified order.
	DataFrom *[]*ExternalSecretSpecDataFrom `field:"optional" json:"dataFrom" yaml:"dataFrom"`
	// RefreshInterval is the amount of time before the values are read again from the SecretStore provider Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h" May be set to zero to fetch and create it once.
	//
	// Defaults to 1h.
	RefreshInterval *string `field:"optional" json:"refreshInterval" yaml:"refreshInterval"`
}

