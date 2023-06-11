package k8sbase

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
	certmanageracmev1 "github.com/cert-manager/cert-manager/pkg/apis/acme/v1"
	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	certmanagermetav1 "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CertIssuerProps struct {
	Enabled bool   `json:"enabled"`
	Email   string `json:"email"`
	Solver  string `json:"solver"` // dns or http
}

func letsEncryptIssuer(chart packager.Construct, props CertIssuerProps, name string, server string) {
	issuer := &certmanagerv1.ClusterIssuer{
		ObjectMeta: v1.ObjectMeta{Name: name},
		Spec: certmanagerv1.IssuerSpec{IssuerConfig: certmanagerv1.IssuerConfig{
			ACME: &certmanageracmev1.ACMEIssuer{
				Email:  props.Email,
				Server: server,
				PrivateKey: certmanagermetav1.SecretKeySelector{
					LocalObjectReference: certmanagermetav1.LocalObjectReference{Name: name},
				},
			},
		}},
	}
	if props.Solver == "http" {
		issuer.Spec.IssuerConfig.ACME.Solvers = []certmanageracmev1.ACMEChallengeSolver{
			{
				HTTP01: &certmanageracmev1.ACMEChallengeSolverHTTP01{
					Ingress: &certmanageracmev1.ACMEChallengeSolverHTTP01Ingress{
						Class: jsii.String("traefik"),
					},
				},
			},
		}
	}
	if props.Solver == "dns" {
		issuer.Spec.IssuerConfig.ACME.Solvers = []certmanageracmev1.ACMEChallengeSolver{
			{
				DNS01: &certmanageracmev1.ACMEChallengeSolverDNS01{
					Cloudflare: &certmanageracmev1.ACMEIssuerDNS01ProviderCloudflare{
						Email: props.Email,
						APIToken: &certmanagermetav1.SecretKeySelector{
							LocalObjectReference: certmanagermetav1.LocalObjectReference{Name: "cloudflare-api-token"},
							Key:                  "api-token",
						},
					},
				},
			},
		}
	}

	k8sapp.NewK8sObject(chart, jsii.String(name), issuer)
}

func NewCertIssuer(scope packager.Construct, props CertIssuerProps) packager.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := &packager.ChartProps{
		Namespace: "cert-manager",
	}
	chart := packager.NewChart(scope, "cert-issuer", cprops)

	// NewNamespace(chart, jsii.String("namespace"), &NamespaceProps{Name: "cert-manager"})
	letsEncryptIssuer(chart, props, "letsencrypt-prod", "https://acme-v02.api.letsencrypt.org/directory")
	letsEncryptIssuer(chart, props, "letsencrypt-staging", "https://acme-staging-v02.api.letsencrypt.org/directory")

	if props.Solver == "dns" {
		k8sapp.NewExternalSecret(chart, jsii.String("cloudflare-externalsecret"), &k8sapp.ExternalSecretProps{
			Name: "cloudflare-api-token",
			RemoteRefs: map[string]string{
				"api-token": "CLOUDFLARE_API_TOKEN",
			},
		})
	}
	return chart
}
