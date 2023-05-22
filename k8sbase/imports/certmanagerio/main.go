// cert-managerio
package certmanagerio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cert-managerio.Certificate",
		reflect.TypeOf((*Certificate)(nil)).Elem(),
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
			j := jsiiProxy_Certificate{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateProps",
		reflect.TypeOf((*CertificateProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cert-managerio.CertificateRequest",
		reflect.TypeOf((*CertificateRequest)(nil)).Elem(),
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
			j := jsiiProxy_CertificateRequest{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateRequestProps",
		reflect.TypeOf((*CertificateRequestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateRequestSpec",
		reflect.TypeOf((*CertificateRequestSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateRequestSpecIssuerRef",
		reflect.TypeOf((*CertificateRequestSpecIssuerRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cert-managerio.CertificateRequestSpecUsages",
		reflect.TypeOf((*CertificateRequestSpecUsages)(nil)).Elem(),
		map[string]interface{}{
			"SIGNING": CertificateRequestSpecUsages_SIGNING,
			"DIGITAL_SIGNATURE": CertificateRequestSpecUsages_DIGITAL_SIGNATURE,
			"CONTENT_COMMITMENT": CertificateRequestSpecUsages_CONTENT_COMMITMENT,
			"KEY_ENCIPHERMENT": CertificateRequestSpecUsages_KEY_ENCIPHERMENT,
			"KEY_AGREEMENT": CertificateRequestSpecUsages_KEY_AGREEMENT,
			"DATA_ENCIPHERMENT": CertificateRequestSpecUsages_DATA_ENCIPHERMENT,
			"CERT_SIGN": CertificateRequestSpecUsages_CERT_SIGN,
			"CRL_SIGN": CertificateRequestSpecUsages_CRL_SIGN,
			"ENCIPHER_ONLY": CertificateRequestSpecUsages_ENCIPHER_ONLY,
			"DECIPHER_ONLY": CertificateRequestSpecUsages_DECIPHER_ONLY,
			"ANY": CertificateRequestSpecUsages_ANY,
			"SERVER_AUTH": CertificateRequestSpecUsages_SERVER_AUTH,
			"CLIENT_AUTH": CertificateRequestSpecUsages_CLIENT_AUTH,
			"CODE_SIGNING": CertificateRequestSpecUsages_CODE_SIGNING,
			"EMAIL_PROTECTION": CertificateRequestSpecUsages_EMAIL_PROTECTION,
			"S_MIME": CertificateRequestSpecUsages_S_MIME,
			"IPSEC_END_SYSTEM": CertificateRequestSpecUsages_IPSEC_END_SYSTEM,
			"IPSEC_TUNNEL": CertificateRequestSpecUsages_IPSEC_TUNNEL,
			"IPSEC_USER": CertificateRequestSpecUsages_IPSEC_USER,
			"TIMESTAMPING": CertificateRequestSpecUsages_TIMESTAMPING,
			"OCSP_SIGNING": CertificateRequestSpecUsages_OCSP_SIGNING,
			"MICROSOFT_SGC": CertificateRequestSpecUsages_MICROSOFT_SGC,
			"NETSCAPE_SGC": CertificateRequestSpecUsages_NETSCAPE_SGC,
		},
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateSpec",
		reflect.TypeOf((*CertificateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateSpecAdditionalOutputFormats",
		reflect.TypeOf((*CertificateSpecAdditionalOutputFormats)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cert-managerio.CertificateSpecAdditionalOutputFormatsType",
		reflect.TypeOf((*CertificateSpecAdditionalOutputFormatsType)(nil)).Elem(),
		map[string]interface{}{
			"DER": CertificateSpecAdditionalOutputFormatsType_DER,
			"COMBINED_PEM": CertificateSpecAdditionalOutputFormatsType_COMBINED_PEM,
		},
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateSpecIssuerRef",
		reflect.TypeOf((*CertificateSpecIssuerRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateSpecKeystores",
		reflect.TypeOf((*CertificateSpecKeystores)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateSpecKeystoresJks",
		reflect.TypeOf((*CertificateSpecKeystoresJks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateSpecKeystoresJksPasswordSecretRef",
		reflect.TypeOf((*CertificateSpecKeystoresJksPasswordSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateSpecKeystoresPkcs12",
		reflect.TypeOf((*CertificateSpecKeystoresPkcs12)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateSpecKeystoresPkcs12PasswordSecretRef",
		reflect.TypeOf((*CertificateSpecKeystoresPkcs12PasswordSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateSpecPrivateKey",
		reflect.TypeOf((*CertificateSpecPrivateKey)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cert-managerio.CertificateSpecPrivateKeyAlgorithm",
		reflect.TypeOf((*CertificateSpecPrivateKeyAlgorithm)(nil)).Elem(),
		map[string]interface{}{
			"RSA": CertificateSpecPrivateKeyAlgorithm_RSA,
			"ECDSA": CertificateSpecPrivateKeyAlgorithm_ECDSA,
			"ED25519": CertificateSpecPrivateKeyAlgorithm_ED25519,
		},
	)
	_jsii_.RegisterEnum(
		"cert-managerio.CertificateSpecPrivateKeyEncoding",
		reflect.TypeOf((*CertificateSpecPrivateKeyEncoding)(nil)).Elem(),
		map[string]interface{}{
			"PKCS1": CertificateSpecPrivateKeyEncoding_PKCS1,
			"PKCS8": CertificateSpecPrivateKeyEncoding_PKCS8,
		},
	)
	_jsii_.RegisterEnum(
		"cert-managerio.CertificateSpecPrivateKeyRotationPolicy",
		reflect.TypeOf((*CertificateSpecPrivateKeyRotationPolicy)(nil)).Elem(),
		map[string]interface{}{
			"NEVER": CertificateSpecPrivateKeyRotationPolicy_NEVER,
			"ALWAYS": CertificateSpecPrivateKeyRotationPolicy_ALWAYS,
		},
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateSpecSecretTemplate",
		reflect.TypeOf((*CertificateSpecSecretTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.CertificateSpecSubject",
		reflect.TypeOf((*CertificateSpecSubject)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cert-managerio.CertificateSpecUsages",
		reflect.TypeOf((*CertificateSpecUsages)(nil)).Elem(),
		map[string]interface{}{
			"SIGNING": CertificateSpecUsages_SIGNING,
			"DIGITAL_SIGNATURE": CertificateSpecUsages_DIGITAL_SIGNATURE,
			"CONTENT_COMMITMENT": CertificateSpecUsages_CONTENT_COMMITMENT,
			"KEY_ENCIPHERMENT": CertificateSpecUsages_KEY_ENCIPHERMENT,
			"KEY_AGREEMENT": CertificateSpecUsages_KEY_AGREEMENT,
			"DATA_ENCIPHERMENT": CertificateSpecUsages_DATA_ENCIPHERMENT,
			"CERT_SIGN": CertificateSpecUsages_CERT_SIGN,
			"CRL_SIGN": CertificateSpecUsages_CRL_SIGN,
			"ENCIPHER_ONLY": CertificateSpecUsages_ENCIPHER_ONLY,
			"DECIPHER_ONLY": CertificateSpecUsages_DECIPHER_ONLY,
			"ANY": CertificateSpecUsages_ANY,
			"SERVER_AUTH": CertificateSpecUsages_SERVER_AUTH,
			"CLIENT_AUTH": CertificateSpecUsages_CLIENT_AUTH,
			"CODE_SIGNING": CertificateSpecUsages_CODE_SIGNING,
			"EMAIL_PROTECTION": CertificateSpecUsages_EMAIL_PROTECTION,
			"S_MIME": CertificateSpecUsages_S_MIME,
			"IPSEC_END_SYSTEM": CertificateSpecUsages_IPSEC_END_SYSTEM,
			"IPSEC_TUNNEL": CertificateSpecUsages_IPSEC_TUNNEL,
			"IPSEC_USER": CertificateSpecUsages_IPSEC_USER,
			"TIMESTAMPING": CertificateSpecUsages_TIMESTAMPING,
			"OCSP_SIGNING": CertificateSpecUsages_OCSP_SIGNING,
			"MICROSOFT_SGC": CertificateSpecUsages_MICROSOFT_SGC,
			"NETSCAPE_SGC": CertificateSpecUsages_NETSCAPE_SGC,
		},
	)
	_jsii_.RegisterClass(
		"cert-managerio.ClusterIssuer",
		reflect.TypeOf((*ClusterIssuer)(nil)).Elem(),
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
			j := jsiiProxy_ClusterIssuer{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerProps",
		reflect.TypeOf((*ClusterIssuerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpec",
		reflect.TypeOf((*ClusterIssuerSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcme",
		reflect.TypeOf((*ClusterIssuerSpecAcme)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeExternalAccountBinding",
		reflect.TypeOf((*ClusterIssuerSpecAcmeExternalAccountBinding)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cert-managerio.ClusterIssuerSpecAcmeExternalAccountBindingKeyAlgorithm",
		reflect.TypeOf((*ClusterIssuerSpecAcmeExternalAccountBindingKeyAlgorithm)(nil)).Elem(),
		map[string]interface{}{
			"HS256": ClusterIssuerSpecAcmeExternalAccountBindingKeyAlgorithm_HS256,
			"HS384": ClusterIssuerSpecAcmeExternalAccountBindingKeyAlgorithm_HS384,
			"HS512": ClusterIssuerSpecAcmeExternalAccountBindingKeyAlgorithm_HS512,
		},
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeExternalAccountBindingKeySecretRef",
		reflect.TypeOf((*ClusterIssuerSpecAcmeExternalAccountBindingKeySecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmePrivateKeySecretRef",
		reflect.TypeOf((*ClusterIssuerSpecAcmePrivateKeySecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolvers",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolvers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01AcmeDns",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01AcmeDns)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01AcmeDnsAccountSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01AcmeDnsAccountSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01Akamai",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01Akamai)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01AkamaiAccessTokenSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01AkamaiAccessTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01AkamaiClientSecretSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01AkamaiClientSecretSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01AkamaiClientTokenSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01AkamaiClientTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01AzureDns",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01AzureDns)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01AzureDnsClientSecretSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01AzureDnsClientSecretSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01AzureDnsEnvironment",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01AzureDnsEnvironment)(nil)).Elem(),
		map[string]interface{}{
			"AZURE_PUBLIC_CLOUD": ClusterIssuerSpecAcmeSolversDns01AzureDnsEnvironment_AZURE_PUBLIC_CLOUD,
			"AZURE_CHINA_CLOUD": ClusterIssuerSpecAcmeSolversDns01AzureDnsEnvironment_AZURE_CHINA_CLOUD,
			"AZURE_GERMAN_CLOUD": ClusterIssuerSpecAcmeSolversDns01AzureDnsEnvironment_AZURE_GERMAN_CLOUD,
			"AZURE_US_GOVERNMENT_CLOUD": ClusterIssuerSpecAcmeSolversDns01AzureDnsEnvironment_AZURE_US_GOVERNMENT_CLOUD,
		},
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01AzureDnsManagedIdentity",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01AzureDnsManagedIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01CloudDns",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01CloudDns)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01CloudDnsServiceAccountSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01CloudDnsServiceAccountSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01Cloudflare",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01Cloudflare)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01CloudflareApiKeySecretRef",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01CloudflareApiKeySecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01CloudflareApiTokenSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01CloudflareApiTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01CnameStrategy",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01CnameStrategy)(nil)).Elem(),
		map[string]interface{}{
			"NONE": ClusterIssuerSpecAcmeSolversDns01CnameStrategy_NONE,
			"FOLLOW": ClusterIssuerSpecAcmeSolversDns01CnameStrategy_FOLLOW,
		},
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01Digitalocean",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01Digitalocean)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01DigitaloceanTokenSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01DigitaloceanTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01Rfc2136",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01Rfc2136)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01Rfc2136TsigSecretSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01Rfc2136TsigSecretSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01Route53",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01Route53)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01Route53AccessKeyIdSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01Route53AccessKeyIdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01Route53SecretAccessKeySecretRef",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01Route53SecretAccessKeySecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversDns01Webhook",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversDns01Webhook)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01GatewayHttpRoute",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01GatewayHttpRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01GatewayHttpRouteParentRefs",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01GatewayHttpRouteParentRefs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01Ingress",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01Ingress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressIngressTemplate",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressIngressTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressIngressTemplateMetadata",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressIngressTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplate",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateMetadata",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpec",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinity",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinity",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinity",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinity",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecImagePullSecrets",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecImagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecTolerations",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversHttp01IngressPodTemplateSpecTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecAcmeSolversSelector",
		reflect.TypeOf((*ClusterIssuerSpecAcmeSolversSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecCa",
		reflect.TypeOf((*ClusterIssuerSpecCa)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecSelfSigned",
		reflect.TypeOf((*ClusterIssuerSpecSelfSigned)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecVault",
		reflect.TypeOf((*ClusterIssuerSpecVault)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecVaultAuth",
		reflect.TypeOf((*ClusterIssuerSpecVaultAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecVaultAuthAppRole",
		reflect.TypeOf((*ClusterIssuerSpecVaultAuthAppRole)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecVaultAuthAppRoleSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecVaultAuthAppRoleSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecVaultAuthKubernetes",
		reflect.TypeOf((*ClusterIssuerSpecVaultAuthKubernetes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecVaultAuthKubernetesSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecVaultAuthKubernetesSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecVaultAuthKubernetesServiceAccountRef",
		reflect.TypeOf((*ClusterIssuerSpecVaultAuthKubernetesServiceAccountRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecVaultAuthTokenSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecVaultAuthTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecVaultCaBundleSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecVaultCaBundleSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecVenafi",
		reflect.TypeOf((*ClusterIssuerSpecVenafi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecVenafiCloud",
		reflect.TypeOf((*ClusterIssuerSpecVenafiCloud)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecVenafiCloudApiTokenSecretRef",
		reflect.TypeOf((*ClusterIssuerSpecVenafiCloudApiTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecVenafiTpp",
		reflect.TypeOf((*ClusterIssuerSpecVenafiTpp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.ClusterIssuerSpecVenafiTppCredentialsRef",
		reflect.TypeOf((*ClusterIssuerSpecVenafiTppCredentialsRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cert-managerio.Issuer",
		reflect.TypeOf((*Issuer)(nil)).Elem(),
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
			j := jsiiProxy_Issuer{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerProps",
		reflect.TypeOf((*IssuerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpec",
		reflect.TypeOf((*IssuerSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcme",
		reflect.TypeOf((*IssuerSpecAcme)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeExternalAccountBinding",
		reflect.TypeOf((*IssuerSpecAcmeExternalAccountBinding)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cert-managerio.IssuerSpecAcmeExternalAccountBindingKeyAlgorithm",
		reflect.TypeOf((*IssuerSpecAcmeExternalAccountBindingKeyAlgorithm)(nil)).Elem(),
		map[string]interface{}{
			"HS256": IssuerSpecAcmeExternalAccountBindingKeyAlgorithm_HS256,
			"HS384": IssuerSpecAcmeExternalAccountBindingKeyAlgorithm_HS384,
			"HS512": IssuerSpecAcmeExternalAccountBindingKeyAlgorithm_HS512,
		},
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeExternalAccountBindingKeySecretRef",
		reflect.TypeOf((*IssuerSpecAcmeExternalAccountBindingKeySecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmePrivateKeySecretRef",
		reflect.TypeOf((*IssuerSpecAcmePrivateKeySecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolvers",
		reflect.TypeOf((*IssuerSpecAcmeSolvers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01AcmeDns",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01AcmeDns)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01AcmeDnsAccountSecretRef",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01AcmeDnsAccountSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01Akamai",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01Akamai)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01AkamaiAccessTokenSecretRef",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01AkamaiAccessTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01AkamaiClientSecretSecretRef",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01AkamaiClientSecretSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01AkamaiClientTokenSecretRef",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01AkamaiClientTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01AzureDns",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01AzureDns)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01AzureDnsClientSecretSecretRef",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01AzureDnsClientSecretSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cert-managerio.IssuerSpecAcmeSolversDns01AzureDnsEnvironment",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01AzureDnsEnvironment)(nil)).Elem(),
		map[string]interface{}{
			"AZURE_PUBLIC_CLOUD": IssuerSpecAcmeSolversDns01AzureDnsEnvironment_AZURE_PUBLIC_CLOUD,
			"AZURE_CHINA_CLOUD": IssuerSpecAcmeSolversDns01AzureDnsEnvironment_AZURE_CHINA_CLOUD,
			"AZURE_GERMAN_CLOUD": IssuerSpecAcmeSolversDns01AzureDnsEnvironment_AZURE_GERMAN_CLOUD,
			"AZURE_US_GOVERNMENT_CLOUD": IssuerSpecAcmeSolversDns01AzureDnsEnvironment_AZURE_US_GOVERNMENT_CLOUD,
		},
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01AzureDnsManagedIdentity",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01AzureDnsManagedIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01CloudDns",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01CloudDns)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01CloudDnsServiceAccountSecretRef",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01CloudDnsServiceAccountSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01Cloudflare",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01Cloudflare)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01CloudflareApiKeySecretRef",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01CloudflareApiKeySecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01CloudflareApiTokenSecretRef",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01CloudflareApiTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cert-managerio.IssuerSpecAcmeSolversDns01CnameStrategy",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01CnameStrategy)(nil)).Elem(),
		map[string]interface{}{
			"NONE": IssuerSpecAcmeSolversDns01CnameStrategy_NONE,
			"FOLLOW": IssuerSpecAcmeSolversDns01CnameStrategy_FOLLOW,
		},
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01Digitalocean",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01Digitalocean)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01DigitaloceanTokenSecretRef",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01DigitaloceanTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01Rfc2136",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01Rfc2136)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01Rfc2136TsigSecretSecretRef",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01Rfc2136TsigSecretSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01Route53",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01Route53)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01Route53AccessKeyIdSecretRef",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01Route53AccessKeyIdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01Route53SecretAccessKeySecretRef",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01Route53SecretAccessKeySecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversDns01Webhook",
		reflect.TypeOf((*IssuerSpecAcmeSolversDns01Webhook)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01GatewayHttpRoute",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01GatewayHttpRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01GatewayHttpRouteParentRefs",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01GatewayHttpRouteParentRefs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01Ingress",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01Ingress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressIngressTemplate",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressIngressTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressIngressTemplateMetadata",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressIngressTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplate",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateMetadata",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpec",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinity",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinity",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinity",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinity",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecImagePullSecrets",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecImagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecTolerations",
		reflect.TypeOf((*IssuerSpecAcmeSolversHttp01IngressPodTemplateSpecTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecAcmeSolversSelector",
		reflect.TypeOf((*IssuerSpecAcmeSolversSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecCa",
		reflect.TypeOf((*IssuerSpecCa)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecSelfSigned",
		reflect.TypeOf((*IssuerSpecSelfSigned)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecVault",
		reflect.TypeOf((*IssuerSpecVault)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecVaultAuth",
		reflect.TypeOf((*IssuerSpecVaultAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecVaultAuthAppRole",
		reflect.TypeOf((*IssuerSpecVaultAuthAppRole)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecVaultAuthAppRoleSecretRef",
		reflect.TypeOf((*IssuerSpecVaultAuthAppRoleSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecVaultAuthKubernetes",
		reflect.TypeOf((*IssuerSpecVaultAuthKubernetes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecVaultAuthKubernetesSecretRef",
		reflect.TypeOf((*IssuerSpecVaultAuthKubernetesSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecVaultAuthKubernetesServiceAccountRef",
		reflect.TypeOf((*IssuerSpecVaultAuthKubernetesServiceAccountRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecVaultAuthTokenSecretRef",
		reflect.TypeOf((*IssuerSpecVaultAuthTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecVaultCaBundleSecretRef",
		reflect.TypeOf((*IssuerSpecVaultCaBundleSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecVenafi",
		reflect.TypeOf((*IssuerSpecVenafi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecVenafiCloud",
		reflect.TypeOf((*IssuerSpecVenafiCloud)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecVenafiCloudApiTokenSecretRef",
		reflect.TypeOf((*IssuerSpecVenafiCloudApiTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecVenafiTpp",
		reflect.TypeOf((*IssuerSpecVenafiTpp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cert-managerio.IssuerSpecVenafiTppCredentialsRef",
		reflect.TypeOf((*IssuerSpecVenafiTppCredentialsRef)(nil)).Elem(),
	)
}
