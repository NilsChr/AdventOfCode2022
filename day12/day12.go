package day12

import (
	"advent-of-code-2022/utils"
	"fmt"
	"os"
	"os/exec"
)

const (
	START int = 83
	END       = 69
)

func Day12() {
	lines := utils.GetInput("./day12/input-test.txt")
	fmt.Println("Task1: ", task1(lines))
}

func task1(lines []string) int {
	grid, start := parseInput(lines)

	path := searchGrid(grid, *start)

	fmt.Println("Path", path)
	if path == nil {
		return -1
	}

	steps := 0

	next := path.parent

	for next != nil {
		steps++
		next = next.parent
	}

	return steps
}

func task2() {

}

func parseInput(lines []string) ([][]Node, *Node) {
	grid := make([][]Node, 0)
	start := new(Node)

	for y, line := range lines {
		var row []Node
		for x, letter := range line {
			value := int(letter)
			node := new(Node)
			node.x = x
			node.y = y
			node.value = value
			row = append(row, *node)

			if value == START {
				start = node
			}
			//row = append(row, value)
		}
		grid = append(grid, row)
	}

	return grid, start
}

type Node struct {
	x      int
	y      int
	value  int
	parent *Node
}

func (n *Node) getPos() Vec2 {
	v := new(Vec2)
	v.x = n.x
	v.y = n.y
	return *v
}

type Vec2 struct {
	x int
	y int
}

func searchGrid(grid [][]Node, start Node) *Node {

	fmt.Println("Searching from:", start)
	printGrid(grid)

	var visited2 = make(map[Vec2]bool)
	var queue []Node
	queue = append(queue, grid[start.y][start.x])

	for len(queue) > 0 {

		current := queue[0]
		printGrid2(grid, current.getPos(), current.getPos(), visited2, queue)
		fmt.Println(string(current.value), current.getPos())
		queue = queue[1:]
		//visited = append(visited, current)
		visited2[current.getPos()] = true

		if current.value == END {
			fmt.Println("Found: ", current)
			return &current
		}

		// Get Neighbours
		neighbours := getNeighbours(grid, current)
		for _, n := range neighbours {
			next := grid[n.y][n.x]
			printGrid2(grid, current.getPos(), next.getPos(), visited2, queue)

			if visited2[next.getPos()] {
				fmt.Println("Allready added")
				fmt.Scanln()
			} else if current.value+1 <= next.value {
				fmt.Println("Adding to queue:", next)
				fmt.Scanln()
				next.parent = &current
				queue = append(queue, next)
			} else {
				fmt.Println("Nothing done");
				fmt.Scanln();
			}

			/*
			if !visited2[next.getPos()] && current.value+1 < next.value {
				fmt.Println("Adding to queue:", next)
				fmt.Scanln()

				//time.Sleep(1 * time.Second)

				next.parent = &current
				queue = append(queue, next)
			}
			*/

		}

	}
	return nil
}

func getNeighbours(grid [][]Node, pos Node) []Vec2 {
	var out []Vec2

	if pos.x > 0 && pos.y > 0 {
		left := new(Vec2)
		left.y = pos.y
		left.x = pos.x - 1
		out = append(out, *left)

		up := new(Vec2)
		up.y = pos.y - 1
		up.x = pos.x
		out = append(out, *up)
	}

	if pos.x < len(grid[0])-1 && pos.y < len(grid)-1 {
		right := new(Vec2)
		right.y = pos.y
		right.x = pos.x + 1
		out = append(out, *right)

		down := new(Vec2)
		down.y = pos.y + 1
		down.x = pos.x
		out = append(out, *down)
	}

	return out
}

func printGrid(grid [][]Node) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	row := ""
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0])-1; x++ {
			row += string(rune(grid[y][x].value))
		}
		fmt.Println(row)
		row = ""
	}

}

func printGrid2(grid [][]Node, focus Vec2, look Vec2, visited map[Vec2]bool, queue []Node) {
	//cmd := exec.Command("cmd", "/c", "cls")
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("Focus:",     focus, grid[focus.y][focus.x].value, string(rune(grid[focus.y][focus.x].value)))
	fmt.Println("Looking at:", look, grid[look.y][look.x].value  , string(rune(grid[look.y][look.x].value)))

	fmt.Println("Visited", visited)
	fmt.Println("Queue", queue)
	row := ""
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0])-1; x++ {
			if focus.x == x && focus.y == y {
				//row += fmt.Sprint("\033[35m" + "*" + "\033[0m")
				row += fmt.Sprint("\033[35m" +  string(rune(grid[y][x].value)) + "\033[0m")
			} else if look.x == x && look.y == y {
				//row += fmt.Sprint("\033[32m" + "¤" + "\033[0m")
				row += fmt.Sprint("\033[32m" +  string(rune(grid[y][x].value)) + "\033[0m")
			} else {
				row += string(rune(grid[y][x].value))
			}

		}
		fmt.Println(row)
		row = ""
	}
	//time.Sleep(1 * time.Second)
	fmt.Scanln()
}
