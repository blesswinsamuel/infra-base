package acmecertmanagerio


// Optional pod template used to configure the ACME challenge solver pods used for HTTP01 challenges.
type ChallengeSpecSolverHttp01IngressPodTemplate struct {
	// ObjectMeta overrides for the pod used to solve HTTP01 challenges.
	//
	// Only the 'labels' and 'annotations' fields may be set. If labels or annotations overlap with in-built values, the values here will override the in-built values.
	Metadata *ChallengeSpecSolverHttp01IngressPodTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// PodSpec defines overrides for the HTTP01 challenge solver pod.
	//
	// Check ACMEChallengeSolverHTTP01IngressPodSpec to find out currently supported fields. All other fields will be ignored.
	Spec *ChallengeSpecSolverHttp01IngressPodTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

