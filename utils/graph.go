package utils

import "fmt"

type Node struct {
	parent   *Node
	children []*Node
	data     string
}

type Graph struct {
	root *Node
}

func CreateGraph(data string) *Graph {
	graph := new(Graph)
	node  := new(Node)
	node.data = data
	graph.root = node;
	return graph
}

func (g *Graph) Add(parent *Node, data string ) *Node {
	node := new(Node)
	node.parent = parent;
	node.data = data;
	parent.children = append(parent.children, node)
	return node
}

func (g *Graph) DFS(target string) (*Node, string) {
	path := ""
	var stack []Node
    stack = append(stack, *g.root)
	for len(stack) > 0 {
        current := stack[len(stack)-1]
		path += current.data
        stack = stack[:len(stack)-1]

        if current.data == target {
            return &current, path
        }
        for _, child := range current.children {
            stack = append(stack, *child)
        }
    }
    return nil,path
}

func (g *Graph) BFS(target string) (*Node, string) {
	path := ""
	var stack []Node
    stack = append(stack, *g.root)
	for len(stack) > 0 {
        current := stack[0]
		fmt.Print(current.data)
        stack = stack[1:]

        if current.data == target {
            return &current, path
        }
        for _, child := range current.children {
            stack = append(stack, *child)
        }
    }
    return nil, path
}