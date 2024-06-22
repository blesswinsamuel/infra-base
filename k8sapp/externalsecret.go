package k8sapp

import (
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"
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
	SecretLabels    map[string]string
	SecretStore     ExternalSecretStoreProps
}

func NewExternalSecret(scope kubegogen.Scope, props *ExternalSecretProps) kubegogen.ApiObject {
	var data []externalsecretsv1beta1.ExternalSecretData
	globals := GetGlobals(scope)
	for k, v := range props.RemoteRefs {
		var remoteRef externalsecretsv1beta1.ExternalSecretDataRemoteRef
		var ref string = v
		vParts := strings.Split(v, "/")
		switch globals.SecretsProvider {
		case "1password":
			var clusternamespace string
			if len(vParts) != 2 {
				panic("Invalid 1Password remote ref: " + v)
			}
			switch vParts[0] {
			case "cluster":
				clusternamespace = globals.ClusterName
			case "commons":
				clusternamespace = "commons"
			default:
				panic("Invalid 1Password remote ref: " + v)
			}
			ref = vParts[1]
			remoteRef = externalsecretsv1beta1.ExternalSecretDataRemoteRef{Key: "Kubernetes " + clusternamespace, Property: ref}
		case "doppler":
			if len(vParts) == 1 {
				ref = vParts[0]
			} else if len(vParts) == 2 {
				ref = vParts[1]
			} else {
				panic("Invalid Doppler remote ref: " + v)
			}
			remoteRef = externalsecretsv1beta1.ExternalSecretDataRemoteRef{Key: ref}
		}
		data = append(data, externalsecretsv1beta1.ExternalSecretData{
			SecretKey: k,
			RemoteRef: remoteRef,
		})
	}
	slices.SortFunc(data, func(a externalsecretsv1beta1.ExternalSecretData, b externalsecretsv1beta1.ExternalSecretData) int {
		return strings.Compare(a.SecretKey, b.SecretKey)
	})
	externalsecret := externalsecretsv1beta1.ExternalSecret{
		ObjectMeta: metav1.ObjectMeta{Name: props.Name}, // , Namespace: infrahelpers.StrPtrIfNonEmpty(props.Namespace)
		Spec: externalsecretsv1beta1.ExternalSecretSpec{
			RefreshInterval: infrahelpers.ToDuration(infrahelpers.UseOrDefault(props.RefreshInterval, globals.Defaults.ExternalSecretRefreshInterval)),
			SecretStoreRef: externalsecretsv1beta1.SecretStoreRef{
				Name: infrahelpers.UseOrDefault(props.SecretStore.Name, globals.Defaults.SecretStoreName),
				Kind: infrahelpers.UseOrDefault(props.SecretStore.Kind, globals.Defaults.SecretStoreKind),
			},
			Data: data,
		},
	}
	if len(props.Template) > 0 {
		externalsecret.Spec.Target.Template = &externalsecretsv1beta1.ExternalSecretTemplate{
			Type: corev1.SecretType(props.SecretType),
			Metadata: externalsecretsv1beta1.ExternalSecretTemplateMetadata{
				Labels: props.SecretLabels,
			},
			Data: props.Template,
		}
	}

	return scope.AddApiObject(&externalsecret)
}
