package certmanagerio


// Configures cert-manager to attempt to complete authorizations by performing the HTTP01 challenge flow.
//
// It is not possible to obtain certificates for wildcard domain names (e.g. `*.example.com`) using the HTTP01 challenge mechanism.
type IssuerSpecAcmeSolversHttp01 struct {
	// The Gateway API is a sig-network community API that models service networking in Kubernetes (https://gateway-api.sigs.k8s.io/). The Gateway solver will create HTTPRoutes with the specified labels in the same namespace as the challenge. This solver is experimental, and fields / behaviour may change in the future.
	GatewayHttpRoute *IssuerSpecAcmeSolversHttp01GatewayHttpRoute `field:"optional" json:"gatewayHttpRoute" yaml:"gatewayHttpRoute"`
	// The ingress based HTTP01 challenge solver will solve challenges by creating or modifying Ingress resources in order to route requests for '/.well-known/acme-challenge/XYZ' to 'challenge solver' pods that are provisioned by cert-manager for each Challenge to be completed.
	Ingress *IssuerSpecAcmeSolversHttp01Ingress `field:"optional" json:"ingress" yaml:"ingress"`
}

