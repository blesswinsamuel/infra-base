// middlewarestcp_traefikio
package middlewarestcp_traefikio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"middlewarestcp_traefikio.MiddlewareTcp",
		reflect.TypeOf((*MiddlewareTcp)(nil)).Elem(),
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
			j := jsiiProxy_MiddlewareTcp{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"middlewarestcp_traefikio.MiddlewareTcpProps",
		reflect.TypeOf((*MiddlewareTcpProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewarestcp_traefikio.MiddlewareTcpSpec",
		reflect.TypeOf((*MiddlewareTcpSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewarestcp_traefikio.MiddlewareTcpSpecInFlightConn",
		reflect.TypeOf((*MiddlewareTcpSpecInFlightConn)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewarestcp_traefikio.MiddlewareTcpSpecIpWhiteList",
		reflect.TypeOf((*MiddlewareTcpSpecIpWhiteList)(nil)).Elem(),
	)
}
