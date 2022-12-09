package utils

import (
	"fmt"
	"testing"
)

func TestDFS(t *testing.T) {
	//t.Skip()
	graph := CreateGraph("a")

	graph.Add(graph.root, "b")
	nodeB := graph.Add(graph.root, "c")
	graph.Add(graph.root, "d")


	graph.Add(nodeB, "e")
	graph.Add(nodeB, "f")
	graph.Add(nodeB, "g")
	graph.Add(nodeB, "h")

	found, pathTraversed := graph.DFS("g")
	fmt.Println(found)
	fmt.Println(pathTraversed)
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

func TestBFS(t *testing.T) {
	//t.Skip()
	graph := CreateGraph("a")

	graph.Add(graph.root, "b")
	nodeB := graph.Add(graph.root, "c")
	graph.Add(graph.root, "d")


	graph.Add(nodeB, "e")
	graph.Add(nodeB, "f")
	graph.Add(nodeB, "g")
	graph.Add(nodeB, "h")

	found, pathTraversed := graph.BFS("g")
	fmt.Println(found)
	fmt.Println(pathTraversed)

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