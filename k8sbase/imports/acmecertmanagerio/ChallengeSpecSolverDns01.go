// acmecert-managerio
package acmecertmanagerio


// Configures cert-manager to attempt to complete authorizations by performing the DNS01 challenge flow.
type ChallengeSpecSolverDns01 struct {
	// Use the 'ACME DNS' (https://github.com/joohoi/acme-dns) API to manage DNS01 challenge records.
	AcmeDns *ChallengeSpecSolverDns01AcmeDns `field:"optional" json:"acmeDns" yaml:"acmeDns"`
	// Use the Akamai DNS zone management API to manage DNS01 challenge records.
	Akamai *ChallengeSpecSolverDns01Akamai `field:"optional" json:"akamai" yaml:"akamai"`
	// Use the Microsoft Azure DNS API to manage DNS01 challenge records.
	AzureDns *ChallengeSpecSolverDns01AzureDns `field:"optional" json:"azureDns" yaml:"azureDns"`
	// Use the Google Cloud DNS API to manage DNS01 challenge records.
	CloudDns *ChallengeSpecSolverDns01CloudDns `field:"optional" json:"cloudDns" yaml:"cloudDns"`
	// Use the Cloudflare API to manage DNS01 challenge records.
	Cloudflare *ChallengeSpecSolverDns01Cloudflare `field:"optional" json:"cloudflare" yaml:"cloudflare"`
	// CNAMEStrategy configures how the DNS01 provider should handle CNAME records when found in DNS zones.
	CnameStrategy ChallengeSpecSolverDns01CnameStrategy `field:"optional" json:"cnameStrategy" yaml:"cnameStrategy"`
	// Use the DigitalOcean DNS API to manage DNS01 challenge records.
	Digitalocean *ChallengeSpecSolverDns01Digitalocean `field:"optional" json:"digitalocean" yaml:"digitalocean"`
	// Use RFC2136 ("Dynamic Updates in the Domain Name System") (https://datatracker.ietf.org/doc/rfc2136/) to manage DNS01 challenge records.
	Rfc2136 *ChallengeSpecSolverDns01Rfc2136 `field:"optional" json:"rfc2136" yaml:"rfc2136"`
	// Use the AWS Route53 API to manage DNS01 challenge records.
	Route53 *ChallengeSpecSolverDns01Route53 `field:"optional" json:"route53" yaml:"route53"`
	// Configure an external webhook based DNS01 challenge solver to manage DNS01 challenge records.
	Webhook *ChallengeSpecSolverDns01Webhook `field:"optional" json:"webhook" yaml:"webhook"`
}

