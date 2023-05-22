package certmanagerio


// PodSpec defines overrides for the HTTP01 challenge solver pod.
//
// Only the 'priorityClassName', 'nodeSelector', 'affinity', 'serviceAccountName' and 'tolerations' fields are supported currently. All other fields will be ignored.
type IssuerSpecAcmeSolversHttp01IngressPodTemplateSpec struct {
	// If specified, the pod's scheduling constraints.
	Affinity *IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinity `field:"optional" json:"affinity" yaml:"affinity"`
	// NodeSelector is a selector which must be true for the pod to fit on a node.
	//
	// Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	NodeSelector *map[string]*string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	// If specified, the pod's priorityClassName.
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	// If specified, the pod's service account.
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// If specified, the pod's tolerations.
	Tolerations *[]*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecTolerations `field:"optional" json:"tolerations" yaml:"tolerations"`
}

