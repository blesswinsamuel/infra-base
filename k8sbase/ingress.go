package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type IngressProps struct {
	CertManager CertManagerProps `json:"certManager"`
	CertIssuer  CertIssuerProps  `json:"certIssuer"`
	Traefik     TraefikProps     `json:"traefik"`
}

func NewIngress(scope constructs.Construct, props IngressProps) constructs.Construct {
	defer logModuleTiming("ingress")()

	{
		chart := k8sapp.NewNamespaceChart(scope, "cert-manager")
		NewCertManager(chart, props.CertManager)
		NewCertIssuer(chart, props.CertIssuer)
	}

	{
		chart := k8sapp.NewNamespaceChart(scope, "ingress")
		NewTraefik(chart, props.Traefik)
	}

	return scope
}
