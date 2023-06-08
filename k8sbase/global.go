package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/ingressroute_traefikio"
	"github.com/blesswinsamuel/infra-base/k8sbase/infraglobal"
)

func SetGlobalContext(scope constructs.Construct, props infraglobal.GlobalProps) {
	scope.Node().SetContext(jsii.String("global"), infrahelpers.ToYamlString(props))
}

func GetDomain(scope constructs.Construct) string {
	return infraglobal.GetGlobal(scope).Domain
}

func GetTraefikAuthMiddlewareName(scope constructs.Construct) string {
	switch infraglobal.GetGlobal(scope).InternetAuthType {
	case "traefik-forward-auth":
		return "auth-traefik-forward-auth@kubernetescrd"
	case "authelia":
		return "auth-forwardauth-authelia@kubernetescrd"
	}
	panic("Invalid internetAuthType")
}

func GetTraefikCRAuthMiddleware(scope constructs.Construct) *ingressroute_traefikio.IngressRouteSpecRoutesMiddlewares {
	switch infraglobal.GetGlobal(scope).InternetAuthType {
	case "traefik-forward-auth":
		return &ingressroute_traefikio.IngressRouteSpecRoutesMiddlewares{Name: jsii.String("traefik-forward-auth"), Namespace: jsii.String("auth")}
	case "authelia":
		return &ingressroute_traefikio.IngressRouteSpecRoutesMiddlewares{Name: jsii.String("forwardauth-authelia"), Namespace: jsii.String("auth")}
	}
	panic("Invalid internetAuthType")
}
