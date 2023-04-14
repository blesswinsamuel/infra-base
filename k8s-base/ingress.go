package resourcesbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type IngressProps struct {
	CertManager CertManagerProps `yaml:"certManager"`
	CertIssuer  CertIssuerProps  `yaml:"certIssuer"`
	Traefik     TraefikProps     `yaml:"traefik"`
}

func NewIngress(scope constructs.Construct, props IngressProps) constructs.Construct {
	construct := constructs.NewConstruct(scope, jsii.String("ingress"))

	NewNamespace(construct, "ingress")
	NewCertManager(construct, props.CertManager)
	NewCertIssuer(construct, props.CertIssuer)
	NewTraefik(construct, props.Traefik)

	return construct
}
