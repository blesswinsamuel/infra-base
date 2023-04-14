// acmecert-managerio
package acmecertmanagerio


// CNAMEStrategy configures how the DNS01 provider should handle CNAME records when found in DNS zones.
type ChallengeSpecSolverDns01CnameStrategy string

const (
	// None.
	ChallengeSpecSolverDns01CnameStrategy_NONE ChallengeSpecSolverDns01CnameStrategy = "NONE"
	// Follow.
	ChallengeSpecSolverDns01CnameStrategy_FOLLOW ChallengeSpecSolverDns01CnameStrategy = "FOLLOW"
)

