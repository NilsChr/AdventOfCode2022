package day12

import (
	"advent-of-code-2022/utils"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	lines := utils.GetInput("./input-test.txt")
	grid, start := parseInput(lines)
	printGrid(grid)
	start.x = 4
	start.y = 0
	neighbours := getNeighbours(grid, *start)
	fmt.Println(start)
	fmt.Println(neighbours)
	for _, n := range neighbours {
		fmt.Println(n)
		fmt.Println(grid[n.y][n.x])
	}
}
