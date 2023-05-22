// ingressrouteudp_traefikio
package ingressrouteudp_traefikio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"ingressrouteudp_traefikio.IngressRouteUdp",
		reflect.TypeOf((*IngressRouteUdp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IngressRouteUdp{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"ingressrouteudp_traefikio.IngressRouteUdpProps",
		reflect.TypeOf((*IngressRouteUdpProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressrouteudp_traefikio.IngressRouteUdpSpec",
		reflect.TypeOf((*IngressRouteUdpSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressrouteudp_traefikio.IngressRouteUdpSpecRoutes",
		reflect.TypeOf((*IngressRouteUdpSpecRoutes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressrouteudp_traefikio.IngressRouteUdpSpecRoutesServices",
		reflect.TypeOf((*IngressRouteUdpSpecRoutesServices)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"ingressrouteudp_traefikio.IngressRouteUdpSpecRoutesServicesPort",
		reflect.TypeOf((*IngressRouteUdpSpecRoutesServicesPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IngressRouteUdpSpecRoutesServicesPort{}
		},
	)
}
