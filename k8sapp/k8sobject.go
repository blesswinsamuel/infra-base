package k8sapp

import (
	"github.com/blesswinsamuel/infra-base/packager"
	"k8s.io/apimachinery/pkg/runtime"
)

func NewK8sObject(scope packager.Construct, id *string, obj runtime.Object) packager.ApiObject {
	return scope.ApiObject(*id, obj)
}
