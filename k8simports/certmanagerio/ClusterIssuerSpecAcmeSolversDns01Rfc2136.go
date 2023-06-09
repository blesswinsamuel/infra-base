package certmanagerio


// Use RFC2136 ("Dynamic Updates in the Domain Name System") (https://datatracker.ietf.org/doc/rfc2136/) to manage DNS01 challenge records.
type ClusterIssuerSpecAcmeSolversDns01Rfc2136 struct {
	// The IP address or hostname of an authoritative DNS server supporting RFC2136 in the form host:port.
	//
	// If the host is an IPv6 address it must be enclosed in square brackets (e.g [2001:db8::1]) ; port is optional. This field is required.
	Nameserver *string `field:"required" json:"nameserver" yaml:"nameserver"`
	// The TSIG Algorithm configured in the DNS supporting RFC2136.
	//
	// Used only when ``tsigSecretSecretRef`` and ``tsigKeyName`` are defined. Supported values are (case-insensitive): ``HMACMD5`` (default), ``HMACSHA1``, ``HMACSHA256`` or ``HMACSHA512``.
	TsigAlgorithm *string `field:"optional" json:"tsigAlgorithm" yaml:"tsigAlgorithm"`
	// The TSIG Key name configured in the DNS.
	//
	// If ``tsigSecretSecretRef`` is defined, this field is required.
	TsigKeyName *string `field:"optional" json:"tsigKeyName" yaml:"tsigKeyName"`
	// The name of the secret containing the TSIG value.
	//
	// If ``tsigKeyName`` is defined, this field is required.
	TsigSecretSecretRef *ClusterIssuerSpecAcmeSolversDns01Rfc2136TsigSecretSecretRef `field:"optional" json:"tsigSecretSecretRef" yaml:"tsigSecretSecretRef"`
}

