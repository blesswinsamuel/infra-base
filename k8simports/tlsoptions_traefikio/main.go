// tlsoptions_traefikio
package tlsoptions_traefikio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"tlsoptions_traefikio.TlsOption",
		reflect.TypeOf((*TlsOption)(nil)).Elem(),
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
			j := jsiiProxy_TlsOption{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"tlsoptions_traefikio.TlsOptionProps",
		reflect.TypeOf((*TlsOptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"tlsoptions_traefikio.TlsOptionSpec",
		reflect.TypeOf((*TlsOptionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"tlsoptions_traefikio.TlsOptionSpecClientAuth",
		reflect.TypeOf((*TlsOptionSpecClientAuth)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"tlsoptions_traefikio.TlsOptionSpecClientAuthClientAuthType",
		reflect.TypeOf((*TlsOptionSpecClientAuthClientAuthType)(nil)).Elem(),
		map[string]interface{}{
			"NO_CLIENT_CERT": TlsOptionSpecClientAuthClientAuthType_NO_CLIENT_CERT,
			"REQUEST_CLIENT_CERT": TlsOptionSpecClientAuthClientAuthType_REQUEST_CLIENT_CERT,
			"REQUIRE_ANY_CLIENT_CERT": TlsOptionSpecClientAuthClientAuthType_REQUIRE_ANY_CLIENT_CERT,
			"VERIFY_CLIENT_CERT_IF_GIVEN": TlsOptionSpecClientAuthClientAuthType_VERIFY_CLIENT_CERT_IF_GIVEN,
			"REQUIRE_AND_VERIFY_CLIENT_CERT": TlsOptionSpecClientAuthClientAuthType_REQUIRE_AND_VERIFY_CLIENT_CERT,
		},
	)
}
