// acmecert-managerio
package acmecertmanagerio


// The type of ACME challenge this resource represents.
//
// One of "HTTP-01" or "DNS-01".
type ChallengeSpecType string

const (
	// HTTP-01.
	ChallengeSpecType_HTTP_01 ChallengeSpecType = "HTTP_01"
	// DNS-01.
	ChallengeSpecType_DNS_01 ChallengeSpecType = "DNS_01"
)

