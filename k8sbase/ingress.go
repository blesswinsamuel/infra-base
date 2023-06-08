package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type IngressProps struct {
	CertManager CertManagerProps `yaml:"certManager"`
	CertIssuer  CertIssuerProps  `yaml:"certIssuer"`
	Traefik     TraefikProps     `yaml:"traefik"`
}

func NewIngress(scope constructs.Construct, props IngressProps) constructs.Construct {
	defer logModuleTiming("ingress")()

	chart := k8sapp.NewNamespaceChart(scope, "ingress")
	NewCertManager(chart, props.CertManager)
	NewCertIssuer(chart, props.CertIssuer)
	NewTraefik(chart, props.Traefik)

	return chart
}
