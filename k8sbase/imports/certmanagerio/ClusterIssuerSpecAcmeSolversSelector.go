// cert-managerio
package certmanagerio


// Selector selects a set of DNSNames on the Certificate resource that should be solved using this challenge solver.
//
// If not specified, the solver will be treated as the 'default' solver with the lowest priority, i.e. if any other solver has a more specific match, it will be used instead.
type ClusterIssuerSpecAcmeSolversSelector struct {
	// List of DNSNames that this solver will be used to solve.
	//
	// If specified and a match is found, a dnsNames selector will take precedence over a dnsZones selector. If multiple solvers match with the same dnsNames value, the solver with the most matching labels in matchLabels will be selected. If neither has more matches, the solver defined earlier in the list will be selected.
	DnsNames *[]*string `field:"optional" json:"dnsNames" yaml:"dnsNames"`
	// List of DNSZones that this solver will be used to solve.
	//
	// The most specific DNS zone match specified here will take precedence over other DNS zone matches, so a solver specifying sys.example.com will be selected over one specifying example.com for the domain www.sys.example.com. If multiple solvers match with the same dnsZones value, the solver with the most matching labels in matchLabels will be selected. If neither has more matches, the solver defined earlier in the list will be selected.
	DnsZones *[]*string `field:"optional" json:"dnsZones" yaml:"dnsZones"`
	// A label selector that is used to refine the set of certificate's that this challenge solver will apply to.
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

