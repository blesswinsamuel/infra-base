package certmanagerio


// The Gateway API is a sig-network community API that models service networking in Kubernetes (https://gateway-api.sigs.k8s.io/). The Gateway solver will create HTTPRoutes with the specified labels in the same namespace as the challenge. This solver is experimental, and fields / behaviour may change in the future.
type ClusterIssuerSpecAcmeSolversHttp01GatewayHttpRoute struct {
	// Custom labels that will be applied to HTTPRoutes created by cert-manager while solving HTTP-01 challenges.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// When solving an HTTP-01 challenge, cert-manager creates an HTTPRoute.
	//
	// cert-manager needs to know which parentRefs should be used when creating the HTTPRoute. Usually, the parentRef references a Gateway. See: https://gateway-api.sigs.k8s.io/api-types/httproute/#attaching-to-gateways
	ParentRefs *[]*ClusterIssuerSpecAcmeSolversHttp01GatewayHttpRouteParentRefs `field:"optional" json:"parentRefs" yaml:"parentRefs"`
	// Optional service type for Kubernetes solver service.
	//
	// Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
}

