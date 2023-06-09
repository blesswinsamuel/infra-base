package certmanagerio


// Configure an external webhook based DNS01 challenge solver to manage DNS01 challenge records.
type IssuerSpecAcmeSolversDns01Webhook struct {
	// The API group name that should be used when POSTing ChallengePayload resources to the webhook apiserver.
	//
	// This should be the same as the GroupName specified in the webhook provider implementation.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The name of the solver to use, as defined in the webhook provider implementation.
	//
	// This will typically be the name of the provider, e.g. 'cloudflare'.
	SolverName *string `field:"required" json:"solverName" yaml:"solverName"`
	// Additional configuration that should be passed to the webhook apiserver when challenges are processed.
	//
	// This can contain arbitrary JSON data. Secret values should not be specified in this stanza. If secret values are needed (e.g. credentials for a DNS service), you should use a SecretKeySelector to reference a Secret resource. For details on the schema of this field, consult the webhook provider implementation's documentation.
	Config interface{} `field:"optional" json:"config" yaml:"config"`
}

