package k8sapp

import (
	"github.com/blesswinsamuel/infra-base/packager"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewNamespaceChart(scope packager.Construct, namespaceName string) packager.Chart {
	chart := scope.Chart(namespaceName, packager.ChartProps{
		Namespace: namespaceName,
	})
	SetNamespaceContext(chart, namespaceName)
	NewNamespace(chart, "namespace", namespaceName)
	return chart
}

func NewNamespace(scope packager.Construct, id string, namespaceName string) packager.Construct {
	return NewK8sObject(scope, id, &corev1.Namespace{
		ObjectMeta: v1.ObjectMeta{
			Name: namespaceName,
			Labels: map[string]string{
				"name": namespaceName,
			},
		},
	})
}

func SetNamespaceContext(scope packager.Construct, namespaceName string) packager.Construct {
	scope.SetContext("namespace", namespaceName)
	return nil
}

func GetNamespaceContext(scope packager.Construct) string {
	ns, _ := scope.GetContext("namespace").(string)
	return ns
}
