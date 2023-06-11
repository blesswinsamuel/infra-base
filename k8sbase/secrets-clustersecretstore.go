package k8sbase

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
	externalsecretsv1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	externalsecretsmetav1 "github.com/external-secrets/external-secrets/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterSecretStoreProps struct {
	DopplerServiceToken string `json:"dopplerServiceToken"`
}

// https://external-secrets.io/v0.5.8/provider-kubernetes/
// https://external-secrets.io/v0.5.8/spec/

func NewClusterSecretStore(scope packager.Construct, props ClusterSecretStoreProps) packager.Chart {
	cprops := &packager.ChartProps{}
	chart := packager.NewChart(scope, "cluster-secret-store", cprops)
	k8sapp.NewSecret(chart, jsii.String("secret"), &k8sapp.SecretProps{
		Name:      "doppler-token-auth-api",
		Namespace: "default",
		Data: map[string][]byte{
			"dopplerToken": []byte(props.DopplerServiceToken),
		},
	})
	k8sapp.NewK8sObject(chart, jsii.String("cluster-secret-store"), &externalsecretsv1beta1.ClusterSecretStore{
		ObjectMeta: metav1.ObjectMeta{Name: GetGlobal(scope).ClusterExternalSecretStoreName},
		Spec: externalsecretsv1beta1.SecretStoreSpec{
			Controller:      "",
			RefreshInterval: 0,
			Provider: &externalsecretsv1beta1.SecretStoreProvider{
				Doppler: &externalsecretsv1beta1.DopplerProvider{
					Auth: &externalsecretsv1beta1.DopplerAuth{
						SecretRef: externalsecretsv1beta1.DopplerAuthSecretRef{
							DopplerToken: externalsecretsmetav1.SecretKeySelector{
								Name:      "doppler-token-auth-api",
								Namespace: infrahelpers.Ptr("default"),
								Key:       "dopplerToken",
							},
						},
					},
				},
			},
		},
	})
	return chart
}
