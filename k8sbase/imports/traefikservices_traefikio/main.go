// traefikservices_traefikio
package traefikservices_traefikio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"traefikservices_traefikio.TraefikService",
		reflect.TypeOf((*TraefikService)(nil)).Elem(),
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
			j := jsiiProxy_TraefikService{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceProps",
		reflect.TypeOf((*TraefikServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpec",
		reflect.TypeOf((*TraefikServiceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecMirroring",
		reflect.TypeOf((*TraefikServiceSpecMirroring)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"traefikservices_traefikio.TraefikServiceSpecMirroringKind",
		reflect.TypeOf((*TraefikServiceSpecMirroringKind)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE": TraefikServiceSpecMirroringKind_SERVICE,
			"TRAEFIK_SERVICE": TraefikServiceSpecMirroringKind_TRAEFIK_SERVICE,
		},
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecMirroringMirrors",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrors)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"traefikservices_traefikio.TraefikServiceSpecMirroringMirrorsKind",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsKind)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE": TraefikServiceSpecMirroringMirrorsKind_SERVICE,
			"TRAEFIK_SERVICE": TraefikServiceSpecMirroringMirrorsKind_TRAEFIK_SERVICE,
		},
	)
	_jsii_.RegisterClass(
		"traefikservices_traefikio.TraefikServiceSpecMirroringMirrorsPort",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecMirroringMirrorsPort{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecMirroringMirrorsResponseForwarding",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsResponseForwarding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecMirroringMirrorsSticky",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecMirroringMirrorsStickyCookie",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsStickyCookie)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikservices_traefikio.TraefikServiceSpecMirroringPort",
		reflect.TypeOf((*TraefikServiceSpecMirroringPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecMirroringPort{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecMirroringResponseForwarding",
		reflect.TypeOf((*TraefikServiceSpecMirroringResponseForwarding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecMirroringSticky",
		reflect.TypeOf((*TraefikServiceSpecMirroringSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecMirroringStickyCookie",
		reflect.TypeOf((*TraefikServiceSpecMirroringStickyCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecWeighted",
		reflect.TypeOf((*TraefikServiceSpecWeighted)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecWeightedServices",
		reflect.TypeOf((*TraefikServiceSpecWeightedServices)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"traefikservices_traefikio.TraefikServiceSpecWeightedServicesKind",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesKind)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE": TraefikServiceSpecWeightedServicesKind_SERVICE,
			"TRAEFIK_SERVICE": TraefikServiceSpecWeightedServicesKind_TRAEFIK_SERVICE,
		},
	)
	_jsii_.RegisterClass(
		"traefikservices_traefikio.TraefikServiceSpecWeightedServicesPort",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecWeightedServicesPort{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecWeightedServicesResponseForwarding",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesResponseForwarding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecWeightedServicesSticky",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecWeightedServicesStickyCookie",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesStickyCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecWeightedSticky",
		reflect.TypeOf((*TraefikServiceSpecWeightedSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikio.TraefikServiceSpecWeightedStickyCookie",
		reflect.TypeOf((*TraefikServiceSpecWeightedStickyCookie)(nil)).Elem(),
	)
}
