package k8sapp

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func NewK8sObject(scope constructs.Construct, id *string, obj runtime.Object) cdk8s.ApiObject {
	groupVersionKinds, _, err := infrahelpers.Scheme.ObjectKinds(obj)
	if err != nil {
		panic(err)
	}
	if len(groupVersionKinds) != 1 {
		panic(fmt.Errorf("expected 1 groupVersionKind, got %d: %v", len(groupVersionKinds), groupVersionKinds))
	}
	var metadata *cdk8s.ApiObjectMetadata
	if obj, ok := obj.(metav1.Object); ok {
		if obj.GetNamespace() == "" {
			if namespaceCtx := GetNamespaceContext(scope); namespaceCtx != "" {
				obj.SetNamespace(namespaceCtx)
			}
		}
		metadata = &cdk8s.ApiObjectMetadata{
			Name:        jsii.String(obj.GetName()),
			Namespace:   infrahelpers.PtrIfNonEmpty(obj.GetNamespace()),
			Labels:      infrahelpers.PtrMap(obj.GetLabels()),
			Annotations: infrahelpers.PtrMap(obj.GetAnnotations()),
		}
	}
	groupVersion := groupVersionKinds[0]
	apiobj := cdk8s.NewApiObject(scope, id, &cdk8s.ApiObjectProps{
		ApiVersion: jsii.String(groupVersion.GroupVersion().String()),
		Kind:       jsii.String(groupVersion.Kind),
		Metadata:   metadata,
	})
	mobj := infrahelpers.K8sObjectToMap(obj)
	for _, field := range infrahelpers.MapKeys(mobj) {
		if field == "apiVersion" || field == "kind" || field == "metadata" {
			continue
		}
		v := mobj[field]
		if v != nil {
			apiobj.AddJsonPatch(cdk8s.JsonPatch_Replace(jsii.String("/"+field), v))
		}
	}
	return apiobj
}
