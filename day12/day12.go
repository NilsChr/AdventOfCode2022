package day12

import (
	"advent-of-code-2022/utils"
	"fmt"
	"sort"

	//"math"
	"os"
	"os/exec"
	"time"
)

const (
	START  int = 83
	TARGET     = 69
)

func Day12() {
	lines := utils.GetInput("./day12/input.txt")
	fmt.Println("Task1: ", task1(lines))
	fmt.Println("Task2: ", task2(lines))

}

func task1(lines []string) int {
	grid, start, end := parseInput(lines)

	path := searchGrid(grid, *start, *end)

	//fmt.Println("Path", path)
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

func task2(lines []string) int {
	grid, _, end := parseInput(lines)

	var pathLengs []int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x].value != 97 {
				continue
			}
			start := grid[y][x]
			path := searchGrid(grid, start, *end)
			if path == nil {
				continue
			}

			steps := 0

			next := path.parent

			for next != nil {
				steps++
				next = next.parent
			}
			pathLengs = append(pathLengs, steps)

		}
	}

	/*
	sort.Slice(pathLengs, func(i,j int) bool {
		return i < j
	})*/

	sort.Ints (pathLengs)
/*

	fmt.Println(pathLengs[0])
	fmt.Println(pathLengs[1])
	fmt.Println(pathLengs[2])

	fmt.Println(pathLengs)
*/
	return pathLengs[0]
}

func parseInput(lines []string) ([][]Node, *Node, *Node) {
	grid := make([][]Node, 0)
	start := new(Node)
	end := new(Node)

	for y, line := range lines {
		var row []Node
		//fmt.Println(line)

		for x, letter := range line {
			value := int(letter)
			node := new(Node)
			node.x = x
			node.y = y
			node.value = value

			if value == START {
				//fmt.Println("FOUND START")
				start = node
				node.value = int('a')
			}

			if value == TARGET {
				//fmt.Println("FOUND END")
				end = node
				node.value = int('z')
			}
			row = append(row, *node)

			//row = append(row, value)
		}
		grid = append(grid, row)
	}

	//printGrid(grid)
	//wait()
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

	//fmt.Println("Searching from:", start)
	//printGrid(grid)

	var visited2 []Node //make(map[Node]bool)
	var queue []Node
	queue = append(queue, grid[start.y][start.x])

	//in := 0
	for len(queue) > 0 {
		//in++
		current := queue[0]
		queue = queue[1:]
		//current := queue[len(queue)-1]
		//queue = queue[:len(queue)-1]
		//printGrid2(grid, current, current, visited2, queue, []Node)
		//fmt.Println(string(current.value), current.getPos())

		visited2 = append(visited2, current)

		//visited2[current] = true

		if current.x == end.x && current.y == end.y {
			//fmt.Println("Found: ", current)

			return &current
		}

		// Get Neighbours
		neighbours := getNeighbours(grid, current)
		for _, n := range neighbours {
			next := grid[n.y][n.x]
			//if in % 0 == 0 {
			//	printGrid2(grid, current, next, visited2, queue, neighbours)
			//}

			if next.value == TARGET {
				//fmt.Println("Found target! Adding:", next, (next.value >= current.value))
				//fmt.Scanln()
				wait()
				next.parent = &current
				queue = append(queue, next)
				continue

			}

			if Contains(visited2, next) {
				//fmt.Println("Allready added")
				//fmt.Scanln()
				//wait()
				continue
			}
			//diff := math.Max(float64(current.value), float64(next.value)) - math.Min(float64(current.value), float64(next.value))
			//fmt.Println("DIFF", diff)
			//if diff <= 1 || current.value == START /*current.value+1 <= next.value*/ {
			if next.value <= current.value+1 {
				if !Contains(queue, next) {
					//	fmt.Println("Adding to queue:", next, (next.value >= current.value))
					//fmt.Scanln()
					//wait()
					next.parent = &current
					queue = append(queue, next)
				} else {
					//	fmt.Println("Allready queued")
					//fmt.Scanln()
					wait()
				}

			} else {
				//fmt.Println("Nothing done")
				//fmt.Scanln()
				wait()
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
	//fmt.Println("Check contains,", node, list)
	for _, n := range list {
		if node.x == n.x && node.y == n.y {
			//fmt.Println("true")
			//fmt.Scanln()
			return true
		}
	}
	//fmt.Println("false")
	//fmt.Scanln()
	return false
}

func printGrid(grid [][]Node) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	row := ""
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			row += string(rune(grid[y][x].value))
		}
		fmt.Println(row)
		row = ""
	}

}

func printGrid2(grid [][]Node, focus Node, look Node, visited []Node, queue []Node, neighbours []Vec2) {
	//cmd := exec.Command("cmd", "/c", "cls")
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	/*
		fmt.Println("Focus:", focus, grid[focus.y][focus.x].value, string(rune(grid[focus.y][focus.x].value)))
		fmt.Println("Looking at:", look, grid[look.y][look.x].value, string(rune(grid[look.y][look.x].value)))

		fmt.Println("Visited", visited)
		fmt.Println("Queue", queue)
		fmt.Println("Neighbours", neighbours)
		fmt.Println("")
	*/
	row := ""
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			val := ""
			if focus.x == x && focus.y == y {
				//row += fmt.Sprint("\033[35m" + "*" + "\033[0m")
				val = fmt.Sprint("\033[35m" + string(rune(grid[y][x].value)) + "\033[0m")
			} else if look.x == x && look.y == y {
				//row += fmt.Sprint("\033[32m" + "¤" + "\033[0m")
				val = fmt.Sprint("\033[32m" + string(rune(grid[y][x].value)) + "\033[0m")
			} else {
				val = string(rune(grid[y][x].value))
				if Contains(visited, grid[y][x]) {
					val = fmt.Sprint("\033[22m" + "." + "\033[0m")
				}
			}

			row += val

		}
		fmt.Println(row)
		row = ""
	}
	fmt.Println("QUEUE:", len(queue))
	fmt.Println("VISITED:", len(visited))
	//time.Sleep(1 * time.Second)
	//fmt.Scanln()
	wait()
}

func printGrid3(grid [][]Node, focus Node, visited []Node, queue []Node) {
	//cmd := exec.Command("cmd", "/c", "cls")
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	/*
		fmt.Println("Focus:", focus, grid[focus.y][focus.x].value, string(rune(grid[focus.y][focus.x].value)))
		fmt.Println("Looking at:", look, grid[look.y][look.x].value, string(rune(grid[look.y][look.x].value)))

		fmt.Println("Visited", visited)
		fmt.Println("Queue", queue)
		fmt.Println("Neighbours", neighbours)
		fmt.Println("")
	*/
	row := ""
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			val := ""
			if focus.x == x && focus.y == y {
				//row += fmt.Sprint("\033[35m" + "*" + "\033[0m")
				val = fmt.Sprint("\033[35m" + string(rune(grid[y][x].value)) + "\033[0m")
			} else {
				val = string(rune(grid[y][x].value))
				if Contains(visited, grid[y][x]) {
					val = fmt.Sprint("\033[22m" + "." + "\033[0m")
				}
			}

			row += val

		}
		fmt.Println(row)
		row = ""
	}

	fmt.Println("QUEUE:", len(queue))
	fmt.Println("VISITED:", len(visited))

	//time.Sleep(1 * time.Second)
	//fmt.Scanln()
	wait()
}

func wait() {
	//fmt.Scanln()
	time.Sleep(1 * time.Nanosecond)
}
