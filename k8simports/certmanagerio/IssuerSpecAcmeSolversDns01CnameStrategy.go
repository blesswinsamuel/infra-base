package certmanagerio


// CNAMEStrategy configures how the DNS01 provider should handle CNAME records when found in DNS zones.
type IssuerSpecAcmeSolversDns01CnameStrategy string

const (
	// None.
	IssuerSpecAcmeSolversDns01CnameStrategy_NONE IssuerSpecAcmeSolversDns01CnameStrategy = "NONE"
	// Follow.
	IssuerSpecAcmeSolversDns01CnameStrategy_FOLLOW IssuerSpecAcmeSolversDns01CnameStrategy = "FOLLOW"
)

