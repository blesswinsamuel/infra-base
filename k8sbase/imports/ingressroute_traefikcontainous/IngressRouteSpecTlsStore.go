// ingressroute_traefikcontainous
package ingressroute_traefikcontainous


// Store defines the reference to the TLSStore, that will be used to store certificates.
//
// Please note that only `default` TLSStore can be used.
type IngressRouteSpecTlsStore struct {
	// Name defines the name of the referenced TLSStore.
	//
	// More info: https://doc.traefik.io/traefik/v2.9/routing/providers/kubernetes-crd/#kind-tlsstore
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace defines the namespace of the referenced TLSStore.
	//
	// More info: https://doc.traefik.io/traefik/v2.9/routing/providers/kubernetes-crd/#kind-tlsstore
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

