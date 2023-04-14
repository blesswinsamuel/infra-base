package ingressroute_traefikcontainous

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"ingressroute_traefikcontainous.IngressRoute",
		reflect.TypeOf((*IngressRoute)(nil)).Elem(),
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
			j := jsiiProxy_IngressRoute{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"ingressroute_traefikcontainous.IngressRouteProps",
		reflect.TypeOf((*IngressRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroute_traefikcontainous.IngressRouteSpec",
		reflect.TypeOf((*IngressRouteSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroute_traefikcontainous.IngressRouteSpecRoutes",
		reflect.TypeOf((*IngressRouteSpecRoutes)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ingressroute_traefikcontainous.IngressRouteSpecRoutesKind",
		reflect.TypeOf((*IngressRouteSpecRoutesKind)(nil)).Elem(),
		map[string]interface{}{
			"RULE": IngressRouteSpecRoutesKind_RULE,
		},
	)
	_jsii_.RegisterStruct(
		"ingressroute_traefikcontainous.IngressRouteSpecRoutesMiddlewares",
		reflect.TypeOf((*IngressRouteSpecRoutesMiddlewares)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroute_traefikcontainous.IngressRouteSpecRoutesServices",
		reflect.TypeOf((*IngressRouteSpecRoutesServices)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"ingressroute_traefikcontainous.IngressRouteSpecRoutesServicesKind",
		reflect.TypeOf((*IngressRouteSpecRoutesServicesKind)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE": IngressRouteSpecRoutesServicesKind_SERVICE,
			"TRAEFIK_SERVICE": IngressRouteSpecRoutesServicesKind_TRAEFIK_SERVICE,
		},
	)
	_jsii_.RegisterClass(
		"ingressroute_traefikcontainous.IngressRouteSpecRoutesServicesPort",
		reflect.TypeOf((*IngressRouteSpecRoutesServicesPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IngressRouteSpecRoutesServicesPort{}
		},
	)
	_jsii_.RegisterStruct(
		"ingressroute_traefikcontainous.IngressRouteSpecRoutesServicesResponseForwarding",
		reflect.TypeOf((*IngressRouteSpecRoutesServicesResponseForwarding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroute_traefikcontainous.IngressRouteSpecRoutesServicesSticky",
		reflect.TypeOf((*IngressRouteSpecRoutesServicesSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroute_traefikcontainous.IngressRouteSpecRoutesServicesStickyCookie",
		reflect.TypeOf((*IngressRouteSpecRoutesServicesStickyCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroute_traefikcontainous.IngressRouteSpecTls",
		reflect.TypeOf((*IngressRouteSpecTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroute_traefikcontainous.IngressRouteSpecTlsDomains",
		reflect.TypeOf((*IngressRouteSpecTlsDomains)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroute_traefikcontainous.IngressRouteSpecTlsOptions",
		reflect.TypeOf((*IngressRouteSpecTlsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"ingressroute_traefikcontainous.IngressRouteSpecTlsStore",
		reflect.TypeOf((*IngressRouteSpecTlsStore)(nil)).Elem(),
	)
}
