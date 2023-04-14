// ingressroutetcp_traefikcontainous
package ingressroutetcp_traefikcontainous


// TLS defines the TLS configuration on a layer 4 / TCP Route.
//
// More info: https://doc.traefik.io/traefik/v2.9/routing/routers/#tls_1
type IngressRouteTcpSpecTls struct {
	// CertResolver defines the name of the certificate resolver to use.
	//
	// Cert resolvers have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v2.9/https/acme/#certificate-resolvers
	CertResolver *string `field:"optional" json:"certResolver" yaml:"certResolver"`
	// Domains defines the list of domains that will be used to issue certificates.
	//
	// More info: https://doc.traefik.io/traefik/v2.9/routing/routers/#domains
	Domains *[]*IngressRouteTcpSpecTlsDomains `field:"optional" json:"domains" yaml:"domains"`
	// Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection.
	//
	// If not defined, the `default` TLSOption is used. More info: https://doc.traefik.io/traefik/v2.9/https/tls/#tls-options
	Options *IngressRouteTcpSpecTlsOptions `field:"optional" json:"options" yaml:"options"`
	// Passthrough defines whether a TLS router will terminate the TLS connection.
	Passthrough *bool `field:"optional" json:"passthrough" yaml:"passthrough"`
	// SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// Store defines the reference to the TLSStore, that will be used to store certificates.
	//
	// Please note that only `default` TLSStore can be used.
	Store *IngressRouteTcpSpecTlsStore `field:"optional" json:"store" yaml:"store"`
}

