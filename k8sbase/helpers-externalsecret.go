package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/externalsecretsio"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"golang.org/x/exp/slices"
)

const clusterExternalSecretStoreName = "secretstore"

type ExternalSecretProps struct {
	Name            *string
	Namespace       *string
	RefreshInterval *string
	Secrets         map[string]string
	Template        *externalsecretsio.ExternalSecretV1Beta1SpecTargetTemplate
}

func NewExternalSecret(scope constructs.Construct, id *string, props *ExternalSecretProps) constructs.Construct {
	// construct := constructs.NewConstruct(scope, id)
	var data []*externalsecretsio.ExternalSecretV1Beta1SpecData
	for k, v := range props.Secrets {
		data = append(data, &externalsecretsio.ExternalSecretV1Beta1SpecData{
			SecretKey: jsii.String(k),
			RemoteRef: &externalsecretsio.ExternalSecretV1Beta1SpecDataRemoteRef{Key: jsii.String(v)},
		})
	}
	slices.SortFunc(data, func(a *externalsecretsio.ExternalSecretV1Beta1SpecData, b *externalsecretsio.ExternalSecretV1Beta1SpecData) bool {
		return *a.SecretKey < *b.SecretKey
	})
	return externalsecretsio.NewExternalSecretV1Beta1(scope, id, &externalsecretsio.ExternalSecretV1Beta1Props{
		Metadata: &cdk8s.ApiObjectMetadata{Name: props.Name, Namespace: props.Namespace},
		Spec: &externalsecretsio.ExternalSecretV1Beta1Spec{
			RefreshInterval: props.RefreshInterval,
			SecretStoreRef: &externalsecretsio.ExternalSecretV1Beta1SpecSecretStoreRef{
				Name: jsii.String(clusterExternalSecretStoreName),
				Kind: jsii.String("ClusterSecretStore"),
			},
			Target: &externalsecretsio.ExternalSecretV1Beta1SpecTarget{
				Template: props.Template,
			},
			Data: &data,
		},
	})
}
