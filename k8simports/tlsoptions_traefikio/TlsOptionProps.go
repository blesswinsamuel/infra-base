package tlsoptions_traefikio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// TLSOption is the CRD implementation of a Traefik TLS Option, allowing to configure some parameters of the TLS connection.
//
// More info: https://doc.traefik.io/traefik/v2.10/https/tls/#tls-options
type TlsOptionProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// TLSOptionSpec defines the desired state of a TLSOption.
	Spec *TlsOptionSpec `field:"required" json:"spec" yaml:"spec"`
}

