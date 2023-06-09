package middlewarestcp_traefikio


// InFlightConn defines the InFlightConn middleware configuration.
type MiddlewareTcpSpecInFlightConn struct {
	// Amount defines the maximum amount of allowed simultaneous connections.
	//
	// The middleware closes the connection if there are already amount connections opened.
	Amount *float64 `field:"optional" json:"amount" yaml:"amount"`
}

