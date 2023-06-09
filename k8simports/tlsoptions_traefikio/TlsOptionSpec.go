package tlsoptions_traefikio


// TLSOptionSpec defines the desired state of a TLSOption.
type TlsOptionSpec struct {
	// ALPNProtocols defines the list of supported application level protocols for the TLS handshake, in order of preference.
	//
	// More info: https://doc.traefik.io/traefik/v2.10/https/tls/#alpn-protocols
	AlpnProtocols *[]*string `field:"optional" json:"alpnProtocols" yaml:"alpnProtocols"`
	// CipherSuites defines the list of supported cipher suites for TLS versions up to TLS 1.2. More info: https://doc.traefik.io/traefik/v2.10/https/tls/#cipher-suites.
	CipherSuites *[]*string `field:"optional" json:"cipherSuites" yaml:"cipherSuites"`
	// ClientAuth defines the server's policy for TLS Client Authentication.
	ClientAuth *TlsOptionSpecClientAuth `field:"optional" json:"clientAuth" yaml:"clientAuth"`
	// CurvePreferences defines the preferred elliptic curves in a specific order.
	//
	// More info: https://doc.traefik.io/traefik/v2.10/https/tls/#curve-preferences
	CurvePreferences *[]*string `field:"optional" json:"curvePreferences" yaml:"curvePreferences"`
	// MaxVersion defines the maximum TLS version that Traefik will accept.
	//
	// Possible values: VersionTLS10, VersionTLS11, VersionTLS12, VersionTLS13. Default: None.
	MaxVersion *string `field:"optional" json:"maxVersion" yaml:"maxVersion"`
	// MinVersion defines the minimum TLS version that Traefik will accept.
	//
	// Possible values: VersionTLS10, VersionTLS11, VersionTLS12, VersionTLS13. Default: VersionTLS10.
	MinVersion *string `field:"optional" json:"minVersion" yaml:"minVersion"`
	// PreferServerCipherSuites defines whether the server chooses a cipher suite among his own instead of among the client's.
	//
	// It is enabled automatically when minVersion or maxVersion is set. Deprecated: https://github.com/golang/go/issues/45430
	PreferServerCipherSuites *bool `field:"optional" json:"preferServerCipherSuites" yaml:"preferServerCipherSuites"`
	// SniStrict defines whether Traefik allows connections from clients connections that do not specify a server_name extension.
	SniStrict *bool `field:"optional" json:"sniStrict" yaml:"sniStrict"`
}

