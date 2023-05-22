package middlewarestcp_traefikio


// MiddlewareTCPSpec defines the desired state of a MiddlewareTCP.
type MiddlewareTcpSpec struct {
	// InFlightConn defines the InFlightConn middleware configuration.
	InFlightConn *MiddlewareTcpSpecInFlightConn `field:"optional" json:"inFlightConn" yaml:"inFlightConn"`
	// IPWhiteList defines the IPWhiteList middleware configuration.
	IpWhiteList *MiddlewareTcpSpecIpWhiteList `field:"optional" json:"ipWhiteList" yaml:"ipWhiteList"`
}

