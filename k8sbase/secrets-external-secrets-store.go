package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	externalsecretsv1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	externalsecretsmetav1 "github.com/external-secrets/external-secrets/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ExternalSecretsStore struct {
	DopplerTokenSecretName      string `json:"dopplerTokenSecretName"`
	DopplerTokenSecretNamespace string `json:"dopplerTokenSecretNamespace"`
	DopplerTokenSecretKey       string `json:"dopplerTokenSecretKey"`
}

// https://external-secrets.io/v0.5.8/provider-kubernetes/
// https://external-secrets.io/v0.5.8/spec/

func (props *ExternalSecretsStore) Render(scope kubegogen.Scope) {
	secretStoreSpec := externalsecretsv1beta1.SecretStoreSpec{
		Controller:      "",
		RefreshInterval: 0,
		Provider: &externalsecretsv1beta1.SecretStoreProvider{
			Doppler: &externalsecretsv1beta1.DopplerProvider{
				Auth: &externalsecretsv1beta1.DopplerAuth{
					SecretRef: externalsecretsv1beta1.DopplerAuthSecretRef{
						DopplerToken: externalsecretsmetav1.SecretKeySelector{
							Name:      props.DopplerTokenSecretName,
							Namespace: infrahelpers.PtrIfNonEmpty(props.DopplerTokenSecretNamespace),
							Key:       props.DopplerTokenSecretKey,
						},
					},
				},
			},
		},
	}
	switch GetGlobals(scope).Defaults.SecretStoreKind {
	case "ClusterSecretStore":
		scope.AddApiObject(&externalsecretsv1beta1.ClusterSecretStore{
			ObjectMeta: metav1.ObjectMeta{Name: GetGlobals(scope).Defaults.SecretStoreName},
			Spec:       secretStoreSpec,
		})
	case "SecretStore":
		scope.AddApiObject(&externalsecretsv1beta1.SecretStore{
			ObjectMeta: metav1.ObjectMeta{Name: GetGlobals(scope).Defaults.SecretStoreName},
			Spec:       secretStoreSpec,
		})
	}
}
