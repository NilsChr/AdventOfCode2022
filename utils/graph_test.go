package utils

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	//t.Skip()
	graph := CreateGraph("a")

	nodeB := graph.Add(graph.root, "b")
	graph.Add(graph.root, "c")
	graph.Add(graph.root, "d")


	graph.Add(nodeB, "e")
	graph.Add(nodeB, "f")
	graph.Add(nodeB, "g")
	graph.Add(nodeB, "h")

	found := graph.DFS("g")
	fmt.Println(found)

	current := found
	path := ""
	for {
		path += current.data

		current = current.parent
		
		if current.parent == nil {
			path += current.data

			break
		}
	}
	fmt.Println(path)
}
