// ingressroute_traefikcontainous
package ingressroute_traefikcontainous


// Domain holds a domain name with SANs.
type IngressRouteSpecTlsDomains struct {
	// Main defines the main domain name.
	Main *string `field:"optional" json:"main" yaml:"main"`
	// SANs defines the subject alternative domain names.
	Sans *[]*string `field:"optional" json:"sans" yaml:"sans"`
}

