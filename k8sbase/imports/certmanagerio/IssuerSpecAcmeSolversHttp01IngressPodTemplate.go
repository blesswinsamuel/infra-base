package certmanagerio


// Optional pod template used to configure the ACME challenge solver pods used for HTTP01 challenges.
type IssuerSpecAcmeSolversHttp01IngressPodTemplate struct {
	// ObjectMeta overrides for the pod used to solve HTTP01 challenges.
	//
	// Only the 'labels' and 'annotations' fields may be set. If labels or annotations overlap with in-built values, the values here will override the in-built values.
	Metadata *IssuerSpecAcmeSolversHttp01IngressPodTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// PodSpec defines overrides for the HTTP01 challenge solver pod.
	//
	// Only the 'priorityClassName', 'nodeSelector', 'affinity', 'serviceAccountName' and 'tolerations' fields are supported currently. All other fields will be ignored.
	Spec *IssuerSpecAcmeSolversHttp01IngressPodTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

