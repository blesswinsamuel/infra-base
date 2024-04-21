package kubegogen

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/rs/zerolog/log"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

type Construct interface { // Scope
	ID() string
	Namespace() string
	Chart(id string, props ChartProps) Construct // CreateScope
	GetContext(key string) any
	SetContext(key string, value any)
	ApiObject(obj runtime.Object) ApiObject          // AddApiObject
	ApiObjectFromMap(props map[string]any) ApiObject // AddApiObjectFromMap
}

type ChartProps struct { // ScopeProps
	Namespace string
}

type scope struct {
	id       string
	props    ChartProps
	context  map[string]any
	parent   *scope
	children []*scope
	objects  []ApiObject
}

func newScope(id string, props ChartProps) Construct {
	return &scope{
		id:      id,
		props:   props,
		context: map[string]any{},
	}
}

func (c *scope) SetContext(key string, value any) {
	c.context[key] = value
}

func (c *scope) GetContext(key string) any {
	for s := c; s != nil; s = s.parent {
		if ctx, ok := s.context[key]; ok {
			return ctx
		}
	}
	return c.context[key]
}

func (c *scope) ID() string {
	return c.id
}

func (c *scope) Chart(id string, props ChartProps) Construct {
	childScope := &scope{
		id:      id,
		props:   props,
		parent:  c,
		context: map[string]any{},
	}
	c.children = append(c.children, childScope)
	if props.Namespace != "" {
		childScope.context["namespace"] = props.Namespace
	}
	return childScope
}

func (c *scope) Namespace() string {
	return c.props.Namespace
}

func (c *scope) ApiObject(obj runtime.Object) ApiObject {
	groupVersionKinds, _, err := infrahelpers.Scheme.ObjectKinds(obj)
	if err != nil {
		log.Panic().Err(err).Msg("ObjectKinds")
	}
	if len(groupVersionKinds) != 1 {
		log.Panic().Msgf("expected 1 groupVersionKind, got %d: %v", len(groupVersionKinds), groupVersionKinds)
	}
	groupVersion := groupVersionKinds[0]
	mobj := infrahelpers.K8sObjectToMap(obj)
	mobj["apiVersion"] = groupVersion.GroupVersion().String()
	mobj["kind"] = groupVersion.Kind
	return c.ApiObjectFromMap(mobj)
}

func (c *scope) ApiObjectFromMap(obj map[string]any) ApiObject {
	props := ApiObjectProps{Unstructured: unstructured.Unstructured{Object: obj}}
	if props.GetNamespace() == "" {
		if namespaceCtx := getNamespaceContext(c); namespaceCtx != "" {
			props.SetNamespace(namespaceCtx)
		}
	}

	apiObject := &apiObject{ApiObjectProps: props}

	// fmt.Println(apiObject.GetKind(), apiObject.GetAPIVersion())
	c.objects = append(c.objects, apiObject)
	return apiObject
}
