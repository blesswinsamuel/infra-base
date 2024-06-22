package k8sapp

import (
	"fmt"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	traefiktypes "github.com/traefik/traefik/v3/pkg/types"
	networkingv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NameNamespace struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type IngressHostPath struct {
	Path            string `json:"path"`
	ServiceName     string `json:"serviceName"`
	ServicePortName string `json:"servicePortName"`
}

type IngressHost struct {
	Host  string            `json:"host"`
	Paths []IngressHostPath `json:"paths"`
	Tls   bool              `json:"tls"`
}

type IngressProps struct {
	Name                   string             `json:"name"`
	TraefikMiddlewareNames []NameNamespace    `json:"traefikMiddlewares"`
	Hosts                  []IngressHost      `json:"hosts"`
	IngressType            string             `json:"ingressType"`
	CertIssuer             CertIssuerRefProps `json:"certIssuer"`
}

type CertIssuerRefProps struct {
	Name string
	Kind string
}

func NewIngress(scope kubegogen.Scope, props *IngressProps) kubegogen.Scope {
	if props.IngressType == "" {
		props.IngressType = "kubernetes"
	}
	globals := GetGlobals(scope)
	if globals.Ingress.DisableTls {
		for i, host := range props.Hosts {
			host.Tls = false
			props.Hosts[i] = host
		}
	}
	if props.IngressType == "traefik" {
		ingressRules := []traefikv1alpha1.Route{}
		tlsHosts := map[string]bool{}
		for _, host := range props.Hosts {
			hostPaths := []traefikv1alpha1.Service{}
			if host.Tls {
				tlsHosts[host.Host] = true
			}
			for _, path := range host.Paths {
				pathStr := path.Path
				if pathStr == "" {
					pathStr = "/"
				}
				var kind string
				if strings.Contains(path.ServiceName, "@") {
					kind = "TraefikService"
				} else {
					kind = "Service"
				}
				hostPaths = append(hostPaths, traefikv1alpha1.Service{LoadBalancerSpec: traefikv1alpha1.LoadBalancerSpec{
					Name: path.ServiceName,
					Kind: kind,
				}})
			}
			middlewares := []traefikv1alpha1.MiddlewareRef{}
			for _, middleware := range props.TraefikMiddlewareNames {
				middlewares = append(middlewares, traefikv1alpha1.MiddlewareRef{
					Name:      middleware.Name,
					Namespace: middleware.Namespace,
				})
			}
			ingressRules = append(ingressRules, traefikv1alpha1.Route{
				Match:       "Host(`" + host.Host + "`)",
				Kind:        "Rule",
				Services:    hostPaths,
				Middlewares: middlewares,
			})
		}
		tlsDomains := []traefiktypes.Domain{}
		if len(tlsHosts) > 0 {
			NewCertificate(scope, &CertificateProps{
				Name:       props.Name,
				Hosts:      infrahelpers.MapKeys(tlsHosts),
				CertIssuer: props.CertIssuer,
			})
		}
		for _, host := range infrahelpers.MapKeys(tlsHosts) {
			tlsDomains = append(tlsDomains, traefiktypes.Domain{
				Main: host,
			})
		}
		scope.AddApiObject(&traefikv1alpha1.IngressRoute{
			ObjectMeta: v1.ObjectMeta{
				Name: props.Name,
			},
			Spec: traefikv1alpha1.IngressRouteSpec{
				EntryPoints: infrahelpers.If(globals.Ingress.DisableTls, []string{"web"}, []string{"websecure"}),
				Routes:      ingressRules,
				TLS: infrahelpers.If(len(tlsDomains) > 0, &traefikv1alpha1.TLS{
					SecretName: fmt.Sprintf("%s-tls", props.Name),
					Domains:    tlsDomains,
				}, nil),
			},
		})
	} else if props.IngressType == "kubernetes" {
		ingressRules := []networkingv1.IngressRule{}
		tlsHosts := map[string]bool{}
		for _, host := range props.Hosts {
			hostPaths := []networkingv1.HTTPIngressPath{}
			tlsHosts[host.Host] = true
			for _, path := range host.Paths {
				pathStr := path.Path
				if pathStr == "" {
					pathStr = "/"
				}
				hostPaths = append(hostPaths, networkingv1.HTTPIngressPath{
					Path:     pathStr,
					PathType: infrahelpers.Ptr(networkingv1.PathType("Prefix")),
					Backend: networkingv1.IngressBackend{
						Service: &networkingv1.IngressServiceBackend{
							Name: path.ServiceName,
							Port: networkingv1.ServiceBackendPort{
								Name: path.ServicePortName,
							},
						},
					},
				})
			}
			ingressRules = append(ingressRules, networkingv1.IngressRule{
				Host: host.Host,
				IngressRuleValue: networkingv1.IngressRuleValue{
					HTTP: &networkingv1.HTTPIngressRuleValue{
						Paths: hostPaths,
					},
				},
			})
		}
		traefikMiddlwareNames := []string{}
		for _, traefikMiddleware := range props.TraefikMiddlewareNames {
			traefikMiddlwareNames = append(traefikMiddlwareNames, fmt.Sprintf("%s-%s@kubernetescrd", traefikMiddleware.Namespace, traefikMiddleware.Name))
		}
		annotations := map[string]string{}
		clusterIssuerAnnotationKey := map[string]string{
			"ClusterIssuer": "cert-manager.io/cluster-issuer",
			"Issuer":        "cert-manager.io/issuer",
		}[infrahelpers.UseOrDefault(props.CertIssuer.Kind, globals.Cert.CertIssuerKind)]
		annotations[clusterIssuerAnnotationKey] = infrahelpers.UseOrDefault(props.CertIssuer.Name, globals.Cert.CertIssuerName)
		if len(traefikMiddlwareNames) > 0 {
			annotations["traefik.ingress.kubernetes.io/router.middlewares"] = strings.Join(traefikMiddlwareNames, ",")
		}
		scope.AddApiObject(&networkingv1.Ingress{
			ObjectMeta: v1.ObjectMeta{
				Name:        props.Name,
				Annotations: annotations,
			},
			Spec: networkingv1.IngressSpec{
				Rules: ingressRules,
				TLS: []networkingv1.IngressTLS{
					{
						Hosts:      infrahelpers.MapKeys(tlsHosts),
						SecretName: fmt.Sprintf("%s-tls", props.Name),
					},
				},
			},
		})
	} else {
		panic("Invalid ingressType")
	}

	return scope
}

type HomepageIngressAnnotationsProps struct {
	Group       string
	Name        string
	Icon        string
	Description string
}

func HomepageIngressAnnotations(props HomepageIngressAnnotationsProps) map[string]string {
	return map[string]string{
		"gethomepage.dev/description": props.Description,
		"gethomepage.dev/enabled":     "true",
		"gethomepage.dev/group":       props.Group,
		"gethomepage.dev/icon":        props.Icon,
		"gethomepage.dev/name":        props.Name,
	}
}
