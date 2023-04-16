// tlsoptions_traefikcontainous
package tlsoptions_traefikcontainous


// ClientAuth defines the server's policy for TLS Client Authentication.
type TlsOptionSpecClientAuth struct {
	// ClientAuthType defines the client authentication type to apply.
	ClientAuthType TlsOptionSpecClientAuthClientAuthType `field:"optional" json:"clientAuthType" yaml:"clientAuthType"`
	// SecretNames defines the names of the referenced Kubernetes Secret storing certificate details.
	SecretNames *[]*string `field:"optional" json:"secretNames" yaml:"secretNames"`
}

