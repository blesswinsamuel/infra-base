package packager

import "fmt"

// https://kendru.github.io/go/2021/10/26/sorting-a-dependency-graph-in-go/
// https://github.com/kendru/darwin/blob/877d6a81060c1ed6cf6db7b0d6dd2fd4307d6f86/go/depgraph/depgraph.go#L21

// Node represents a node in the DAG.
type Node[T any] interface {
	ID() string
	SetContext(key string, value any)
	TryGetContext(key string) any
}

type node[T any] struct {
	id       string
	Value    T
	Parent   *node[T]
	Children []*node[T]
}

func (n *node[T]) ID() string {
	return n.id
}

func (n *node[T]) SetContext(key string, value any) {
}

func (n *node[T]) TryGetContext(key string) any {
	return nil
}

type DAG[T any] interface {
	AddNode(id string, value T) *node[T]
}

// DAG represents the Directed Acyclic Graph.
type dag[T any] struct {
	Nodes []*node[T]
}

func NewDAG[T any]() DAG[T] {
	return &dag[T]{}
}

// AddNode adds a new node to the DAG.
func (d *dag[T]) AddNode(id string, value T) *node[T] {
	node := &node[T]{id: id, Value: value}
	d.Nodes = append(d.Nodes, node)
	return node
}

// AddEdge adds an edge between two nodes in the DAG.
func (d *dag[T]) AddEdge(parent *node[T], child *node[T]) {
	parent.Children = append(parent.Children, child)
	child.Parent = parent
}

// DepthFirstSearch performs a depth-first search traversal of the DAG.
func (d *dag[T]) DepthFirstSearch(node *node[T]) {
	fmt.Printf("%s ", node.ID())

	for _, child := range node.Children {
		d.DepthFirstSearch(child)
	}
}

// func main() {
// 	// Create a new DAG.
// 	dag := &dag{}

// 	// Create nodes.
// 	node1 := dag.AddNode(1)
// 	node2 := dag.AddNode(2)
// 	node3 := dag.AddNode(3)
// 	node4 := dag.AddNode(4)
// 	node5 := dag.AddNode(5)
// 	node6 := dag.AddNode(6)

// 	// Add edges between nodes.
// 	dag.AddEdge(node1, node2)
// 	dag.AddEdge(node1, node3)
// 	dag.AddEdge(node2, node4)
// 	dag.AddEdge(node2, node5)
// 	dag.AddEdge(node3, node5)
// 	dag.AddEdge(node4, node6)
// 	dag.AddEdge(node5, node6)

// 	// Perform a depth-first search on the DAG starting from node1.
// 	fmt.Println("Depth-First Search:")
// 	dag.DepthFirstSearch(node1)
// 	fmt.Println()
// }
