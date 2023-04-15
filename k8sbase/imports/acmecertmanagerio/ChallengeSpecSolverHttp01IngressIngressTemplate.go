// acmecert-managerio
package acmecertmanagerio


// Optional ingress template used to configure the ACME challenge solver ingress used for HTTP01 challenges.
type ChallengeSpecSolverHttp01IngressIngressTemplate struct {
	// ObjectMeta overrides for the ingress used to solve HTTP01 challenges.
	//
	// Only the 'labels' and 'annotations' fields may be set. If labels or annotations overlap with in-built values, the values here will override the in-built values.
	Metadata *ChallengeSpecSolverHttp01IngressIngressTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}
