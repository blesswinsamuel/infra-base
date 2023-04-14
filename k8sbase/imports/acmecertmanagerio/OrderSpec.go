// acmecert-managerio
package acmecertmanagerio


type OrderSpec struct {
	// IssuerRef references a properly configured ACME-type Issuer which should be used to create this Order.
	//
	// If the Issuer does not exist, processing will be retried. If the Issuer is not an 'ACME' Issuer, an error will be returned and the Order will be marked as failed.
	IssuerRef *OrderSpecIssuerRef `field:"required" json:"issuerRef" yaml:"issuerRef"`
	// Certificate signing request bytes in DER encoding.
	//
	// This will be used when finalizing the order. This field must be set on the order.
	Request *string `field:"required" json:"request" yaml:"request"`
	// CommonName is the common name as specified on the DER encoded CSR.
	//
	// If specified, this value must also be present in `dnsNames` or `ipAddresses`. This field must match the corresponding field on the DER encoded CSR.
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// DNSNames is a list of DNS names that should be included as part of the Order validation process.
	//
	// This field must match the corresponding field on the DER encoded CSR.
	DnsNames *[]*string `field:"optional" json:"dnsNames" yaml:"dnsNames"`
	// Duration is the duration for the not after date for the requested certificate.
	//
	// this is set on order creation as pe the ACME spec.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// IPAddresses is a list of IP addresses that should be included as part of the Order validation process.
	//
	// This field must match the corresponding field on the DER encoded CSR.
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
}

