package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/ingressroute_traefikio"
)

type GlobalProps struct {
	Domain                         string `yaml:"domain"`
	CertIssuer                     string `yaml:"clusterCertIssuerName"`
	ClusterExternalSecretStoreName string `yaml:"clusterExternalSecretStoreName"`
	InternetAuthType               string `yaml:"internetAuthType"`
}

func SetGlobalContext(scope constructs.Construct, props GlobalProps) {
	scope.Node().SetContext(jsii.String("global"), ToYamlString(props))
}

func GetDomain(scope constructs.Construct) string {
	return GetGlobal(scope).Domain
}

func GetGlobal(scope constructs.Construct) GlobalProps {
	globalValues := scope.Node().TryGetContext(jsii.String("global")).(string)
	return FromYamlString[GlobalProps](globalValues)
}

func GetCertIssuerAnnotation(scope constructs.Construct) map[string]string {
	return map[string]string{"cert-manager.io/cluster-issuer": GetCertIssuer(scope)}
}

func GetTraefikAuthMiddlewareName(scope constructs.Construct) string {
	switch GetGlobal(scope).InternetAuthType {
	case "traefik-forward-auth":
		return "auth-traefik-forward-auth@kubernetescrd"
	case "authelia":
		return "auth-forwardauth-authelia@kubernetescrd"
	}
	panic("Invalid internetAuthType")
}

func GetTraefikCRAuthMiddleware(scope constructs.Construct) *ingressroute_traefikio.IngressRouteSpecRoutesMiddlewares {
	switch GetGlobal(scope).InternetAuthType {
	case "traefik-forward-auth":
		return &ingressroute_traefikio.IngressRouteSpecRoutesMiddlewares{Name: jsii.String("traefik-forward-auth"), Namespace: jsii.String("auth")}
	case "authelia":
		return &ingressroute_traefikio.IngressRouteSpecRoutesMiddlewares{Name: jsii.String("forwardauth-authelia"), Namespace: jsii.String("auth")}
	}
	panic("Invalid internetAuthType")
}

func JSIIMap[K comparable, V any](m map[K]V) *map[K]*V {
	out := make(map[K]*V)
	for k, v := range m {
		v := v
		out[k] = &v
	}
	return &out
}

func GetCertIssuer(scope constructs.Construct) string {
	return GetGlobal(scope).CertIssuer
}
