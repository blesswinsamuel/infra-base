package packager

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
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

type ApiObjectProps struct {
	unstructured.Unstructured
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
	groupVersion := groupVersionKinds[0]
	mobj := infrahelpers.K8sObjectToMap(obj)
	mobj["apiVersion"] = groupVersion.GroupVersion().String()
	mobj["kind"] = groupVersion.Kind
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
}
