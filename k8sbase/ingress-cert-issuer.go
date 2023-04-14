package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/certmanagerio"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type CertIssuerProps struct {
	Enabled bool   `yaml:"enabled"`
	Email   string `yaml:"email"`
	Solver  string `yaml:"solver"` // dns or http
}

func letsEncryptIssuer(chart constructs.Construct, props CertIssuerProps, name string, server string) {
	certmanagerio.NewClusterIssuer(chart, jsii.String(name), &certmanagerio.ClusterIssuerProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String(name),
		},
		Spec: &certmanagerio.ClusterIssuerSpec{
			Acme: &certmanagerio.ClusterIssuerSpecAcme{
				Email:  jsii.String(props.Email),
				Server: jsii.String(server),
				PrivateKeySecretRef: &certmanagerio.ClusterIssuerSpecAcmePrivateKeySecretRef{
					Name: jsii.String(name),
				},
				Solvers: Ptr(MergeLists(
					Ternary(props.Solver == "http", []*certmanagerio.ClusterIssuerSpecAcmeSolvers{
						{
							Http01: &certmanagerio.ClusterIssuerSpecAcmeSolversHttp01{
								Ingress: &certmanagerio.ClusterIssuerSpecAcmeSolversHttp01Ingress{
									Class: jsii.String("traefik"),
								},
							},
						},
					}, nil),
					Ternary(props.Solver == "dns", []*certmanagerio.ClusterIssuerSpecAcmeSolvers{
						{
							Dns01: &certmanagerio.ClusterIssuerSpecAcmeSolversDns01{
								Cloudflare: &certmanagerio.ClusterIssuerSpecAcmeSolversDns01Cloudflare{
									Email: jsii.String(props.Email),
									ApiTokenSecretRef: &certmanagerio.ClusterIssuerSpecAcmeSolversDns01CloudflareApiTokenSecretRef{
										Name: jsii.String("cloudflare-api-token"),
										Key:  jsii.String("api-token"),
									},
								},
							},
						},
					}, nil),
				)),
			},
		},
	})
}

func NewCertIssuer(scope constructs.Construct, props CertIssuerProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: jsii.String("cert-manager"),
	}
	chart := cdk8s.NewChart(scope, jsii.String("cert-issuer"), &cprops)

	// NewNamespace(chart, jsii.String("namespace"), &NamespaceProps{Name: "cert-manager"})
	letsEncryptIssuer(chart, props, "letsencrypt-prod", "https://acme-v02.api.letsencrypt.org/directory")
	letsEncryptIssuer(chart, props, "letsencrypt-staging", "https://acme-staging-v02.api.letsencrypt.org/directory")

	if props.Solver == "dns" {
		NewExternalSecret(chart, jsii.String("cloudflare-externalsecret"), &ExternalSecretProps{
			Name:            jsii.String("cloudflare-api-token"),
			RefreshInterval: jsii.String("2m"),
			Secrets: map[string]string{
				"api-token": "CLOUDFLARE_API_TOKEN",
			},
		})
	}
	return chart
}
