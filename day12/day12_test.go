package day12

import (
	"advent-of-code-2022/utils"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	lines := utils.GetInput("./input-test.txt")
	grid, start, end := parseInput(lines)
	fmt.Println(end)
	printGrid(grid)
	start.x = 2
	start.y = 4
	neighbours := getNeighbours(grid, *start)
	fmt.Println(start)
	fmt.Println(neighbours)
	for _, n := range neighbours {
		fmt.Println(n)
		fmt.Println(grid[n.y][n.x])
	}
}
