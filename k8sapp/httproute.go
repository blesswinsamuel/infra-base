package k8sapp

import (
	"github.com/blesswinsamuel/kgen"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

type HTTPRoutePath struct {
	Path              string `json:"path"`
	ServiceName       string `json:"serviceName"`
	ServicePortNumber int32  `json:"servicePortNumber"`
}

type HTTPRouteHost struct {
	Host  string          `json:"host"`
	Paths []HTTPRoutePath `json:"paths"`
}

type HTTPRoute struct {
	Name      string          `json:"name"`
	Gateway   NameNamespace   `json:"gateway"`
	Hostnames []string        `json:"hostnames"`
	Hosts     []HTTPRouteHost `json:"hosts"`
}

func NewHTTPRoute(scope kgen.Scope, props *HTTPRoute) kgen.Scope {
	hostRules := []gatewayv1.HTTPRouteRule{}
	for _, host := range props.Hosts {
		backendRefs := []gatewayv1.HTTPBackendRef{}
		for _, path := range host.Paths {
			backendRefs = append(backendRefs, gatewayv1.HTTPBackendRef{
				BackendRef: gatewayv1.BackendRef{
					BackendObjectReference: gatewayv1.BackendObjectReference{
						Name: gatewayv1.ObjectName(path.ServiceName),
						Port: ptr.To(gatewayv1.PortNumber(path.ServicePortNumber)),
					},
				},
			})
		}
		hostRules = append(hostRules, gatewayv1.HTTPRouteRule{
			Matches: []gatewayv1.HTTPRouteMatch{
				{Path: &gatewayv1.HTTPPathMatch{Type: ptr.To(gatewayv1.PathMatchPathPrefix), Value: ptr.To("/")}},
			},
			BackendRefs: backendRefs,
		})
	}
	hostnames := []gatewayv1.Hostname{}
	for _, hostname := range props.Hostnames {
		hostnames = append(hostnames, gatewayv1.Hostname(hostname))
	}
	if props.Gateway.Namespace == "" {
		props.Gateway.Namespace = scope.Namespace()
	}
	scope.AddApiObject(&gatewayv1.HTTPRoute{
		ObjectMeta: v1.ObjectMeta{
			Name: props.Name,
		},
		Spec: gatewayv1.HTTPRouteSpec{
			Hostnames: hostnames,
			CommonRouteSpec: gatewayv1.CommonRouteSpec{
				ParentRefs: []gatewayv1.ParentReference{
					{Name: gatewayv1.ObjectName(props.Gateway.Name), Namespace: ptr.To(gatewayv1.Namespace(props.Gateway.Namespace))},
				},
			},
			Rules: hostRules,
		},
	})

	return scope
}
