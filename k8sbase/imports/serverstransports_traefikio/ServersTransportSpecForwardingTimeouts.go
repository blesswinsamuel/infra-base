package serverstransports_traefikio


// ForwardingTimeouts defines the timeouts for requests forwarded to the backend servers.
type ServersTransportSpecForwardingTimeouts struct {
	// DialTimeout is the amount of time to wait until a connection to a backend server can be established.
	DialTimeout ServersTransportSpecForwardingTimeoutsDialTimeout `field:"optional" json:"dialTimeout" yaml:"dialTimeout"`
	// IdleConnTimeout is the maximum period for which an idle HTTP keep-alive connection will remain open before closing itself.
	IdleConnTimeout ServersTransportSpecForwardingTimeoutsIdleConnTimeout `field:"optional" json:"idleConnTimeout" yaml:"idleConnTimeout"`
	// PingTimeout is the timeout after which the HTTP/2 connection will be closed if a response to ping is not received.
	PingTimeout ServersTransportSpecForwardingTimeoutsPingTimeout `field:"optional" json:"pingTimeout" yaml:"pingTimeout"`
	// ReadIdleTimeout is the timeout after which a health check using ping frame will be carried out if no frame is received on the HTTP/2 connection.
	ReadIdleTimeout ServersTransportSpecForwardingTimeoutsReadIdleTimeout `field:"optional" json:"readIdleTimeout" yaml:"readIdleTimeout"`
	// ResponseHeaderTimeout is the amount of time to wait for a server's response headers after fully writing the request (including its body, if any).
	ResponseHeaderTimeout ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout `field:"optional" json:"responseHeaderTimeout" yaml:"responseHeaderTimeout"`
}

