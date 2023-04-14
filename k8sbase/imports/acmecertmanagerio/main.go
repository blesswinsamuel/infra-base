package acmecertmanagerio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"acmecert-managerio.Challenge",
		reflect.TypeOf((*Challenge)(nil)).Elem(),
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
			j := jsiiProxy_Challenge{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeProps",
		reflect.TypeOf((*ChallengeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpec",
		reflect.TypeOf((*ChallengeSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecIssuerRef",
		reflect.TypeOf((*ChallengeSpecIssuerRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolver",
		reflect.TypeOf((*ChallengeSpecSolver)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01",
		reflect.TypeOf((*ChallengeSpecSolverDns01)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01AcmeDns",
		reflect.TypeOf((*ChallengeSpecSolverDns01AcmeDns)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01AcmeDnsAccountSecretRef",
		reflect.TypeOf((*ChallengeSpecSolverDns01AcmeDnsAccountSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01Akamai",
		reflect.TypeOf((*ChallengeSpecSolverDns01Akamai)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01AkamaiAccessTokenSecretRef",
		reflect.TypeOf((*ChallengeSpecSolverDns01AkamaiAccessTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01AkamaiClientSecretSecretRef",
		reflect.TypeOf((*ChallengeSpecSolverDns01AkamaiClientSecretSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01AkamaiClientTokenSecretRef",
		reflect.TypeOf((*ChallengeSpecSolverDns01AkamaiClientTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01AzureDns",
		reflect.TypeOf((*ChallengeSpecSolverDns01AzureDns)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01AzureDnsClientSecretSecretRef",
		reflect.TypeOf((*ChallengeSpecSolverDns01AzureDnsClientSecretSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"acmecert-managerio.ChallengeSpecSolverDns01AzureDnsEnvironment",
		reflect.TypeOf((*ChallengeSpecSolverDns01AzureDnsEnvironment)(nil)).Elem(),
		map[string]interface{}{
			"AZURE_PUBLIC_CLOUD": ChallengeSpecSolverDns01AzureDnsEnvironment_AZURE_PUBLIC_CLOUD,
			"AZURE_CHINA_CLOUD": ChallengeSpecSolverDns01AzureDnsEnvironment_AZURE_CHINA_CLOUD,
			"AZURE_GERMAN_CLOUD": ChallengeSpecSolverDns01AzureDnsEnvironment_AZURE_GERMAN_CLOUD,
			"AZURE_US_GOVERNMENT_CLOUD": ChallengeSpecSolverDns01AzureDnsEnvironment_AZURE_US_GOVERNMENT_CLOUD,
		},
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01AzureDnsManagedIdentity",
		reflect.TypeOf((*ChallengeSpecSolverDns01AzureDnsManagedIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01CloudDns",
		reflect.TypeOf((*ChallengeSpecSolverDns01CloudDns)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01CloudDnsServiceAccountSecretRef",
		reflect.TypeOf((*ChallengeSpecSolverDns01CloudDnsServiceAccountSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01Cloudflare",
		reflect.TypeOf((*ChallengeSpecSolverDns01Cloudflare)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01CloudflareApiKeySecretRef",
		reflect.TypeOf((*ChallengeSpecSolverDns01CloudflareApiKeySecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01CloudflareApiTokenSecretRef",
		reflect.TypeOf((*ChallengeSpecSolverDns01CloudflareApiTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"acmecert-managerio.ChallengeSpecSolverDns01CnameStrategy",
		reflect.TypeOf((*ChallengeSpecSolverDns01CnameStrategy)(nil)).Elem(),
		map[string]interface{}{
			"NONE": ChallengeSpecSolverDns01CnameStrategy_NONE,
			"FOLLOW": ChallengeSpecSolverDns01CnameStrategy_FOLLOW,
		},
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01Digitalocean",
		reflect.TypeOf((*ChallengeSpecSolverDns01Digitalocean)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01DigitaloceanTokenSecretRef",
		reflect.TypeOf((*ChallengeSpecSolverDns01DigitaloceanTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01Rfc2136",
		reflect.TypeOf((*ChallengeSpecSolverDns01Rfc2136)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01Rfc2136TsigSecretSecretRef",
		reflect.TypeOf((*ChallengeSpecSolverDns01Rfc2136TsigSecretSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01Route53",
		reflect.TypeOf((*ChallengeSpecSolverDns01Route53)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01Route53AccessKeyIdSecretRef",
		reflect.TypeOf((*ChallengeSpecSolverDns01Route53AccessKeyIdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01Route53SecretAccessKeySecretRef",
		reflect.TypeOf((*ChallengeSpecSolverDns01Route53SecretAccessKeySecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverDns01Webhook",
		reflect.TypeOf((*ChallengeSpecSolverDns01Webhook)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01",
		reflect.TypeOf((*ChallengeSpecSolverHttp01)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01GatewayHttpRoute",
		reflect.TypeOf((*ChallengeSpecSolverHttp01GatewayHttpRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01GatewayHttpRouteParentRefs",
		reflect.TypeOf((*ChallengeSpecSolverHttp01GatewayHttpRouteParentRefs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01Ingress",
		reflect.TypeOf((*ChallengeSpecSolverHttp01Ingress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressIngressTemplate",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressIngressTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressIngressTemplateMetadata",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressIngressTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplate",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateMetadata",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpec",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinity",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinity",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinity",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinity",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverHttp01IngressPodTemplateSpecTolerations",
		reflect.TypeOf((*ChallengeSpecSolverHttp01IngressPodTemplateSpecTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.ChallengeSpecSolverSelector",
		reflect.TypeOf((*ChallengeSpecSolverSelector)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"acmecert-managerio.ChallengeSpecType",
		reflect.TypeOf((*ChallengeSpecType)(nil)).Elem(),
		map[string]interface{}{
			"HTTP_01": ChallengeSpecType_HTTP_01,
			"DNS_01": ChallengeSpecType_DNS_01,
		},
	)
	_jsii_.RegisterClass(
		"acmecert-managerio.Order",
		reflect.TypeOf((*Order)(nil)).Elem(),
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
			j := jsiiProxy_Order{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.OrderProps",
		reflect.TypeOf((*OrderProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.OrderSpec",
		reflect.TypeOf((*OrderSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"acmecert-managerio.OrderSpecIssuerRef",
		reflect.TypeOf((*OrderSpecIssuerRef)(nil)).Elem(),
	)
}
