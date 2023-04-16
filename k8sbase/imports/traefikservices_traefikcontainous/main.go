package traefikservices_traefikcontainous

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"traefikservices_traefikcontainous.TraefikService",
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
		"traefikservices_traefikcontainous.TraefikServiceProps",
		reflect.TypeOf((*TraefikServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpec",
		reflect.TypeOf((*TraefikServiceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecMirroring",
		reflect.TypeOf((*TraefikServiceSpecMirroring)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"traefikservices_traefikcontainous.TraefikServiceSpecMirroringKind",
		reflect.TypeOf((*TraefikServiceSpecMirroringKind)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE": TraefikServiceSpecMirroringKind_SERVICE,
			"TRAEFIK_SERVICE": TraefikServiceSpecMirroringKind_TRAEFIK_SERVICE,
		},
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecMirroringMirrors",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrors)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"traefikservices_traefikcontainous.TraefikServiceSpecMirroringMirrorsKind",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsKind)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE": TraefikServiceSpecMirroringMirrorsKind_SERVICE,
			"TRAEFIK_SERVICE": TraefikServiceSpecMirroringMirrorsKind_TRAEFIK_SERVICE,
		},
	)
	_jsii_.RegisterClass(
		"traefikservices_traefikcontainous.TraefikServiceSpecMirroringMirrorsPort",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecMirroringMirrorsPort{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecMirroringMirrorsResponseForwarding",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsResponseForwarding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecMirroringMirrorsSticky",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecMirroringMirrorsStickyCookie",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsStickyCookie)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikservices_traefikcontainous.TraefikServiceSpecMirroringPort",
		reflect.TypeOf((*TraefikServiceSpecMirroringPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecMirroringPort{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecMirroringResponseForwarding",
		reflect.TypeOf((*TraefikServiceSpecMirroringResponseForwarding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecMirroringSticky",
		reflect.TypeOf((*TraefikServiceSpecMirroringSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecMirroringStickyCookie",
		reflect.TypeOf((*TraefikServiceSpecMirroringStickyCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecWeighted",
		reflect.TypeOf((*TraefikServiceSpecWeighted)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecWeightedServices",
		reflect.TypeOf((*TraefikServiceSpecWeightedServices)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"traefikservices_traefikcontainous.TraefikServiceSpecWeightedServicesKind",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesKind)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE": TraefikServiceSpecWeightedServicesKind_SERVICE,
			"TRAEFIK_SERVICE": TraefikServiceSpecWeightedServicesKind_TRAEFIK_SERVICE,
		},
	)
	_jsii_.RegisterClass(
		"traefikservices_traefikcontainous.TraefikServiceSpecWeightedServicesPort",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecWeightedServicesPort{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecWeightedServicesResponseForwarding",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesResponseForwarding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecWeightedServicesSticky",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecWeightedServicesStickyCookie",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesStickyCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecWeightedSticky",
		reflect.TypeOf((*TraefikServiceSpecWeightedSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikservices_traefikcontainous.TraefikServiceSpecWeightedStickyCookie",
		reflect.TypeOf((*TraefikServiceSpecWeightedStickyCookie)(nil)).Elem(),
	)
}
