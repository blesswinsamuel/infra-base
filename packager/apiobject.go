package packager

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/goccy/go-yaml"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

type ApiObject interface {
	Construct
	metav1.Type
	metav1.Object
	ToYAML() []byte
}

type cdk8sApiObject struct {
	cdk8sConstruct
	ApiObjectProps
	construct cdk8s.ApiObject
}

func (c *cdk8sConstruct) ApiObject(id string, obj runtime.Object) ApiObject {
	// groupVersionKinds, _, err := infrahelpers.Scheme.ObjectKinds(obj)
	// if err != nil {
	// 	panic(err)
	// }
	// if len(groupVersionKinds) != 1 {
	// 	panic(fmt.Errorf("expected 1 groupVersionKind, got %d: %v", len(groupVersionKinds), groupVersionKinds))
	// }
	// var metadata metav1.ObjectMeta
	// if obj, ok := obj.(metav1.Object); ok {
	// 	metadata = metav1.ObjectMeta{
	// 		Name:        obj.GetName(),
	// 		Namespace:   obj.GetNamespace(),
	// 		Labels:      obj.GetLabels(),
	// 		Annotations: obj.GetAnnotations(),
	// 	}
	// }
	// groupVersion := groupVersionKinds[0]
	mobj := infrahelpers.K8sObjectToMap(obj)
	// return c.ApiObjectFromMap(id, ApiObjectProps{
	// 	TypeMeta: metav1.TypeMeta{
	// 		Kind:       groupVersion.Kind,
	// 		APIVersion: groupVersion.GroupVersion().String(),
	// 	},
	// 	ObjectMeta: metadata,
	// 	Object:     mobj,
	// })
	return c.ApiObjectFromMap(id, mobj)
}

type ApiObjectProps struct {
	unstructured.Unstructured
}

func (c *cdk8sConstruct) ApiObjectFromMap(id string, obj map[string]any) ApiObject {
	props := ApiObjectProps{Unstructured: unstructured.Unstructured{Object: obj}}
	var apiMetadata *cdk8s.ApiObjectMetadata
	if props.GetNamespace() == "" {
		if namespaceCtx := getNamespaceContext(c); namespaceCtx != "" {
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
	apiobj := cdk8s.NewApiObject(c.construct, &id, &cdk8s.ApiObjectProps{
		ApiVersion: jsii.String(props.GetAPIVersion()),
		Kind:       jsii.String(props.GetKind()),
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
	return &cdk8sApiObject{construct: apiobj, ApiObjectProps: props}
}

func (a *cdk8sApiObject) ToYAML() []byte {
	return nil
}

func getNamespaceContext(scope Construct) string {
	ns, _ := scope.GetContext("namespace").(string)
	return ns
}

type apiObject struct {
	construct
	ApiObjectProps
}

func (c *construct) ApiObject(id string, obj runtime.Object) ApiObject {
	groupVersionKinds, _, err := infrahelpers.Scheme.ObjectKinds(obj)
	if err != nil {
		panic(err)
	}
	if len(groupVersionKinds) != 1 {
		panic(fmt.Errorf("expected 1 groupVersionKind, got %d: %v", len(groupVersionKinds), groupVersionKinds))
	}
	// var metadata metav1.ObjectMeta
	// if obj, ok := obj.(metav1.Object); ok {
	// 	metadata = metav1.ObjectMeta{
	// 		Name:        obj.GetName(),
	// 		Namespace:   obj.GetNamespace(),
	// 		Labels:      obj.GetLabels(),
	// 		Annotations: obj.GetAnnotations(),
	// 	}
	// }
	groupVersion := groupVersionKinds[0]
	mobj := infrahelpers.K8sObjectToMap(obj)
	mobj["apiVersion"] = groupVersion.GroupVersion().String()
	mobj["kind"] = groupVersion.Kind
	// apiobj := ApiObjectProps{
	// 	TypeMeta: metav1.TypeMeta{
	// 		Kind:       groupVersion.Kind,
	// 		APIVersion: groupVersion.GroupVersion().String(),
	// 	},
	// 	ObjectMeta: metadata,
	// 	Object:     mobj,
	// }
	return c.ApiObjectFromMap(id, mobj)
}

func (c *construct) ApiObjectFromMap(id string, obj map[string]any) ApiObject {
	props := ApiObjectProps{Unstructured: unstructured.Unstructured{Object: obj}}
	if props.GetNamespace() == "" {
		if namespaceCtx := getNamespaceContext(c); namespaceCtx != "" {
			props.SetNamespace(namespaceCtx)
		}
	}

	apiObject := &apiObject{ApiObjectProps: props}

	// fmt.Println(apiObject.GetKind(), apiObject.GetAPIVersion())
	c.node.AddChildNode(id, apiObject)
	return apiObject
}

func (a *apiObject) ToYAML() []byte {
	b := bytes.NewBuffer(nil)
	enc := yaml.NewEncoder(b, yaml.IndentSequence(true), yaml.UseLiteralStyleIfMultiline(true), yaml.UseSingleQuote(false))
	sortedMap := yaml.MapSlice{}
	keys := []string{}
	for k := range a.Object {
		if k == "apiVersion" || k == "kind" || k == "metadata" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	keys = append([]string{"apiVersion", "kind", "metadata"}, keys...)
	for _, key := range keys {
		sortedMap = append(sortedMap, yaml.MapItem{
			Key:   key,
			Value: a.Object[key],
		})
	}
	err := enc.Encode(sortedMap)
	if err != nil {
		panic(err)
	}
	return b.Bytes()
	// return infrahelpers.K8sObjectToYaml(&a.Unstructured)
}
