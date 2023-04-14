// cert-managerio
package certmanagerio


// The ingress based HTTP01 challenge solver will solve challenges by creating or modifying Ingress resources in order to route requests for '/.well-known/acme-challenge/XYZ' to 'challenge solver' pods that are provisioned by cert-manager for each Challenge to be completed.
type IssuerSpecAcmeSolversHttp01Ingress struct {
	// The ingress class to use when creating Ingress resources to solve ACME challenges that use this challenge solver.
	//
	// Only one of 'class' or 'name' may be specified.
	Class *string `field:"optional" json:"class" yaml:"class"`
	// Optional ingress template used to configure the ACME challenge solver ingress used for HTTP01 challenges.
	IngressTemplate *IssuerSpecAcmeSolversHttp01IngressIngressTemplate `field:"optional" json:"ingressTemplate" yaml:"ingressTemplate"`
	// The name of the ingress resource that should have ACME challenge solving routes inserted into it in order to solve HTTP01 challenges.
	//
	// This is typically used in conjunction with ingress controllers like ingress-gce, which maintains a 1:1 mapping between external IPs and ingress resources.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional pod template used to configure the ACME challenge solver pods used for HTTP01 challenges.
	PodTemplate *IssuerSpecAcmeSolversHttp01IngressPodTemplate `field:"optional" json:"podTemplate" yaml:"podTemplate"`
	// Optional service type for Kubernetes solver service.
	//
	// Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
}

