package acmecertmanagerio


// Use the Akamai DNS zone management API to manage DNS01 challenge records.
type ChallengeSpecSolverDns01Akamai struct {
	// A reference to a specific 'key' within a Secret resource.
	//
	// In some instances, `key` is a required field.
	AccessTokenSecretRef *ChallengeSpecSolverDns01AkamaiAccessTokenSecretRef `field:"required" json:"accessTokenSecretRef" yaml:"accessTokenSecretRef"`
	// A reference to a specific 'key' within a Secret resource.
	//
	// In some instances, `key` is a required field.
	ClientSecretSecretRef *ChallengeSpecSolverDns01AkamaiClientSecretSecretRef `field:"required" json:"clientSecretSecretRef" yaml:"clientSecretSecretRef"`
	// A reference to a specific 'key' within a Secret resource.
	//
	// In some instances, `key` is a required field.
	ClientTokenSecretRef *ChallengeSpecSolverDns01AkamaiClientTokenSecretRef `field:"required" json:"clientTokenSecretRef" yaml:"clientTokenSecretRef"`
	ServiceConsumerDomain *string `field:"required" json:"serviceConsumerDomain" yaml:"serviceConsumerDomain"`
}
