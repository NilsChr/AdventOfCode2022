package day7

import (
	u "advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

const (
	CD int = 0
	LS
	MK
)

// TOO HIGH = 1102256

func Day7() {
	lines := u.GetInput2("./day7/input.txt")
	fmt.Println("Task1: ", task1(lines))
	//fmt.Println("Task2: ", lines)
}

func task1(lines []string) int {

	root, dirs := buildTree(lines)
	fmt.Println(root)

	sum := 0
	for i := len(dirs) - 1; i >= 0; i-- {
		fmt.Println(dirs[i].name, dirs[i].size)
		if dirs[i].size <= 100000 {
			sum += dirs[i].size
		}
	}

	return sum
}

func buildTree(lines []string) (*Node, []*Node) {
	root := new(Node)
	root.name = "root"

	current := root
	var dirs []*Node
	var files []*Node
	for _, line := range lines {
		isCommand, values := parseLine(line)
		if isCommand {
			if values[0] == "cd" {
				if values[1] == `/` {
					//	fmt.Println("cd root")
					current = root
				} else if values[1] == ".." {
					//	fmt.Println("cd ", values[1])
					current = current.parent
				} else {
					//	fmt.Println("cd ", values[1])
					current = getChild(*current, values[1])
				}
			} else {
				//fmt.Println("ls ")
			}
		} else {
			if values[0] == "dir" {
				//fmt.Println("Create folder,", values[1])
				child := new(Node)
				child.parent = current
				child.name = values[1]
				dirs = append(dirs, child)
				current.addChild(child)
			} else {
				child := new(Node)
				child.parent = current
				child.size, _ = strconv.Atoi(values[0])
				child.name = values[1]
				files = append(files, child)
				//fmt.Println("Create file,", values[1])
				current.children = append(current.children, child)
			}
		}
	}

	for i := len(files) - 1; i >= 0; i-- {
		files[i].passSizeToParent()
	}

	return root, dirs
}

func parseLine(line string) (bool, []string) {
	parts := strings.Split(line, " ")
	if parts[0] == "$" {
		return true, parts[1:]
	}
	return false, parts
}

func getChild(node Node, search string) *Node {
	//fmt.Println("GET CHILD", search)
	var out *Node
	for _, child := range node.children {
		if child.name == search {
			out = child
		}
	}
	// fmt.Println("Return child, ", out)
	return out
}
