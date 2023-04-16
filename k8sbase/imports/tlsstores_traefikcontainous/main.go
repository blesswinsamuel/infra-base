package tlsstores_traefikcontainous

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"tlsstores_traefikcontainous.TlsStore",
		reflect.TypeOf((*TlsStore)(nil)).Elem(),
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
			j := jsiiProxy_TlsStore{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"tlsstores_traefikcontainous.TlsStoreProps",
		reflect.TypeOf((*TlsStoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"tlsstores_traefikcontainous.TlsStoreSpec",
		reflect.TypeOf((*TlsStoreSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"tlsstores_traefikcontainous.TlsStoreSpecCertificates",
		reflect.TypeOf((*TlsStoreSpecCertificates)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"tlsstores_traefikcontainous.TlsStoreSpecDefaultCertificate",
		reflect.TypeOf((*TlsStoreSpecDefaultCertificate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"tlsstores_traefikcontainous.TlsStoreSpecDefaultGeneratedCert",
		reflect.TypeOf((*TlsStoreSpecDefaultGeneratedCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"tlsstores_traefikcontainous.TlsStoreSpecDefaultGeneratedCertDomain",
		reflect.TypeOf((*TlsStoreSpecDefaultGeneratedCertDomain)(nil)).Elem(),
	)
}
