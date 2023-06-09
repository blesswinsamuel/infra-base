package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8simports/externalsecretsio"
	"github.com/blesswinsamuel/infra-base/k8simports/k8s"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type ClusterSecretStoreProps struct {
	DopplerServiceToken string `yaml:"dopplerServiceToken"`
}

// https://external-secrets.io/v0.5.8/provider-kubernetes/
// https://external-secrets.io/v0.5.8/spec/

func NewClusterSecretStore(scope constructs.Construct, props ClusterSecretStoreProps) cdk8s.Chart {
	cprops := cdk8s.ChartProps{}
	chart := cdk8s.NewChart(scope, jsii.String("cluster-secret-store"), &cprops)
	k8s.NewKubeSecret(chart, jsii.String("secret"), &k8s.KubeSecretProps{
		Metadata: &k8s.ObjectMeta{
			Name:      jsii.String("doppler-token-auth-api"),
			Namespace: jsii.String("default"),
		},
		Data: &map[string]*string{
			"dopplerToken": jsii.String(props.DopplerServiceToken),
		},
	})
	externalsecretsio.NewClusterSecretStoreV1Beta1(chart, jsii.String("cluster-secret-store"), &externalsecretsio.ClusterSecretStoreV1Beta1Props{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String(GetGlobal(scope).ClusterExternalSecretStoreName),
		},
		Spec: &externalsecretsio.ClusterSecretStoreV1Beta1Spec{
			//   # controller: doppler  # like ingressClassName definition
			Provider: &externalsecretsio.ClusterSecretStoreV1Beta1SpecProvider{
				Doppler: &externalsecretsio.ClusterSecretStoreV1Beta1SpecProviderDoppler{
					Auth: &externalsecretsio.ClusterSecretStoreV1Beta1SpecProviderDopplerAuth{
						SecretRef: &externalsecretsio.ClusterSecretStoreV1Beta1SpecProviderDopplerAuthSecretRef{
							DopplerToken: &externalsecretsio.ClusterSecretStoreV1Beta1SpecProviderDopplerAuthSecretRefDopplerToken{
								Name:      jsii.String("doppler-token-auth-api"),
								Key:       jsii.String("dopplerToken"),
								Namespace: jsii.String("default"),
							},
						},
					},
				},
			},
		},
	})
	return chart
}
