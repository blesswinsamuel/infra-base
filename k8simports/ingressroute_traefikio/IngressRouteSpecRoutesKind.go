package ingressroute_traefikio


// Kind defines the kind of the route.
//
// Rule is the only supported kind.
type IngressRouteSpecRoutesKind string

const (
	// Rule.
	IngressRouteSpecRoutesKind_RULE IngressRouteSpecRoutesKind = "RULE"
)

