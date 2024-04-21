package k8sapp

import (
	"github.com/blesswinsamuel/infra-base/kubegogen"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewNamespace(scope kubegogen.Scope, namespaceName string) {
	scope.AddApiObject(&corev1.Namespace{
		ObjectMeta: v1.ObjectMeta{
			Name: namespaceName,
			Labels: map[string]string{
				"name": namespaceName,
			},
		},
	})
}
