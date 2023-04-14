// acmecert-managerio
package acmecertmanagerio


type ChallengeSpec struct {
	// The URL to the ACME Authorization resource that this challenge is a part of.
	AuthorizationUrl *string `field:"required" json:"authorizationUrl" yaml:"authorizationUrl"`
	// dnsName is the identifier that this challenge is for, e.g. example.com. If the requested DNSName is a 'wildcard', this field MUST be set to the non-wildcard domain, e.g. for `*.example.com`, it must be `example.com`.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// References a properly configured ACME-type Issuer which should be used to create this Challenge.
	//
	// If the Issuer does not exist, processing will be retried. If the Issuer is not an 'ACME' Issuer, an error will be returned and the Challenge will be marked as failed.
	IssuerRef *ChallengeSpecIssuerRef `field:"required" json:"issuerRef" yaml:"issuerRef"`
	// The ACME challenge key for this challenge For HTTP01 challenges, this is the value that must be responded with to complete the HTTP01 challenge in the format: `<private key JWK thumbprint>.<key from acme server for challenge>`. For DNS01 challenges, this is the base64 encoded SHA256 sum of the `<private key JWK thumbprint>.<key from acme server for challenge>` text that must be set as the TXT record content.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Contains the domain solving configuration that should be used to solve this challenge resource.
	Solver *ChallengeSpecSolver `field:"required" json:"solver" yaml:"solver"`
	// The ACME challenge token for this challenge.
	//
	// This is the raw value returned from the ACME server.
	Token *string `field:"required" json:"token" yaml:"token"`
	// The type of ACME challenge this resource represents.
	//
	// One of "HTTP-01" or "DNS-01".
	Type ChallengeSpecType `field:"required" json:"type" yaml:"type"`
	// The URL of the ACME Challenge resource for this challenge.
	//
	// This can be used to lookup details about the status of this challenge.
	Url *string `field:"required" json:"url" yaml:"url"`
	// wildcard will be true if this challenge is for a wildcard identifier, for example '*.example.com'.
	Wildcard *bool `field:"optional" json:"wildcard" yaml:"wildcard"`
}

