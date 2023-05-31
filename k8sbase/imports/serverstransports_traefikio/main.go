// serverstransports_traefikio
package serverstransports_traefikio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"serverstransports_traefikio.ServersTransport",
		reflect.TypeOf((*ServersTransport)(nil)).Elem(),
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
			j := jsiiProxy_ServersTransport{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"serverstransports_traefikio.ServersTransportProps",
		reflect.TypeOf((*ServersTransportProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"serverstransports_traefikio.ServersTransportSpec",
		reflect.TypeOf((*ServersTransportSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"serverstransports_traefikio.ServersTransportSpecForwardingTimeouts",
		reflect.TypeOf((*ServersTransportSpecForwardingTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"serverstransports_traefikio.ServersTransportSpecForwardingTimeoutsDialTimeout",
		reflect.TypeOf((*ServersTransportSpecForwardingTimeoutsDialTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ServersTransportSpecForwardingTimeoutsDialTimeout{}
		},
	)
	_jsii_.RegisterClass(
		"serverstransports_traefikio.ServersTransportSpecForwardingTimeoutsIdleConnTimeout",
		reflect.TypeOf((*ServersTransportSpecForwardingTimeoutsIdleConnTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ServersTransportSpecForwardingTimeoutsIdleConnTimeout{}
		},
	)
	_jsii_.RegisterClass(
		"serverstransports_traefikio.ServersTransportSpecForwardingTimeoutsPingTimeout",
		reflect.TypeOf((*ServersTransportSpecForwardingTimeoutsPingTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ServersTransportSpecForwardingTimeoutsPingTimeout{}
		},
	)
	_jsii_.RegisterClass(
		"serverstransports_traefikio.ServersTransportSpecForwardingTimeoutsReadIdleTimeout",
		reflect.TypeOf((*ServersTransportSpecForwardingTimeoutsReadIdleTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ServersTransportSpecForwardingTimeoutsReadIdleTimeout{}
		},
	)
	_jsii_.RegisterClass(
		"serverstransports_traefikio.ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout",
		reflect.TypeOf((*ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout{}
		},
	)
}