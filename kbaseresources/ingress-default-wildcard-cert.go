package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DefaultWildcardCertificateProps struct {
	Name string
}

func (props *DefaultWildcardCertificateProps) Render(scope kgen.Scope) {
	globals := k8sapp.GetGlobals(scope)
	k8sapp.NewCertificate(scope, &k8sapp.CertificateProps{
		Name:  props.Name,
		Hosts: []string{"*." + globals.Domain, globals.Domain},
	})
	scope.AddApiObject(&traefikv1alpha1.TLSStore{
		ObjectMeta: metav1.ObjectMeta{Name: "default"},
		Spec: traefikv1alpha1.TLSStoreSpec{
			DefaultCertificate: &traefikv1alpha1.Certificate{
				SecretName: props.Name + "-tls",
			},
		},
	})
}
