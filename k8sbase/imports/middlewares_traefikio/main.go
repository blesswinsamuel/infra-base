// middlewares_traefikio
package middlewares_traefikio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"middlewares_traefikio.Middleware",
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
		"middlewares_traefikio.MiddlewareProps",
		reflect.TypeOf((*MiddlewareProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpec",
		reflect.TypeOf((*MiddlewareSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecAddPrefix",
		reflect.TypeOf((*MiddlewareSpecAddPrefix)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecBasicAuth",
		reflect.TypeOf((*MiddlewareSpecBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecBuffering",
		reflect.TypeOf((*MiddlewareSpecBuffering)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecChain",
		reflect.TypeOf((*MiddlewareSpecChain)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecChainMiddlewares",
		reflect.TypeOf((*MiddlewareSpecChainMiddlewares)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecCircuitBreaker",
		reflect.TypeOf((*MiddlewareSpecCircuitBreaker)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"middlewares_traefikio.MiddlewareSpecCircuitBreakerCheckPeriod",
		reflect.TypeOf((*MiddlewareSpecCircuitBreakerCheckPeriod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecCircuitBreakerCheckPeriod{}
		},
	)
	_jsii_.RegisterClass(
		"middlewares_traefikio.MiddlewareSpecCircuitBreakerFallbackDuration",
		reflect.TypeOf((*MiddlewareSpecCircuitBreakerFallbackDuration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecCircuitBreakerFallbackDuration{}
		},
	)
	_jsii_.RegisterClass(
		"middlewares_traefikio.MiddlewareSpecCircuitBreakerRecoveryDuration",
		reflect.TypeOf((*MiddlewareSpecCircuitBreakerRecoveryDuration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecCircuitBreakerRecoveryDuration{}
		},
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecCompress",
		reflect.TypeOf((*MiddlewareSpecCompress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecContentType",
		reflect.TypeOf((*MiddlewareSpecContentType)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecDigestAuth",
		reflect.TypeOf((*MiddlewareSpecDigestAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecErrors",
		reflect.TypeOf((*MiddlewareSpecErrors)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecErrorsService",
		reflect.TypeOf((*MiddlewareSpecErrorsService)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"middlewares_traefikio.MiddlewareSpecErrorsServiceKind",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceKind)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE": MiddlewareSpecErrorsServiceKind_SERVICE,
			"TRAEFIK_SERVICE": MiddlewareSpecErrorsServiceKind_TRAEFIK_SERVICE,
		},
	)
	_jsii_.RegisterClass(
		"middlewares_traefikio.MiddlewareSpecErrorsServicePort",
		reflect.TypeOf((*MiddlewareSpecErrorsServicePort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecErrorsServicePort{}
		},
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecErrorsServiceResponseForwarding",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceResponseForwarding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecErrorsServiceSticky",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecErrorsServiceStickyCookie",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceStickyCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecForwardAuth",
		reflect.TypeOf((*MiddlewareSpecForwardAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecForwardAuthTls",
		reflect.TypeOf((*MiddlewareSpecForwardAuthTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecHeaders",
		reflect.TypeOf((*MiddlewareSpecHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecInFlightReq",
		reflect.TypeOf((*MiddlewareSpecInFlightReq)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecInFlightReqSourceCriterion",
		reflect.TypeOf((*MiddlewareSpecInFlightReqSourceCriterion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecInFlightReqSourceCriterionIpStrategy",
		reflect.TypeOf((*MiddlewareSpecInFlightReqSourceCriterionIpStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecIpWhiteList",
		reflect.TypeOf((*MiddlewareSpecIpWhiteList)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecIpWhiteListIpStrategy",
		reflect.TypeOf((*MiddlewareSpecIpWhiteListIpStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecPassTlsClientCert",
		reflect.TypeOf((*MiddlewareSpecPassTlsClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecPassTlsClientCertInfo",
		reflect.TypeOf((*MiddlewareSpecPassTlsClientCertInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecPassTlsClientCertInfoIssuer",
		reflect.TypeOf((*MiddlewareSpecPassTlsClientCertInfoIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecPassTlsClientCertInfoSubject",
		reflect.TypeOf((*MiddlewareSpecPassTlsClientCertInfoSubject)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecRateLimit",
		reflect.TypeOf((*MiddlewareSpecRateLimit)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"middlewares_traefikio.MiddlewareSpecRateLimitPeriod",
		reflect.TypeOf((*MiddlewareSpecRateLimitPeriod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecRateLimitPeriod{}
		},
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecRateLimitSourceCriterion",
		reflect.TypeOf((*MiddlewareSpecRateLimitSourceCriterion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecRateLimitSourceCriterionIpStrategy",
		reflect.TypeOf((*MiddlewareSpecRateLimitSourceCriterionIpStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecRedirectRegex",
		reflect.TypeOf((*MiddlewareSpecRedirectRegex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecRedirectScheme",
		reflect.TypeOf((*MiddlewareSpecRedirectScheme)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecReplacePath",
		reflect.TypeOf((*MiddlewareSpecReplacePath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecReplacePathRegex",
		reflect.TypeOf((*MiddlewareSpecReplacePathRegex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecRetry",
		reflect.TypeOf((*MiddlewareSpecRetry)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"middlewares_traefikio.MiddlewareSpecRetryInitialInterval",
		reflect.TypeOf((*MiddlewareSpecRetryInitialInterval)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecRetryInitialInterval{}
		},
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecStripPrefix",
		reflect.TypeOf((*MiddlewareSpecStripPrefix)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"middlewares_traefikio.MiddlewareSpecStripPrefixRegex",
		reflect.TypeOf((*MiddlewareSpecStripPrefixRegex)(nil)).Elem(),
	)
}
