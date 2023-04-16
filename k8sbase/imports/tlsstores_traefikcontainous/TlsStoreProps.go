// tlsstores_traefikcontainous
package tlsstores_traefikcontainous

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// TLSStore is the CRD implementation of a Traefik TLS Store.
//
// For the time being, only the TLSStore named default is supported. This means that you cannot have two stores that are named default in different Kubernetes namespaces. More info: https://doc.traefik.io/traefik/v2.9/https/tls/#certificates-stores
type TlsStoreProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// TLSStoreSpec defines the desired state of a TLSStore.
	Spec *TlsStoreSpec `field:"required" json:"spec" yaml:"spec"`
}

