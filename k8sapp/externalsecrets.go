package k8sapp

import (
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/packager"
	externalsecretsv1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	"golang.org/x/exp/slices"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func NewExternalSecret(scope packager.Construct, id string, props *ExternalSecretProps) packager.ApiObject {
	// construct := constructs.NewConstruct(scope, id)
	var data []externalsecretsv1beta1.ExternalSecretData
	for k, v := range props.RemoteRefs {
		data = append(data, externalsecretsv1beta1.ExternalSecretData{
			SecretKey: k,
			RemoteRef: externalsecretsv1beta1.ExternalSecretDataRemoteRef{Key: v},
		})
	}
	slices.SortFunc(data, func(a externalsecretsv1beta1.ExternalSecretData, b externalsecretsv1beta1.ExternalSecretData) int {
		return strings.Compare(a.SecretKey, b.SecretKey)
	})
	globals := GetGlobalContext(scope)
	externalsecret := externalsecretsv1beta1.ExternalSecret{
		ObjectMeta: metav1.ObjectMeta{Name: props.Name}, // , Namespace: infrahelpers.StrPtrIfNonEmpty(props.Namespace)
		Spec: externalsecretsv1beta1.ExternalSecretSpec{
			RefreshInterval: infrahelpers.ToDuration(infrahelpers.UseOrDefault(props.RefreshInterval, globals.DefaultExternalSecretRefreshInterval)),
			SecretStoreRef: externalsecretsv1beta1.SecretStoreRef{
				Name: infrahelpers.UseOrDefault(props.SecretStore.Name, globals.DefaultSecretStoreName),
				Kind: infrahelpers.UseOrDefault(props.SecretStore.Kind, globals.DefaultSecretStoreKind),
			},
			Data: data,
		},
	}
	if len(props.Template) > 0 {
		externalsecret.Spec.Target.Template = &externalsecretsv1beta1.ExternalSecretTemplate{
			Type: corev1.SecretType(props.SecretType),
			Data: props.Template,
		}
	}

	return NewK8sObject(scope, id, &externalsecret)
}
