package acmecertmanagerio


// References a properly configured ACME-type Issuer which should be used to create this Challenge.
//
// If the Issuer does not exist, processing will be retried. If the Issuer is not an 'ACME' Issuer, an error will be returned and the Challenge will be marked as failed.
type ChallengeSpecIssuerRef struct {
	// Name of the resource being referred to.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Group of the resource being referred to.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Kind of the resource being referred to.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
}

