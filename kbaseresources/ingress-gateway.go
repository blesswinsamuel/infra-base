package kbaseresources

import (
	"fmt"
	"strings"

	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func init() {
	RegisterModule("gateway", &Gateway{})
}

type Gateway struct {
	Domains []string `json:"domains"`
}

func (props *Gateway) Render(scope kgen.Scope) {
	var listeners []gatewayv1.Listener
	cleanDomainName := func(domain string) string {
		cleaned := strings.ReplaceAll(domain, ".", "-")
		cleaned = strings.ReplaceAll(cleaned, "*", "wildcard")
		return cleaned
	}
	for _, domain := range props.Domains {
		listenerName := cleanDomainName(domain)
		listeners = append(listeners, gatewayv1.Listener{
			Name:     gatewayv1.SectionName(fmt.Sprintf("https-%s", listenerName)),
			Port:     gatewayv1.PortNumber(443),
			Protocol: gatewayv1.HTTPSProtocolType,
			Hostname: ptr.To(gatewayv1.Hostname(domain)),
			TLS: &gatewayv1.GatewayTLSConfig{CertificateRefs: []gatewayv1.SecretObjectReference{
				{Name: gatewayv1.ObjectName(fmt.Sprintf("%s-tls", listenerName)), Namespace: ptr.To(gatewayv1.Namespace(scope.Namespace()))}},
			},
			AllowedRoutes: &gatewayv1.AllowedRoutes{Namespaces: &gatewayv1.RouteNamespaces{From: ptr.To(gatewayv1.NamespacesFromAll)}},
		})
		listeners = append(listeners, gatewayv1.Listener{
			Name:          gatewayv1.SectionName(fmt.Sprintf("http-%s", listenerName)),
			Port:          gatewayv1.PortNumber(80),
			Protocol:      gatewayv1.HTTPProtocolType,
			Hostname:      ptr.To(gatewayv1.Hostname(domain)),
			AllowedRoutes: &gatewayv1.AllowedRoutes{Namespaces: &gatewayv1.RouteNamespaces{From: ptr.To(gatewayv1.NamespacesFromAll)}},
		})
	}
	scope.AddApiObject(&gatewayv1.Gateway{
		ObjectMeta: metav1.ObjectMeta{Name: scope.ID()},
		Spec: gatewayv1.GatewaySpec{
			GatewayClassName: "traefik",
			Listeners:        listeners,
		},
	})
	for _, domain := range props.Domains {
		k8sapp.NewCertificate(scope, &k8sapp.CertificateProps{
			Name:       cleanDomainName(domain),
			Hosts:      []string{domain},
			CertIssuer: k8sapp.CertIssuerRefProps{Name: "letsencrypt-prod", Kind: "ClusterIssuer"},
		})
	}
}
