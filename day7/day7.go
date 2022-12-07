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

func tasks(lines []string) (int,int) {
	root, folders := buildTree(lines)
	calcSize(root)
	diskspace := 70_000_000
	required  := 30_000_000
	unused  := diskspace - root.size;
	var possible_delete []int
	task1 := 0
	task2 := 0

	for _, folder := range folders {
		if folder.size <= 100_000 {
			task1 += folder.size
		}

		if unused + folder.size > required  {
			possible_delete = append(possible_delete, folder.size)
		}
	}

	sort.Slice(possible_delete, func(i, j int) bool {
		return possible_delete[i] < possible_delete[j]
	})

	task2 = possible_delete[0]

	return task1,task2
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
		isCommand, values := parseLine(line)
		if isCommand {
			if values[0] == "cd" {
				if values[1] == "/" {
					current = root
				} else if values[1] == ".." {
					current = current.parent
				} else {
					current = getChild(*current, values[1])
				}
			}
		} else {
			if values[0] == "dir" {
				child := new(Node)
				child.parent = current
				child.name = values[1]
				current.addChild(child)
				folders = append(folders, child)
			} else {
				child := new(Node)
				child.parent = current
				child.size, _ = strconv.Atoi(values[0])
				child.name = values[1]
				current.children = append(current.children, child)
			}
		}
	}

	return root, folders
}

func parseLine(line string) (bool, []string) {
	parts := strings.Split(line, " ")
	if parts[0] == "$" {
		return true, parts[1:]
	}
	return false, parts
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
