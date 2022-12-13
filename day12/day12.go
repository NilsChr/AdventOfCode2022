package day12

import (
	"advent-of-code-2022/utils"
	"fmt"
)

const (
	START  int = 83
	TARGET int = 69
)

func Day12() {
	lines := utils.GetInput("./day12/input.txt")
	fmt.Println("Task1: ", task1(lines))
	fmt.Println("Task2: ", task2(lines))
}

func task1(lines []string) int {
	grid, start, end := parseInput(lines)
	path := searchGrid(grid, *start, *end)
	return getPathLength(path)
}

func task2(lines []string) int {
	grid, _, end := parseInput(lines)
	path := searchGridValue(grid, *end, 97)
	return getPathLength(path)
}

func getPathLength(node *Node) int {
	steps := 0
	if node == nil {
		return -1
	}
	next := node.parent
	for next != nil {
		steps++
		next = next.parent
	}
	return steps
}

func parseInput(lines []string) ([][]Node, *Node, *Node) {
	grid := make([][]Node, 0)
	start := new(Node)
	end := new(Node)

	for y, line := range lines {
		var row []Node
		for x, letter := range line {
			value := int(letter)
			node := new(Node)
			node.x = x
			node.y = y
			node.value = value
			if value == START {
				start = node
				node.value = int('a')
			}

			if value == TARGET {
				end = node
				node.value = int('z')
			}
			row = append(row, *node)
		}
		grid = append(grid, row)
	}
	return grid, start, end
}

type Node struct {
	x      int
	y      int
	value  int
	parent *Node
}

type Vec2 struct {
	x int
	y int
}

func searchGrid(grid [][]Node, start Node, end Node) *Node {
	var visited []Node
	var queue []Node
	queue = append(queue, grid[start.y][start.x])

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visited = append(visited, current)
		if current.x == end.x && current.y == end.y {
			return &current
		}

		neighbours := getNeighbours(grid, current)
		for _, n := range neighbours {
			next := grid[n.y][n.x]
			if Contains(visited, next) {
				continue
			}
			if next.value <= current.value+1 && !Contains(queue, next) {
				next.parent = &current
				queue = append(queue, next)
			}
		}
	}
	return nil
}

func searchGridValue(grid [][]Node, start Node, targetValue int) *Node {
	var visited []Node
	var queue []Node
	queue = append(queue, grid[start.y][start.x])

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visited = append(visited, current)
		if current.value == targetValue {
			return &current
		}
		neighbours := getNeighbours(grid, current)
		for _, n := range neighbours {
			next := grid[n.y][n.x]
			if Contains(visited, next) {
				continue
			}
			if current.value-next.value <= 1 && !Contains(queue, next) {
				next.parent = &current
				queue = append(queue, next)
			}
		}
	}
	return nil
}

func getNeighbours(grid [][]Node, pos Node) []Vec2 {
	var out []Vec2
	if pos.x > 0 {
		left := new(Vec2)
		left.y = pos.y
		left.x = pos.x - 1
		out = append(out, *left)
	}
	if pos.y > 0 {
		up := new(Vec2)
		up.y = pos.y - 1
		up.x = pos.x
		out = append(out, *up)
	}
	if pos.x < len(grid[0])-1 {
		right := new(Vec2)
		right.y = pos.y
		right.x = pos.x + 1
		out = append(out, *right)
	}
	if pos.y < len(grid)-1 {
		down := new(Vec2)
		down.y = pos.y + 1
		down.x = pos.x
		out = append(out, *down)
	}
	return out
}

func Contains(list []Node, node Node) bool {
	for _, n := range list {
		if node.x == n.x && node.y == n.y {
			return true
		}
	}
	return false
}