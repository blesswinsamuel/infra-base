package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/kgen"
	externalsecretsv1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	externalsecretsmetav1 "github.com/external-secrets/external-secrets/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ExternalSecretsStore struct {
	TokenSecretName      string `json:"tokenSecretName"`
	TokenSecretNamespace string `json:"tokenSecretNamespace"`
	TokenSecretKey       string `json:"tokenSecretKey"`

	OnePassword struct {
		ConnectHost string         `json:"connectHost"`
		Vaults      map[string]int `json:"vaults"`
	} `json:"onePassword"`
}

// https://external-secrets.io/v0.5.8/provider-kubernetes/
// https://external-secrets.io/v0.5.8/spec/

func (props *ExternalSecretsStore) Render(scope kgen.Scope) {
	var secretStoreSpec externalsecretsv1beta1.SecretStoreSpec
	switch GetGlobals(scope).ExternalSecret.SecretsProvider {
	case "doppler":
		secretStoreSpec = externalsecretsv1beta1.SecretStoreSpec{
			Controller:      "",
			RefreshInterval: 0,
			Provider: &externalsecretsv1beta1.SecretStoreProvider{
				Doppler: &externalsecretsv1beta1.DopplerProvider{
					Auth: &externalsecretsv1beta1.DopplerAuth{
						SecretRef: externalsecretsv1beta1.DopplerAuthSecretRef{
							DopplerToken: externalsecretsmetav1.SecretKeySelector{
								Name:      props.TokenSecretName,
								Namespace: infrahelpers.PtrIfNonEmpty(props.TokenSecretNamespace),
								Key:       props.TokenSecretKey,
							},
						},
					},
				},
			},
		}
	case "1password":
		secretStoreSpec = externalsecretsv1beta1.SecretStoreSpec{
			Controller:      "",
			RefreshInterval: 0,
			Provider: &externalsecretsv1beta1.SecretStoreProvider{
				OnePassword: &externalsecretsv1beta1.OnePasswordProvider{
					ConnectHost: props.OnePassword.ConnectHost,
					Vaults:      props.OnePassword.Vaults,
					Auth: &externalsecretsv1beta1.OnePasswordAuth{
						SecretRef: &externalsecretsv1beta1.OnePasswordAuthSecretRef{
							ConnectToken: externalsecretsmetav1.SecretKeySelector{
								Name:      props.TokenSecretName,
								Namespace: infrahelpers.Ptr(props.TokenSecretNamespace),
								Key:       props.TokenSecretKey,
							},
						},
					},
				},
			},
		}
	}
	switch GetGlobals(scope).ExternalSecret.SecretStoreKind {
	case "ClusterSecretStore":
		scope.AddApiObject(&externalsecretsv1beta1.ClusterSecretStore{
			ObjectMeta: metav1.ObjectMeta{Name: GetGlobals(scope).ExternalSecret.SecretStoreName},
			Spec:       secretStoreSpec,
		})
	case "SecretStore":
		scope.AddApiObject(&externalsecretsv1beta1.SecretStore{
			ObjectMeta: metav1.ObjectMeta{Name: GetGlobals(scope).ExternalSecret.SecretStoreName},
			Spec:       secretStoreSpec,
		})
	}
}
