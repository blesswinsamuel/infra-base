package middlewarestcp_traefikio


// IPWhiteList defines the IPWhiteList middleware configuration.
type MiddlewareTcpSpecIpWhiteList struct {
	// SourceRange defines the allowed IPs (or ranges of allowed IPs by using CIDR notation).
	SourceRange *[]*string `field:"optional" json:"sourceRange" yaml:"sourceRange"`
}

