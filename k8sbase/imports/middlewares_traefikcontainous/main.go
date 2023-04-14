package middlewares_traefikcontainous

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"middlewares_traefikcontainous.Middleware",
		reflect.TypeOf((*Middleware)(nil)).Elem(),
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
			j := jsiiProxy_Middleware{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareProps",
		reflect.TypeOf((*MiddlewareProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpec",
		reflect.TypeOf((*MiddlewareSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecAddPrefix",
		reflect.TypeOf((*MiddlewareSpecAddPrefix)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecBasicAuth",
		reflect.TypeOf((*MiddlewareSpecBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecBuffering",
		reflect.TypeOf((*MiddlewareSpecBuffering)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecChain",
		reflect.TypeOf((*MiddlewareSpecChain)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecChainMiddlewares",
		reflect.TypeOf((*MiddlewareSpecChainMiddlewares)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecCircuitBreaker",
		reflect.TypeOf((*MiddlewareSpecCircuitBreaker)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"middlewares_traefikcontainous.MiddlewareSpecCircuitBreakerCheckPeriod",
		reflect.TypeOf((*MiddlewareSpecCircuitBreakerCheckPeriod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecCircuitBreakerCheckPeriod{}
		},
	)
	_jsii_.RegisterClass(
		"middlewares_traefikcontainous.MiddlewareSpecCircuitBreakerFallbackDuration",
		reflect.TypeOf((*MiddlewareSpecCircuitBreakerFallbackDuration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecCircuitBreakerFallbackDuration{}
		},
	)
	_jsii_.RegisterClass(
		"middlewares_traefikcontainous.MiddlewareSpecCircuitBreakerRecoveryDuration",
		reflect.TypeOf((*MiddlewareSpecCircuitBreakerRecoveryDuration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecCircuitBreakerRecoveryDuration{}
		},
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecCompress",
		reflect.TypeOf((*MiddlewareSpecCompress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecContentType",
		reflect.TypeOf((*MiddlewareSpecContentType)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecDigestAuth",
		reflect.TypeOf((*MiddlewareSpecDigestAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecErrors",
		reflect.TypeOf((*MiddlewareSpecErrors)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecErrorsService",
		reflect.TypeOf((*MiddlewareSpecErrorsService)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"middlewares_traefikcontainous.MiddlewareSpecErrorsServiceKind",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceKind)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE": MiddlewareSpecErrorsServiceKind_SERVICE,
			"TRAEFIK_SERVICE": MiddlewareSpecErrorsServiceKind_TRAEFIK_SERVICE,
		},
	)
	_jsii_.RegisterClass(
		"middlewares_traefikcontainous.MiddlewareSpecErrorsServicePort",
		reflect.TypeOf((*MiddlewareSpecErrorsServicePort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecErrorsServicePort{}
		},
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecErrorsServiceResponseForwarding",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceResponseForwarding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecErrorsServiceSticky",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecErrorsServiceStickyCookie",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceStickyCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecForwardAuth",
		reflect.TypeOf((*MiddlewareSpecForwardAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecForwardAuthTls",
		reflect.TypeOf((*MiddlewareSpecForwardAuthTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecHeaders",
		reflect.TypeOf((*MiddlewareSpecHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecInFlightReq",
		reflect.TypeOf((*MiddlewareSpecInFlightReq)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecInFlightReqSourceCriterion",
		reflect.TypeOf((*MiddlewareSpecInFlightReqSourceCriterion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecInFlightReqSourceCriterionIpStrategy",
		reflect.TypeOf((*MiddlewareSpecInFlightReqSourceCriterionIpStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecIpWhiteList",
		reflect.TypeOf((*MiddlewareSpecIpWhiteList)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecIpWhiteListIpStrategy",
		reflect.TypeOf((*MiddlewareSpecIpWhiteListIpStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecPassTlsClientCert",
		reflect.TypeOf((*MiddlewareSpecPassTlsClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecPassTlsClientCertInfo",
		reflect.TypeOf((*MiddlewareSpecPassTlsClientCertInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecPassTlsClientCertInfoIssuer",
		reflect.TypeOf((*MiddlewareSpecPassTlsClientCertInfoIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecPassTlsClientCertInfoSubject",
		reflect.TypeOf((*MiddlewareSpecPassTlsClientCertInfoSubject)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecRateLimit",
		reflect.TypeOf((*MiddlewareSpecRateLimit)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"middlewares_traefikcontainous.MiddlewareSpecRateLimitPeriod",
		reflect.TypeOf((*MiddlewareSpecRateLimitPeriod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecRateLimitPeriod{}
		},
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecRateLimitSourceCriterion",
		reflect.TypeOf((*MiddlewareSpecRateLimitSourceCriterion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecRateLimitSourceCriterionIpStrategy",
		reflect.TypeOf((*MiddlewareSpecRateLimitSourceCriterionIpStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecRedirectRegex",
		reflect.TypeOf((*MiddlewareSpecRedirectRegex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecRedirectScheme",
		reflect.TypeOf((*MiddlewareSpecRedirectScheme)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecReplacePath",
		reflect.TypeOf((*MiddlewareSpecReplacePath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecReplacePathRegex",
		reflect.TypeOf((*MiddlewareSpecReplacePathRegex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecRetry",
		reflect.TypeOf((*MiddlewareSpecRetry)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"middlewares_traefikcontainous.MiddlewareSpecRetryInitialInterval",
		reflect.TypeOf((*MiddlewareSpecRetryInitialInterval)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecRetryInitialInterval{}
		},
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecStripPrefix",
		reflect.TypeOf((*MiddlewareSpecStripPrefix)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikcontainous.MiddlewareSpecStripPrefixRegex",
		reflect.TypeOf((*MiddlewareSpecStripPrefixRegex)(nil)).Elem(),
	)
}
