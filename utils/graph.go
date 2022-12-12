package utils

import (
	"fmt"
)

type Node struct {
	parent   *Node
	children []*Node
	data     string
}

func (n *Node) equals(other *Node) bool {
	return n == other
}

type Graph struct {
	root *Node
}

func CreateGraph(data string) *Graph {
	graph := new(Graph)
	node := new(Node)
	node.data = data
	graph.root = node
	return graph
}

func (g *Graph) Add(parent *Node, data string) *Node {
	node := new(Node)
	node.parent = parent
	node.data = data
	parent.children = append(parent.children, node)
	return node
}

func (g *Graph) DFS(target string) (*Node, string) {
	path := ""
	var stack []Node
	var visited []Node
	stack = append(stack, *g.root)
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		visited = append(visited, current)
		path += current.data
		stack = stack[:len(stack)-1]

		if current.data == target {
			return &current, path
		}
		for _, child := range current.children {
			if !ContainsGeneric(visited, *child) {
				stack = append(stack, *child)
			}
		}
	}
	return nil, path
}

func (g *Graph) BFS(target string) (*Node, string) {
	path := ""
	var stack []Node
	var visited []Node
	stack = append(stack, *g.root)
	for len(stack) > 0 {
		current := stack[0]
		visited = append(visited, current)
		fmt.Print(current.data)
		stack = stack[1:]

		if current.data == target {
			return &current, path
		}
		for _, child := range current.children {
			if !ContainsGeneric(visited, *child) {
				stack = append(stack, *child)
			}
		}
	}
	return nil, path
}
