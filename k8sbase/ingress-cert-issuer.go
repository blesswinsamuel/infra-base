package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
	certmanageracmev1 "github.com/cert-manager/cert-manager/pkg/apis/acme/v1"
	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	certmanagermetav1 "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CertIssuerProps struct {
	Email  string `json:"email"`
	Solver string `json:"solver"` // dns or http
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
						Class: infrahelpers.Ptr("traefik"),
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

	k8sapp.NewK8sObject(chart, name, issuer)
}

func (props *CertIssuerProps) Chart(scope packager.Construct) packager.Construct {
	cprops := packager.ChartProps{
		Namespace: "cert-manager",
	}
	chart := scope.Chart("cert-issuer", cprops)

	// NewNamespace(chart, ("namespace"), &NamespaceProps{Name: "cert-manager"})
	letsEncryptIssuer(chart, *props, "letsencrypt-prod", "https://acme-v02.api.letsencrypt.org/directory")
	letsEncryptIssuer(chart, *props, "letsencrypt-staging", "https://acme-staging-v02.api.letsencrypt.org/directory")

	if props.Solver == "dns" {
		k8sapp.NewExternalSecret(chart, "cloudflare-externalsecret", &k8sapp.ExternalSecretProps{
			Name: "cloudflare-api-token",
			RemoteRefs: map[string]string{
				"api-token": "CLOUDFLARE_API_TOKEN",
			},
		})
	}
	return chart
}
