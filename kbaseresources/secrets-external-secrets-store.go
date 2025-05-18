package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	externalsecretsv1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1"
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
	var secretStoreSpec externalsecretsv1.SecretStoreSpec
	switch k8sapp.GetGlobals(scope).ExternalSecret.SecretsProvider {
	case "doppler":
		secretStoreSpec = externalsecretsv1.SecretStoreSpec{
			Controller:      "",
			RefreshInterval: 0,
			Provider: &externalsecretsv1.SecretStoreProvider{
				Doppler: &externalsecretsv1.DopplerProvider{
					Auth: &externalsecretsv1.DopplerAuth{
						SecretRef: externalsecretsv1.DopplerAuthSecretRef{
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
		secretStoreSpec = externalsecretsv1.SecretStoreSpec{
			Controller:      "",
			RefreshInterval: 0,
			Provider: &externalsecretsv1.SecretStoreProvider{
				OnePassword: &externalsecretsv1.OnePasswordProvider{
					ConnectHost: props.OnePassword.ConnectHost,
					Vaults:      props.OnePassword.Vaults,
					Auth: &externalsecretsv1.OnePasswordAuth{
						SecretRef: &externalsecretsv1.OnePasswordAuthSecretRef{
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
	switch k8sapp.GetGlobals(scope).ExternalSecret.SecretStoreKind {
	case "ClusterSecretStore":
		scope.AddApiObject(&externalsecretsv1.ClusterSecretStore{
			ObjectMeta: metav1.ObjectMeta{Name: k8sapp.GetGlobals(scope).ExternalSecret.SecretStoreName},
			Spec:       secretStoreSpec,
		})
	case "SecretStore":
		scope.AddApiObject(&externalsecretsv1.SecretStore{
			ObjectMeta: metav1.ObjectMeta{Name: k8sapp.GetGlobals(scope).ExternalSecret.SecretStoreName},
			Spec:       secretStoreSpec,
		})
	}
}
