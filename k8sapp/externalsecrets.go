package k8sapp

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/externalsecretsio"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"golang.org/x/exp/slices"
)

type ExternalSecretStoreProps struct {
	Name string
	Kind string
}

type ExternalSecretProps struct {
	Name string
	// Namespace       string // optional
	RefreshInterval string // optional
	RemoteRefs      map[string]string
	Template        map[string]string
	SecretType      string
	SecretStore     ExternalSecretStoreProps
}

func NewExternalSecret(scope constructs.Construct, id *string, props *ExternalSecretProps) externalsecretsio.ExternalSecretV1Beta1 {
	// construct := constructs.NewConstruct(scope, id)
	var data []*externalsecretsio.ExternalSecretV1Beta1SpecData
	for k, v := range props.RemoteRefs {
		data = append(data, &externalsecretsio.ExternalSecretV1Beta1SpecData{
			SecretKey: jsii.String(k),
			RemoteRef: &externalsecretsio.ExternalSecretV1Beta1SpecDataRemoteRef{Key: jsii.String(v)},
		})
	}
	slices.SortFunc(data, func(a *externalsecretsio.ExternalSecretV1Beta1SpecData, b *externalsecretsio.ExternalSecretV1Beta1SpecData) bool {
		return *a.SecretKey < *b.SecretKey
	})
	globals := GetGlobalContext(scope)
	var target *externalsecretsio.ExternalSecretV1Beta1SpecTarget
	if len(props.Template) > 0 {
		target = &externalsecretsio.ExternalSecretV1Beta1SpecTarget{
			Template: &externalsecretsio.ExternalSecretV1Beta1SpecTargetTemplate{
				Type: infrahelpers.PtrIfNonEmpty(props.SecretType),
				Data: infrahelpers.PtrMap(props.Template),
			},
		}
	}
	return externalsecretsio.NewExternalSecretV1Beta1(scope, id, &externalsecretsio.ExternalSecretV1Beta1Props{
		Metadata: &cdk8s.ApiObjectMetadata{Name: &props.Name}, // , Namespace: infrahelpers.StrPtrIfNonEmpty(props.Namespace)
		Spec: &externalsecretsio.ExternalSecretV1Beta1Spec{
			RefreshInterval: jsii.String(infrahelpers.UseOrDefault(props.RefreshInterval, globals.DefaultExternalSecretRefreshInterval)),
			SecretStoreRef: &externalsecretsio.ExternalSecretV1Beta1SpecSecretStoreRef{
				Name: jsii.String(infrahelpers.UseOrDefault(props.SecretStore.Name, globals.DefaultSecretStoreName)),
				Kind: jsii.String(infrahelpers.UseOrDefault(props.SecretStore.Kind, globals.DefaultSecretStoreKind)),
			},
			Target: target,
			Data:   &data,
		},
	})
}
