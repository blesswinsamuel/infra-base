package k8sapp

import (
	"fmt"
	"strings"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	traefikv1alpha1 "github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	traefiktypes "github.com/traefik/traefik/v2/pkg/types"
	networkingv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NameNamespace struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

type IngressHostPath struct {
	Path            string `yaml:"path"`
	ServiceName     string `yaml:"serviceName"`
	ServicePortName string `yaml:"servicePortName"`
}

type IngressHost struct {
	Host  string            `yaml:"host"`
	Paths []IngressHostPath `yaml:"paths"`
	Tls   bool              `yaml:"tls"`
}

type IngressProps struct {
	Name                   string             `yaml:"name"`
	TraefikMiddlewareNames []NameNamespace    `yaml:"traefikMiddlewares"`
	Hosts                  []IngressHost      `yaml:"hosts"`
	IngressType            string             `yaml:"ingressType"`
	CertIssuer             CertIssuerRefProps `yaml:"certIssuer"`
}

type CertIssuerRefProps struct {
	Name string
	Kind string
}

func NewIngress(scope constructs.Construct, id *string, props *IngressProps) constructs.Construct {
	if props.IngressType == "" {
		props.IngressType = "kubernetes"
	}
	globals := GetGlobalContext(scope)
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
			NewCertificate(scope, jsii.String(*id+"-cert"), &CertificateProps{
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
		NewK8sObject(scope, id, &traefikv1alpha1.IngressRoute{
			ObjectMeta: v1.ObjectMeta{
				Name: props.Name,
			},
			Spec: traefikv1alpha1.IngressRouteSpec{
				EntryPoints: []string{"websecure"},
				Routes:      ingressRules,
				TLS: &traefikv1alpha1.TLS{
					SecretName: fmt.Sprintf("%s-tls", props.Name),
					Domains:    tlsDomains,
				},
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
		}[infrahelpers.UseOrDefault(props.CertIssuer.Kind, globals.DefaultCertIssuerKind)]
		annotations[clusterIssuerAnnotationKey] = infrahelpers.UseOrDefault(props.CertIssuer.Name, globals.DefaultCertIssuerName)
		if len(traefikMiddlwareNames) > 0 {
			annotations["traefik.ingress.kubernetes.io/router.middlewares"] = strings.Join(traefikMiddlwareNames, ",")
		}
		NewK8sObject(scope, id, &networkingv1.Ingress{
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
