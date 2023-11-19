package k8sapp

import (
	"github.com/blesswinsamuel/infra-base/kubegogen"
	"k8s.io/apimachinery/pkg/runtime"
)

func NewK8sObject(scope kubegogen.Construct, id string, obj runtime.Object) kubegogen.ApiObject {
	return scope.ApiObject(id, obj)
}
