package k8sapp

import (
	"github.com/blesswinsamuel/infra-base/kubegogen"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewNamespaceChart(scope kubegogen.Scope, namespaceName string) kubegogen.Scope {
	chart := scope.CreateScope(namespaceName, kubegogen.ScopeProps{
		Namespace: namespaceName,
	})
	SetNamespaceContext(chart, namespaceName)
	if namespaceName != "default" {
		NewNamespace(chart, "namespace", namespaceName)
	}
	return chart
}

func NewNamespace(scope kubegogen.Scope, id string, namespaceName string) {
	scope.AddApiObject(&corev1.Namespace{
		ObjectMeta: v1.ObjectMeta{
			Name: namespaceName,
			Labels: map[string]string{
				"name": namespaceName,
			},
		},
	})
}

func SetNamespaceContext(scope kubegogen.Scope, namespaceName string) kubegogen.Scope {
	scope.SetContext("namespace", namespaceName)
	return nil
}

func GetNamespaceContext(scope kubegogen.Scope) string {
	ns, _ := scope.GetContext("namespace").(string)
	return ns
}
