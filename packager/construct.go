package packager

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type Construct interface {
	Node() Node[any]
	// Children() bool
}

type cdk8sConstruct struct {
	constructs.Construct
}

func getCdk8sConstruct(scope Construct) constructs.Construct {
	switch scope := scope.(type) {
	case *cdk8sConstruct:
		return scope.Construct
	case *cdk8schart:
		return scope.cdk8sConstruct.Construct
	case *cdk8sApiObject:
		return scope.cdk8sConstruct.Construct
	case *cdk8sapp:
		return scope.cdk8sConstruct.Construct
	}
	return nil
}

func NewCdk8sConstruct(parent Construct, id string) Construct {
	construct := constructs.NewConstruct(getCdk8sConstruct(parent), jsii.String(id))
	return &cdk8sConstruct{construct}
}

func (c *cdk8sConstruct) Node() Node[any] {
	return &cdk8sNode{c.Construct.Node()}
}

type construct struct {
}

func NewConstruct(parent Construct, id string) Construct {
	return &construct{}
}

func (c *construct) Node() Node[any] {
	return nil
}

type cdk8sNode struct {
	node constructs.Node
}

func (c *cdk8sNode) SetContext(key string, value any) {
	c.node.SetContext(&key, value)
}

func (c *cdk8sNode) TryGetContext(key string) any {
	return c.node.TryGetContext(&key)
}

func (c *cdk8sNode) ID() string {
	return *c.node.Id()
}
