package day7

type Node struct {
	parent   *Node
	children []*Node
	name     string
	size     int
}

func (n *Node) isFolder() bool {
	return len(n.children) == 0
}

func (n *Node) addChild(node *Node) {
	n.children = append(n.children, node)
}

func (n *Node) passSizeToParent() {
	if n.parent == nil {
		return
	}
	n.parent.size += n.size
	n.parent.passSizeToParent()
}
