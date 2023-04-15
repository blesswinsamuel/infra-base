// cert-managerio
package certmanagerio


// ParentReference identifies an API object (usually a Gateway) that can be considered a parent of this resource (usually a route).
//
// The only kind of parent resource with "Core" support is Gateway. This API may be extended in the future to support additional kinds of parent resources, such as HTTPRoute.
// The API object must be valid in the cluster; the Group and Kind must be registered in the cluster for this reference to be valid.
type IssuerSpecAcmeSolversHttp01GatewayHttpRouteParentRefs struct {
	// Name is the name of the referent.
	//
	// Support: Core.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Group is the group of the referent.
	//
	// When unspecified, "gateway.networking.k8s.io" is inferred. To set the core API group (such as for a "Service" kind referent), Group must be explicitly set to "" (empty string).
	// Support: Core.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Kind is kind of the referent.
	//
	// Support: Core (Gateway)
	// Support: Implementation-specific (Other Resources).
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Namespace is the namespace of the referent.
	//
	// When unspecified, this refers to the local namespace of the Route.
	// Note that there are specific rules for ParentRefs which cross namespace boundaries. Cross-namespace references are only valid if they are explicitly allowed by something in the namespace they are referring to. For example: Gateway has the AllowedRoutes field, and ReferenceGrant provides a generic way to enable any other kind of cross-namespace reference.
	// Support: Core.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Port is the network port this Route targets.
	//
	// It can be interpreted differently based on the type of parent resource.
	// When the parent resource is a Gateway, this targets all listeners listening on the specified port that also support this kind of Route(and select this Route). It's not recommended to set `Port` unless the networking behaviors specified in a Route must apply to a specific port as opposed to a listener(s) whose port(s) may be changed. When both Port and SectionName are specified, the name and port of the selected listener must match both specified values.
	// Implementations MAY choose to support other parent resources. Implementations supporting other types of parent resources MUST clearly document how/if Port is interpreted.
	// For the purpose of status, an attachment is considered successful as long as the parent resource accepts it partially. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.
	// Support: Extended
	// <gateway:experimental>.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// SectionName is the name of a section within the target resource.
	//
	// In the following resources, SectionName is interpreted as the following:
	// * Gateway: Listener Name. When both Port (experimental) and SectionName are specified, the name and port of the selected listener must match both specified values.
	// Implementations MAY choose to support attaching Routes to other resources. If that is the case, they MUST clearly document how SectionName is interpreted.
	// When unspecified (empty string), this will reference the entire resource. For the purpose of status, an attachment is considered successful if at least one section in the parent resource accepts it. For example, Gateway listeners can restrict which Routes can attach to them by Route kind, namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from the referencing Route, the Route MUST be considered successfully attached. If no Gateway listeners accept attachment from this Route, the Route MUST be considered detached from the Gateway.
	// Support: Core.
	SectionName *string `field:"optional" json:"sectionName" yaml:"sectionName"`
}
