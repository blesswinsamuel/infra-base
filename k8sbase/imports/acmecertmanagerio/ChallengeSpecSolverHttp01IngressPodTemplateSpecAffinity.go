// acmecert-managerio
package acmecertmanagerio


// If specified, the pod's scheduling constraints.
type ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinity struct {
	// Describes node affinity scheduling rules for the pod.
	NodeAffinity *ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	// Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).
	PodAffinity *ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	// Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).
	PodAntiAffinity *ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

