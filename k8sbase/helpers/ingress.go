package helpers

import (
	"fmt"
	"strings"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/certmanagerio"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/ingressroute_traefikio"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/k8s"
	"github.com/blesswinsamuel/infra-base/k8sbase/infraglobal"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type NameNamespace struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

type Path struct {
	Path            string `yaml:"path"`
	ServiceName     string `yaml:"serviceName"`
	ServicePortName string `yaml:"servicePortName"`
}

type Host struct {
	Host  string `yaml:"host"`
	Paths []Path `yaml:"paths"`
	Tls   bool   `yaml:"tls"`
}

type IngressProps struct {
	Name                   string          `yaml:"name"`
	TraefikMiddlewareNames []NameNamespace `yaml:"traefikMiddlewares"`
	Hosts                  []Host          `yaml:"hosts"`
	IngressType            string          `yaml:"ingressType"`
}

func NewIngress(scope constructs.Construct, id *string, props *IngressProps) constructs.Construct {
	if props.IngressType == "" {
		props.IngressType = "kubernetes"
	}
	if props.IngressType == "traefik" {
		ingressRules := []*ingressroute_traefikio.IngressRouteSpecRoutes{}
		tlsHosts := map[string]bool{}
		for _, host := range props.Hosts {
			hostPaths := []*ingressroute_traefikio.IngressRouteSpecRoutesServices{}
			if host.Tls {
				tlsHosts[host.Host] = true
			}
			for _, path := range host.Paths {
				pathStr := path.Path
				if pathStr == "" {
					pathStr = "/"
				}
				var kind ingressroute_traefikio.IngressRouteSpecRoutesServicesKind
				if strings.Contains(path.ServiceName, "@") {
					kind = ingressroute_traefikio.IngressRouteSpecRoutesServicesKind_TRAEFIK_SERVICE
				} else {
					kind = ingressroute_traefikio.IngressRouteSpecRoutesServicesKind_SERVICE
				}
				hostPaths = append(hostPaths, &ingressroute_traefikio.IngressRouteSpecRoutesServices{
					Name: jsii.String(path.ServiceName),
					Kind: kind,
				})
			}
			middlewares := []*ingressroute_traefikio.IngressRouteSpecRoutesMiddlewares{}
			for _, middleware := range props.TraefikMiddlewareNames {
				middlewares = append(middlewares, &ingressroute_traefikio.IngressRouteSpecRoutesMiddlewares{
					Name:      jsii.String(middleware.Name),
					Namespace: jsii.String(middleware.Namespace),
				})
			}
			ingressRules = append(ingressRules, &ingressroute_traefikio.IngressRouteSpecRoutes{
				Match:       jsii.String("Host(`" + host.Host + "`)"),
				Kind:        ingressroute_traefikio.IngressRouteSpecRoutesKind_RULE,
				Services:    &hostPaths,
				Middlewares: &middlewares,
			})
		}
		tlsDomains := []*ingressroute_traefikio.IngressRouteSpecTlsDomains{}
		if len(tlsHosts) > 0 {
			certmanagerio.NewCertificate(scope, jsii.String(*id+"-cert"), &certmanagerio.CertificateProps{
				Metadata: &cdk8s.ApiObjectMetadata{
					Name: jsii.String(props.Name),
				},
				Spec: &certmanagerio.CertificateSpec{
					DnsNames:   infrahelpers.PtrSlice(infrahelpers.MapKeys(tlsHosts)...),
					SecretName: jsii.String(fmt.Sprintf("%s-tls", props.Name)),
					IssuerRef: &certmanagerio.CertificateSpecIssuerRef{
						Name: jsii.String(infraglobal.GetCertIssuer(scope)),
						Kind: jsii.String("ClusterIssuer"),
					},
				},
			})
		}
		for _, host := range infrahelpers.MapKeys(tlsHosts) {
			tlsDomains = append(tlsDomains, &ingressroute_traefikio.IngressRouteSpecTlsDomains{
				Main: jsii.String(host),
			})
		}
		ingressroute_traefikio.NewIngressRoute(scope, id, &ingressroute_traefikio.IngressRouteProps{
			Metadata: &cdk8s.ApiObjectMetadata{
				Name: jsii.String(props.Name),
			},
			Spec: &ingressroute_traefikio.IngressRouteSpec{
				EntryPoints: jsii.PtrSlice("websecure"),
				Routes:      &ingressRules,
				Tls: &ingressroute_traefikio.IngressRouteSpecTls{
					SecretName: jsii.String(fmt.Sprintf("%s-tls", props.Name)),
					Domains:    &tlsDomains,
				},
			},
		})
	} else if props.IngressType == "kubernetes" {
		ingressRules := []*k8s.IngressRule{}
		tlsHosts := map[string]bool{}
		for _, host := range props.Hosts {
			hostPaths := []*k8s.HttpIngressPath{}
			tlsHosts[host.Host] = true
			for _, path := range host.Paths {
				pathStr := path.Path
				if pathStr == "" {
					pathStr = "/"
				}
				hostPaths = append(hostPaths, &k8s.HttpIngressPath{
					Path:     &pathStr,
					PathType: jsii.String("Prefix"),
					Backend: &k8s.IngressBackend{
						Service: &k8s.IngressServiceBackend{
							Name: jsii.String(path.ServiceName),
							Port: &k8s.ServiceBackendPort{
								Name: jsii.String(path.ServicePortName),
							},
						},
					},
				})
			}
			ingressRules = append(ingressRules, &k8s.IngressRule{
				Host: jsii.String(host.Host),
				Http: &k8s.HttpIngressRuleValue{
					Paths: &hostPaths,
				},
			})
		}
		traefikMiddlwareNames := []string{}
		for _, traefikMiddleware := range props.TraefikMiddlewareNames {
			traefikMiddlwareNames = append(traefikMiddlwareNames, fmt.Sprintf("%s-%s", traefikMiddleware.Name, traefikMiddleware.Namespace))
		}
		annotations := infraglobal.GetCertIssuerAnnotation(scope)
		if len(traefikMiddlwareNames) > 0 {
			annotations["traefik.ingress.kubernetes.io/router.middlewares"] = strings.Join(traefikMiddlwareNames, ",")
		}
		k8s.NewKubeIngress(scope, id, &k8s.KubeIngressProps{
			Metadata: &k8s.ObjectMeta{
				Name:        jsii.String(props.Name),
				Annotations: JSIIMap(annotations),
			},
			Spec: &k8s.IngressSpec{
				Rules: &ingressRules,
				Tls: &[]*k8s.IngressTls{
					{
						Hosts:      infrahelpers.PtrSlice(infrahelpers.MapKeys(tlsHosts)...),
						SecretName: jsii.String(fmt.Sprintf("%s-tls", props.Name)),
					},
				},
			},
		})
	} else {
		panic("Invalid ingressType")
	}

	return scope
}
