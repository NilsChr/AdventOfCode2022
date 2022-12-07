package day7

import (
	u "advent-of-code-2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day7() {
	lines := u.GetInput("./day7/input.txt")
	task1, task2 := tasks(lines)
	fmt.Println("Task1: ", task1)
	fmt.Println("Task2: ", task2)
}

func tasks(lines []string) (int, int) {
	root, folders := buildTree(lines)
	calcSize(root)
	unused := 70_000_000 - root.size
	var possible_delete []int
	task1 := 0

	for _, folder := range folders {
		if folder.size <= 100_000 {
			task1 += folder.size
		}

		if unused+folder.size > 30_000_000 {
			possible_delete = append(possible_delete, folder.size)
		}
	}

	sort.Slice(possible_delete, func(i, j int) bool {
		return possible_delete[i] < possible_delete[j]
	})

	return task1, possible_delete[0]
}

func calcSize(node *Node) int {
	if len(node.children) == 0 {
		return node.size
	}
	for _, child := range node.children {
		node.size += calcSize(child)
	}
	return node.size
}

func buildTree(lines []string) (*Node, []*Node) {
	root := new(Node)
	root.name = "root"
	current := root
	var folders []*Node
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == "/" {
					current = root
				} else if parts[2] == ".." {
					current = current.parent
				} else {
					current = getChild(*current, parts[2])
				}
			}
		} else {
			child := new(Node)
			child.parent = current
			child.name = parts[1]
			current.children = append(current.children, child)
			if parts[0] == "dir" {
				folders = append(folders, child)
			} else {
				child.size, _ = strconv.Atoi(parts[0])
			}
		}
	}

	return root, folders
}

func getChild(node Node, search string) *Node {
	var out *Node
	for _, child := range node.children {
		if child.name == search {
			out = child
		}
	}
	return out
}
