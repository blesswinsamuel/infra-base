// cert-managerio
package certmanagerio


// ObjectMeta overrides for the pod used to solve HTTP01 challenges.
//
// Only the 'labels' and 'annotations' fields may be set. If labels or annotations overlap with in-built values, the values here will override the in-built values.
type ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateMetadata struct {
	// Annotations that should be added to the create ACME HTTP01 solver pods.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Labels that should be added to the created ACME HTTP01 solver pods.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

