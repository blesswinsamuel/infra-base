package k8sapp

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/packager"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewNamespaceChart(scope packager.Construct, namespaceName string) packager.Chart {
	chart := packager.NewChart(scope, namespaceName, &packager.ChartProps{
		Namespace: namespaceName,
	})
	SetNamespaceContext(chart, namespaceName)
	NewNamespace(chart, jsii.String("namespace"), namespaceName)
	return chart
}

func NewNamespace(scope packager.Construct, id *string, namespaceName string) packager.Construct {
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
	scope.Node().SetContext("namespace", namespaceName)
	return nil
}

func GetNamespaceContextPtr(scope packager.Construct) *string {
	return jsii.String(GetNamespaceContext(scope))
}

func GetNamespaceContext(scope packager.Construct) string {
	ns, _ := scope.Node().TryGetContext("namespace").(string)
	return ns
}
