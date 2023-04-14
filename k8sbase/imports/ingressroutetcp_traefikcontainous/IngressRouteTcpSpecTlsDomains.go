// ingressroutetcp_traefikcontainous
package ingressroutetcp_traefikcontainous


// Domain holds a domain name with SANs.
type IngressRouteTcpSpecTlsDomains struct {
	// Main defines the main domain name.
	Main *string `field:"optional" json:"main" yaml:"main"`
	// SANs defines the subject alternative domain names.
	Sans *[]*string `field:"optional" json:"sans" yaml:"sans"`
}

