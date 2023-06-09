// generatorsexternal-secretsio
package generatorsexternalsecretsio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"generatorsexternal-secretsio.AcrAccessToken",
		reflect.TypeOf((*AcrAccessToken)(nil)).Elem(),
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
			j := jsiiProxy_AcrAccessToken{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.AcrAccessTokenProps",
		reflect.TypeOf((*AcrAccessTokenProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.AcrAccessTokenSpec",
		reflect.TypeOf((*AcrAccessTokenSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.AcrAccessTokenSpecAuth",
		reflect.TypeOf((*AcrAccessTokenSpecAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.AcrAccessTokenSpecAuthManagedIdentity",
		reflect.TypeOf((*AcrAccessTokenSpecAuthManagedIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.AcrAccessTokenSpecAuthServicePrincipal",
		reflect.TypeOf((*AcrAccessTokenSpecAuthServicePrincipal)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.AcrAccessTokenSpecAuthServicePrincipalSecretRef",
		reflect.TypeOf((*AcrAccessTokenSpecAuthServicePrincipalSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.AcrAccessTokenSpecAuthServicePrincipalSecretRefClientId",
		reflect.TypeOf((*AcrAccessTokenSpecAuthServicePrincipalSecretRefClientId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.AcrAccessTokenSpecAuthServicePrincipalSecretRefClientSecret",
		reflect.TypeOf((*AcrAccessTokenSpecAuthServicePrincipalSecretRefClientSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.AcrAccessTokenSpecAuthWorkloadIdentity",
		reflect.TypeOf((*AcrAccessTokenSpecAuthWorkloadIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.AcrAccessTokenSpecAuthWorkloadIdentityServiceAccountRef",
		reflect.TypeOf((*AcrAccessTokenSpecAuthWorkloadIdentityServiceAccountRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"generatorsexternal-secretsio.AcrAccessTokenSpecEnvironmentType",
		reflect.TypeOf((*AcrAccessTokenSpecEnvironmentType)(nil)).Elem(),
		map[string]interface{}{
			"PUBLIC_CLOUD": AcrAccessTokenSpecEnvironmentType_PUBLIC_CLOUD,
			"US_GOVERNMENT_CLOUD": AcrAccessTokenSpecEnvironmentType_US_GOVERNMENT_CLOUD,
			"CHINA_CLOUD": AcrAccessTokenSpecEnvironmentType_CHINA_CLOUD,
			"GERMAN_CLOUD": AcrAccessTokenSpecEnvironmentType_GERMAN_CLOUD,
		},
	)
	_jsii_.RegisterClass(
		"generatorsexternal-secretsio.EcrAuthorizationToken",
		reflect.TypeOf((*EcrAuthorizationToken)(nil)).Elem(),
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
			j := jsiiProxy_EcrAuthorizationToken{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.EcrAuthorizationTokenProps",
		reflect.TypeOf((*EcrAuthorizationTokenProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.EcrAuthorizationTokenSpec",
		reflect.TypeOf((*EcrAuthorizationTokenSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.EcrAuthorizationTokenSpecAuth",
		reflect.TypeOf((*EcrAuthorizationTokenSpecAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.EcrAuthorizationTokenSpecAuthJwt",
		reflect.TypeOf((*EcrAuthorizationTokenSpecAuthJwt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.EcrAuthorizationTokenSpecAuthJwtServiceAccountRef",
		reflect.TypeOf((*EcrAuthorizationTokenSpecAuthJwtServiceAccountRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.EcrAuthorizationTokenSpecAuthSecretRef",
		reflect.TypeOf((*EcrAuthorizationTokenSpecAuthSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.EcrAuthorizationTokenSpecAuthSecretRefAccessKeyIdSecretRef",
		reflect.TypeOf((*EcrAuthorizationTokenSpecAuthSecretRefAccessKeyIdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.EcrAuthorizationTokenSpecAuthSecretRefSecretAccessKeySecretRef",
		reflect.TypeOf((*EcrAuthorizationTokenSpecAuthSecretRefSecretAccessKeySecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.EcrAuthorizationTokenSpecAuthSecretRefSessionTokenSecretRef",
		reflect.TypeOf((*EcrAuthorizationTokenSpecAuthSecretRefSessionTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"generatorsexternal-secretsio.Fake",
		reflect.TypeOf((*Fake)(nil)).Elem(),
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
			j := jsiiProxy_Fake{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.FakeProps",
		reflect.TypeOf((*FakeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.FakeSpec",
		reflect.TypeOf((*FakeSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"generatorsexternal-secretsio.GcrAccessToken",
		reflect.TypeOf((*GcrAccessToken)(nil)).Elem(),
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
			j := jsiiProxy_GcrAccessToken{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.GcrAccessTokenProps",
		reflect.TypeOf((*GcrAccessTokenProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.GcrAccessTokenSpec",
		reflect.TypeOf((*GcrAccessTokenSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.GcrAccessTokenSpecAuth",
		reflect.TypeOf((*GcrAccessTokenSpecAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.GcrAccessTokenSpecAuthSecretRef",
		reflect.TypeOf((*GcrAccessTokenSpecAuthSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.GcrAccessTokenSpecAuthSecretRefSecretAccessKeySecretRef",
		reflect.TypeOf((*GcrAccessTokenSpecAuthSecretRefSecretAccessKeySecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.GcrAccessTokenSpecAuthWorkloadIdentity",
		reflect.TypeOf((*GcrAccessTokenSpecAuthWorkloadIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.GcrAccessTokenSpecAuthWorkloadIdentityServiceAccountRef",
		reflect.TypeOf((*GcrAccessTokenSpecAuthWorkloadIdentityServiceAccountRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"generatorsexternal-secretsio.Password",
		reflect.TypeOf((*Password)(nil)).Elem(),
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
			j := jsiiProxy_Password{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.PasswordProps",
		reflect.TypeOf((*PasswordProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.PasswordSpec",
		reflect.TypeOf((*PasswordSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"generatorsexternal-secretsio.VaultDynamicSecret",
		reflect.TypeOf((*VaultDynamicSecret)(nil)).Elem(),
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
			j := jsiiProxy_VaultDynamicSecret{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretProps",
		reflect.TypeOf((*VaultDynamicSecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpec",
		reflect.TypeOf((*VaultDynamicSecretSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProvider",
		reflect.TypeOf((*VaultDynamicSecretSpecProvider)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuth",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthAppRole",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthAppRole)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthAppRoleSecretRef",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthAppRoleSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthCert",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthCertClientCert",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthCertClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthCertSecretRef",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthCertSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthJwt",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthJwt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthJwtKubernetesServiceAccountToken",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthJwtKubernetesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthJwtKubernetesServiceAccountTokenServiceAccountRef",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthJwtKubernetesServiceAccountTokenServiceAccountRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthJwtSecretRef",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthJwtSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthKubernetes",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthKubernetes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthKubernetesSecretRef",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthKubernetesSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthKubernetesServiceAccountRef",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthKubernetesServiceAccountRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthLdap",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthLdap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthLdapSecretRef",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthLdapSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderAuthTokenSecretRef",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderAuthTokenSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderCaProvider",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderCaProvider)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderCaProviderType",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderCaProviderType)(nil)).Elem(),
		map[string]interface{}{
			"SECRET": VaultDynamicSecretSpecProviderCaProviderType_SECRET,
			"CONFIG_MAP": VaultDynamicSecretSpecProviderCaProviderType_CONFIG_MAP,
		},
	)
	_jsii_.RegisterEnum(
		"generatorsexternal-secretsio.VaultDynamicSecretSpecProviderVersion",
		reflect.TypeOf((*VaultDynamicSecretSpecProviderVersion)(nil)).Elem(),
		map[string]interface{}{
			"V1": VaultDynamicSecretSpecProviderVersion_V1,
			"V2": VaultDynamicSecretSpecProviderVersion_V2,
		},
	)
}
