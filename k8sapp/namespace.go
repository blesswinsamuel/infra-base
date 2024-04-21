package k8sapp

import (
	"github.com/blesswinsamuel/infra-base/kubegogen"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewNamespaceChart(scope kubegogen.Construct, namespaceName string) kubegogen.Construct {
	chart := scope.Chart(namespaceName, kubegogen.ChartProps{
		Namespace: namespaceName,
	})
	SetNamespaceContext(chart, namespaceName)
	if namespaceName != "default" {
		NewNamespace(chart, "namespace", namespaceName)
	}
	return chart
}

func NewNamespace(scope kubegogen.Construct, id string, namespaceName string) {
	scope.ApiObject(&corev1.Namespace{
		ObjectMeta: v1.ObjectMeta{
			Name: namespaceName,
			Labels: map[string]string{
				"name": namespaceName,
			},
		},
	})
}

func SetNamespaceContext(scope kubegogen.Construct, namespaceName string) kubegogen.Construct {
	scope.SetContext("namespace", namespaceName)
	return nil
}

func GetNamespaceContext(scope kubegogen.Construct) string {
	ns, _ := scope.GetContext("namespace").(string)
	return ns
}
