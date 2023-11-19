package kubegogen

import "strings"

// https://kendru.github.io/go/2021/10/26/sorting-a-dependency-graph-in-go/
// https://github.com/kendru/darwin/blob/877d6a81060c1ed6cf6db7b0d6dd2fd4307d6f86/go/depgraph/depgraph.go#L21

// Node represents a node in the DAG.
// type Node[T any] interface {
// 	ID() string
// 	SetContext(key string, value any)
// 	TryGetContext(key string) any
// 	Parent() Node[T]
// 	Children() []Node[T]
// }

type node[T any] struct {
	id       string
	value    T
	parent   *node[T]
	children []*node[T]
	context  map[string]any
}

func newNode[T any](id string, value T) *node[T] {
	return &node[T]{id: id, value: value, context: make(map[string]any)}
}

func (n *node[T]) ID() string {
	return n.id
}

func (n *node[T]) SetContext(key string, value any) {
	n.context[key] = value
}

func (n *node[T]) GetContext(key string) any {
	for n := n; n != nil; n = n.parent {
		if ctx, ok := n.context[key]; ok {
			return ctx
		}
	}
	// panic(fmt.Sprintf("context not found: %s", key))
	return nil
}

// AddNode adds a new node to the DAG.
func (n *node[T]) AddChildNode(id string, value T) *node[T] {
	node := newNode(id, value)
	node.parent = n
	if n != nil {
		n.children = append(n.children, node)
	}
	return node
}

func (n *node[T]) FullID() string {
	ids := []string{}
	for n := n; n != nil; n = n.parent {
		if n.id == "$$root" {
			continue
		}
		ids = append([]string{n.id}, ids...)
	}
	return strings.Join(ids, "-")
}
