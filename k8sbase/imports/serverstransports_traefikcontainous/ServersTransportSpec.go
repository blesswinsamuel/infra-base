// serverstransports_traefikcontainous
package serverstransports_traefikcontainous


// ServersTransportSpec defines the desired state of a ServersTransport.
type ServersTransportSpec struct {
	// CertificatesSecrets defines a list of secret storing client certificates for mTLS.
	CertificatesSecrets *[]*string `field:"optional" json:"certificatesSecrets" yaml:"certificatesSecrets"`
	// DisableHTTP2 disables HTTP/2 for connections with backend servers.
	DisableHttp2 *bool `field:"optional" json:"disableHttp2" yaml:"disableHttp2"`
	// ForwardingTimeouts defines the timeouts for requests forwarded to the backend servers.
	ForwardingTimeouts *ServersTransportSpecForwardingTimeouts `field:"optional" json:"forwardingTimeouts" yaml:"forwardingTimeouts"`
	// InsecureSkipVerify disables SSL certificate verification.
	InsecureSkipVerify *bool `field:"optional" json:"insecureSkipVerify" yaml:"insecureSkipVerify"`
	// MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host.
	MaxIdleConnsPerHost *float64 `field:"optional" json:"maxIdleConnsPerHost" yaml:"maxIdleConnsPerHost"`
	// PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.
	PeerCertUri *string `field:"optional" json:"peerCertUri" yaml:"peerCertUri"`
	// RootCAsSecrets defines a list of CA secret used to validate self-signed certificate.
	RootCAsSecrets *[]*string `field:"optional" json:"rootCAsSecrets" yaml:"rootCAsSecrets"`
	// ServerName defines the server name used to contact the server.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
}

