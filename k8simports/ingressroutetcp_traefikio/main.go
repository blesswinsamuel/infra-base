// ingressroutetcp_traefikio
package ingressroutetcp_traefikio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"ingressroutetcp_traefikio.IngressRouteTcp",
		reflect.TypeOf((*IngressRouteTcp)(nil)).Elem(),
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
			j := jsiiProxy_IngressRouteTcp{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"ingressroutetcp_traefikio.IngressRouteTcpProps",
		reflect.TypeOf((*IngressRouteTcpProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroutetcp_traefikio.IngressRouteTcpSpec",
		reflect.TypeOf((*IngressRouteTcpSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroutetcp_traefikio.IngressRouteTcpSpecRoutes",
		reflect.TypeOf((*IngressRouteTcpSpecRoutes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroutetcp_traefikio.IngressRouteTcpSpecRoutesMiddlewares",
		reflect.TypeOf((*IngressRouteTcpSpecRoutesMiddlewares)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroutetcp_traefikio.IngressRouteTcpSpecRoutesServices",
		reflect.TypeOf((*IngressRouteTcpSpecRoutesServices)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"ingressroutetcp_traefikio.IngressRouteTcpSpecRoutesServicesPort",
		reflect.TypeOf((*IngressRouteTcpSpecRoutesServicesPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IngressRouteTcpSpecRoutesServicesPort{}
		},
	)
	_jsii_.RegisterStruct(
		"ingressroutetcp_traefikio.IngressRouteTcpSpecRoutesServicesProxyProtocol",
		reflect.TypeOf((*IngressRouteTcpSpecRoutesServicesProxyProtocol)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroutetcp_traefikio.IngressRouteTcpSpecTls",
		reflect.TypeOf((*IngressRouteTcpSpecTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroutetcp_traefikio.IngressRouteTcpSpecTlsDomains",
		reflect.TypeOf((*IngressRouteTcpSpecTlsDomains)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroutetcp_traefikio.IngressRouteTcpSpecTlsOptions",
		reflect.TypeOf((*IngressRouteTcpSpecTlsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroutetcp_traefikio.IngressRouteTcpSpecTlsStore",
		reflect.TypeOf((*IngressRouteTcpSpecTlsStore)(nil)).Elem(),
	)
}
