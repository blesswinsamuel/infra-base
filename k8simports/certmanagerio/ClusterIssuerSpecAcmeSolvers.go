package certmanagerio


// An ACMEChallengeSolver describes how to solve ACME challenges for the issuer it is part of.
//
// A selector may be provided to use different solving strategies for different DNS names. Only one of HTTP01 or DNS01 must be provided.
type ClusterIssuerSpecAcmeSolvers struct {
	// Configures cert-manager to attempt to complete authorizations by performing the DNS01 challenge flow.
	Dns01 *ClusterIssuerSpecAcmeSolversDns01 `field:"optional" json:"dns01" yaml:"dns01"`
	// Configures cert-manager to attempt to complete authorizations by performing the HTTP01 challenge flow.
	//
	// It is not possible to obtain certificates for wildcard domain names (e.g. `*.example.com`) using the HTTP01 challenge mechanism.
	Http01 *ClusterIssuerSpecAcmeSolversHttp01 `field:"optional" json:"http01" yaml:"http01"`
	// Selector selects a set of DNSNames on the Certificate resource that should be solved using this challenge solver.
	//
	// If not specified, the solver will be treated as the 'default' solver with the lowest priority, i.e. if any other solver has a more specific match, it will be used instead.
	Selector *ClusterIssuerSpecAcmeSolversSelector `field:"optional" json:"selector" yaml:"selector"`
}

