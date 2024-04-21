package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	externalsecretsv1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	externalsecretsmetav1 "github.com/external-secrets/external-secrets/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterSecretStoreProps struct {
}

// https://external-secrets.io/v0.5.8/provider-kubernetes/
// https://external-secrets.io/v0.5.8/spec/

func (props *ClusterSecretStoreProps) Render(scope kubegogen.Scope) {
	scope.AddApiObject(&externalsecretsv1beta1.ClusterSecretStore{
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
}
