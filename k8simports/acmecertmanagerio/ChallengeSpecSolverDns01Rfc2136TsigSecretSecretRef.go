package acmecertmanagerio


// The name of the secret containing the TSIG value.
//
// If ``tsigKeyName`` is defined, this field is required.
type ChallengeSpecSolverDns01Rfc2136TsigSecretSecretRef struct {
	// Name of the resource being referred to.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `field:"required" json:"name" yaml:"name"`
	// The key of the entry in the Secret resource's `data` field to be used.
	//
	// Some instances of this field may be defaulted, in others it may be required.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

