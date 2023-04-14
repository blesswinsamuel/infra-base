// cert-managerio
package certmanagerio


// If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node.
//
// If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.
type IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution struct {
	// Required.
	//
	// A list of node selector terms. The terms are ORed.
	NodeSelectorTerms *[]*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms `field:"required" json:"nodeSelectorTerms" yaml:"nodeSelectorTerms"`
}

