package k8sapp

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	certmanagermetav1 "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CertificateProps struct {
	Name       string
	Hosts      []string
	CertIssuer CertIssuerRefProps
}

func NewCertificate(scope kubegogen.Construct, id string, props *CertificateProps) kubegogen.Construct {
	globals := GetGlobalContext(scope)
	return NewK8sObject(scope, id, &certmanagerv1.Certificate{
		ObjectMeta: metav1.ObjectMeta{Name: props.Name},
		Spec: certmanagerv1.CertificateSpec{
			DNSNames:   props.Hosts,
			SecretName: props.Name + "-tls",
			IssuerRef: certmanagermetav1.ObjectReference{
				Name: infrahelpers.UseOrDefault(props.CertIssuer.Name, globals.DefaultCertIssuerName),
				Kind: infrahelpers.UseOrDefault(props.CertIssuer.Kind, globals.DefaultCertIssuerKind),
			},
		},
	})
}
