package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/helpers"
)

type IngressProps struct {
	CertManager CertManagerProps `yaml:"certManager"`
	CertIssuer  CertIssuerProps  `yaml:"certIssuer"`
	Traefik     TraefikProps     `yaml:"traefik"`
}

func NewIngress(scope constructs.Construct, props IngressProps) constructs.Construct {
	defer logModuleTiming("ingress")()
	construct := constructs.NewConstruct(scope, jsii.String("ingress"))

	helpers.NewNamespace(construct, "ingress")
	NewCertManager(construct, props.CertManager)
	NewCertIssuer(construct, props.CertIssuer)
	NewTraefik(construct, props.Traefik)

	return construct
}
