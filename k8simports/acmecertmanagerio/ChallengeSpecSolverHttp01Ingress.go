package acmecertmanagerio


// The ingress based HTTP01 challenge solver will solve challenges by creating or modifying Ingress resources in order to route requests for '/.well-known/acme-challenge/XYZ' to 'challenge solver' pods that are provisioned by cert-manager for each Challenge to be completed.
type ChallengeSpecSolverHttp01Ingress struct {
	// This field configures the annotation `kubernetes.io/ingress.class` when creating Ingress resources to solve ACME challenges that use this challenge solver. Only one of `class`, `name` or `ingressClassName` may be specified.
	Class *string `field:"optional" json:"class" yaml:"class"`
	// This field configures the field `ingressClassName` on the created Ingress resources used to solve ACME challenges that use this challenge solver.
	//
	// This is the recommended way of configuring the ingress class. Only one of `class`, `name` or `ingressClassName` may be specified.
	IngressClassName *string `field:"optional" json:"ingressClassName" yaml:"ingressClassName"`
	// Optional ingress template used to configure the ACME challenge solver ingress used for HTTP01 challenges.
	IngressTemplate *ChallengeSpecSolverHttp01IngressIngressTemplate `field:"optional" json:"ingressTemplate" yaml:"ingressTemplate"`
	// The name of the ingress resource that should have ACME challenge solving routes inserted into it in order to solve HTTP01 challenges.
	//
	// This is typically used in conjunction with ingress controllers like ingress-gce, which maintains a 1:1 mapping between external IPs and ingress resources. Only one of `class`, `name` or `ingressClassName` may be specified.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional pod template used to configure the ACME challenge solver pods used for HTTP01 challenges.
	PodTemplate *ChallengeSpecSolverHttp01IngressPodTemplate `field:"optional" json:"podTemplate" yaml:"podTemplate"`
	// Optional service type for Kubernetes solver service.
	//
	// Supported values are NodePort or ClusterIP. If unset, defaults to NodePort.
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
}

