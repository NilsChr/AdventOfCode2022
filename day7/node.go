package day7

type Node struct {
	parent   *Node
	children []*Node
	name     string
	size     int
}