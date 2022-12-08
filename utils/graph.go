package utils

func Contains(haystack []Node, needle Node) bool {
    for _, val := range haystack {
        if val.data == needle.data {
            return true
        }
    }
    return false
}

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

func (g *Graph) DFS(target string) *Node {

	var stack []Node
    stack = append(stack, *g.root)
	for len(stack) > 0 {
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        if current.data == target {
            return &current
        }
        for _, child := range current.children {
            stack = append(stack, *child)
        }
    }
    return nil
}