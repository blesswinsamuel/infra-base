package certmanagerio


// CNAMEStrategy configures how the DNS01 provider should handle CNAME records when found in DNS zones.
type ClusterIssuerSpecAcmeSolversDns01CnameStrategy string

const (
	// None.
	ClusterIssuerSpecAcmeSolversDns01CnameStrategy_NONE ClusterIssuerSpecAcmeSolversDns01CnameStrategy = "NONE"
	// Follow.
	ClusterIssuerSpecAcmeSolversDns01CnameStrategy_FOLLOW ClusterIssuerSpecAcmeSolversDns01CnameStrategy = "FOLLOW"
)

