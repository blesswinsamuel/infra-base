package certmanagerio


// PodSpec defines overrides for the HTTP01 challenge solver pod.
//
// Check ACMEChallengeSolverHTTP01IngressPodSpec to find out currently supported fields. All other fields will be ignored.
type ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpec struct {
	// If specified, the pod's scheduling constraints.
	Affinity *ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	// If specified, the pod's imagePullSecrets.
	ImagePullSecrets *[]*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecImagePullSecrets `field:"optional" json:"imagePullSecrets" yaml:"imagePullSecrets"`
	// NodeSelector is a selector which must be true for the pod to fit on a node.
	//
	// Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	// If specified, the pod's priorityClassName.
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	// If specified, the pod's service account.
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// If specified, the pod's tolerations.
	Tolerations *[]*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecTolerations `field:"optional" json:"tolerations" yaml:"tolerations"`
}

