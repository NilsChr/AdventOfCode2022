package day8

import (
	"advent-of-code-2022/utils"
	"fmt"
	"sort"
	"strconv"
)

func Day8() {
	lines := utils.GetInput("./day8/input.txt")
	task1, task2 := tasks(lines)
	fmt.Println("Task 1: ", task1)
	fmt.Println("Task 2: ", task2)
}

func tasks(lines []string) (int, int) {
	grid := inputToArray(lines)
	var scores []int
	count := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			task1, task2 := checkVisibility(grid, x, y)
			scores = append(scores, task2)
			if task1 {
				count++
			}
		}

	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})
	return count, scores[0]
}

func checkVisibility(grid [][]int, x int, y int) (bool, int) {
	a, l0 := checkVisibilityDir(grid, x, y, -1, 0)
	b, r0 := checkVisibilityDir(grid, x, y, 1, 0)
	c, u0 := checkVisibilityDir(grid, x, y, 0, -1)
	d, d0 := checkVisibilityDir(grid, x, y, 0, 1)
	return (a + b + c + d) > 0, l0 * r0 * u0 * d0
}

func checkVisibilityDir(grid [][]int, x int, y int, dx int, dy int) (int, int) {
	sx := x
	sy := y
	value := grid[y][x]
	move := 0

	for {
		sx += dx
		sy += dy

		if (sx < 0 || sx > len(grid[0])-1) || (sy < 0 || sy > len(grid)-1) {
			break
		}
		move++
		if grid[sy][sx] >= value {
			return 0, move
		}

	}
	return 1, move
}

func inputToArray(lines []string) [][]int {
	a := make([][]int, len(lines))
	for i := range a {
		a[i] = make([]int, len(lines[0]))
	}

	for y, line := range lines {
		for x := range line {
			a[y][x], _ = strconv.Atoi(string(lines[y][x]))
		}
	}

	return a
}
