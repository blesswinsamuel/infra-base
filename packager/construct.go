package packager

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"k8s.io/apimachinery/pkg/runtime"
)

type Construct interface {
	// Node() Node[any]
	// Type() string // cdk8s or normal
	Construct(id string) Construct
	Chart(id string, props ChartProps) Chart
	ApiObject(id string, obj runtime.Object) ApiObject
	ApiObjectFromMap(id string, props map[string]any) ApiObject
	GetContext(key string) any
	SetContext(key string, value any)
	ID() string
	// getConstruct() Construct
	// Children() bool
}

type cdk8sConstruct struct {
	construct constructs.Construct
}

// func getCdk8sConstruct(scope Construct) constructs.Construct {
// 	return scope.getConstruct().(*cdk8sConstruct).Construct
// }

// func NewCdk8sConstruct(parent Construct, id string) Construct {
// 	construct := constructs.NewConstruct(getCdk8sConstruct(parent), jsii.String(id))
// 	return &cdk8sConstruct{construct}
// }

// func (c *cdk8sConstruct) Node() Node[any] {
// 	return &cdk8sNode{c.construct.Node()}
// }

func (c *cdk8sConstruct) SetContext(key string, value any) {
	c.construct.Node().SetContext(&key, value)
}

func (c *cdk8sConstruct) GetContext(key string) any {
	return c.construct.Node().TryGetContext(&key)
}

func (c *cdk8sConstruct) ID() string {
	return *c.construct.Node().Id()
}

// func (c *cdk8sConstruct) Type() string {
// 	return "cdk8s"
// }

func (c *cdk8sConstruct) Construct(id string) Construct {
	construct := constructs.NewConstruct(c.construct, &id)
	return &cdk8sConstruct{construct: construct}
}

// func (c *cdk8sConstruct) getConstruct() Construct {
// 	return c
// }

type construct struct {
	node *node[Construct]
}

// func NewConstruct(parent Construct, id string) Construct {
// 	return &construct{}
// }

// func (c *construct) Node() Node[any] {
// 	return nil
// }

// func (c *construct) Type() string {
// 	return "normal"
// }

// func newConstruct(id string, parent *construct, child Construct) *construct {
// 	construct := &construct{}
// 	construct.node = parent.node.AddChildNode(id, child)
// 	return construct
// }

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

// func (c *construct) getConstruct() Construct {
// 	return c
// }

// type cdk8sNode struct {
// 	node constructs.Node
// }

// func (c *cdk8sNode) SetContext(key string, value any) {
// 	c.node.SetContext(&key, value)
// }

// func (c *cdk8sNode) TryGetContext(key string) any {
// 	return c.node.TryGetContext(&key)
// }

// func (c *cdk8sNode) ID() string {
// 	return *c.node.Id()
// }

// func (c *cdk8sNode) Children() []Node[any] {
// 	return nil
// }

// func (c *cdk8sNode) Parent() Node[any] {
// 	return nil
// }
