package packager

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type ApiObject interface {
	Construct
	// Children() bool
}

type cdk8sApiObject struct {
	cdk8sConstruct
	construct cdk8s.ApiObject
}

func NewCdk8sApiObject(construct Construct, id string, obj runtime.Object) ApiObject {
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
			if namespaceCtx := getNamespaceContext(construct); namespaceCtx != "" {
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
	apiobj := cdk8s.NewApiObject(getCdk8sConstruct(construct), &id, &cdk8s.ApiObjectProps{
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
	return &cdk8sApiObject{construct: apiobj}
}

type ApiObjectProps struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Object map[string]interface{}
}

func NewCdk8sApiObjectFromMap(construct Construct, id string, props ApiObjectProps) ApiObject {
	var apiMetadata *cdk8s.ApiObjectMetadata
	if props.GetNamespace() == "" {
		if namespaceCtx := getNamespaceContext(construct); namespaceCtx != "" {
			props.SetNamespace(namespaceCtx)
		}
	}
	apiMetadata = &cdk8s.ApiObjectMetadata{
		Name:        infrahelpers.PtrIfNonEmpty(props.GetName()),
		Namespace:   infrahelpers.PtrIfNonEmpty(props.GetNamespace()),
		Labels:      infrahelpers.PtrMap(props.GetLabels()),
		Annotations: infrahelpers.PtrMap(props.GetAnnotations()),
	}
	if propsObjectMetadata, ok := props.Object["metadata"].(map[string]any); ok {
		if name, ok := propsObjectMetadata["name"].(string); ok {
			apiMetadata.Name = infrahelpers.PtrIfNonEmpty(name)
		}
		if ns, ok := propsObjectMetadata["namespace"].(string); ok {
			apiMetadata.Namespace = infrahelpers.PtrIfNonEmpty(ns)
		}
		mapStringAnyToStringString := func(in map[string]any) *map[string]*string {
			out := map[string]*string{}
			for k, v := range in {
				if v, ok := v.(string); ok {
					out[k] = infrahelpers.PtrIfNonEmpty(v)
				}
			}
			return &out
		}
		if labels, ok := propsObjectMetadata["labels"].(map[string]any); ok {
			apiMetadata.Labels = mapStringAnyToStringString(labels)
		}
		if annotations, ok := propsObjectMetadata["annotations"].(map[string]any); ok {
			apiMetadata.Annotations = mapStringAnyToStringString(annotations)
		}
		if finalizers, ok := propsObjectMetadata["finalizers"].([]any); ok {
			var f []*string
			for _, v := range finalizers {
				if v, ok := v.(string); ok {
					f = append(f, jsii.String(v))
				}
			}
			apiMetadata.Finalizers = &f
		}
	}
	apiobj := cdk8s.NewApiObject(getCdk8sConstruct(construct), &id, &cdk8s.ApiObjectProps{
		ApiVersion: jsii.String(props.APIVersion),
		Kind:       jsii.String(props.Kind),
		Metadata:   apiMetadata,
	})
	for _, field := range infrahelpers.MapKeys(props.Object) {
		if field == "apiVersion" || field == "kind" || field == "metadata" {
			continue
		}
		v := props.Object[field]
		if v != nil {
			apiobj.AddJsonPatch(cdk8s.JsonPatch_Replace(jsii.String("/"+field), v))
		}
	}
	return &cdk8sApiObject{construct: apiobj}
}

func getNamespaceContext(scope Construct) string {
	ns, _ := scope.Node().TryGetContext("namespace").(string)
	return ns
}
