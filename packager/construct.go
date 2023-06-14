package packager

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
)

type Construct interface {
	Construct(id string) Construct
	Chart(id string, props ChartProps) Chart
	ApiObject(id string, obj runtime.Object) ApiObject
	ApiObjectFromMap(id string, props map[string]any) ApiObject
	GetContext(key string) any
	SetContext(key string, value any)
	ID() string
}

type construct struct {
	node *node[Construct]
}

func (c *construct) Construct(id string) Construct {
	fmt.Println("AddConstruct", c.node.FullID(), id)
	construct := &construct{}
	construct.node = c.node.AddChildNode(id, construct)
	return construct
}

func (c *construct) SetContext(key string, value any) {
	fmt.Println("SetContext", c.node.FullID(), key)
	c.node.SetContext(key, value)
}

func (c *construct) GetContext(key string) any {
	return c.node.GetContext(key)
}

func (c *construct) ID() string {
	return c.node.id
}
