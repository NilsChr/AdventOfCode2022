package day8

import (
	"advent-of-code-2022/utils"
	"fmt"
	"sort"
	"strconv"
)

func Day8() {
	lines := utils.GetInput("./day8/input.txt")
	fmt.Println("Task 1: ", task1(lines))
	fmt.Println("Task 2: ", task2(lines))
}

func task1(lines []string) int {
	grid := inputToArray(lines)
	count := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if !checkVisibility(grid, x, y) {
				continue
			}
			count++
		}
	}
	return count
}

func task2(lines []string) int {
	grid := inputToArray(lines)
	var scores []int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			scores = append(scores, checkVisibility2(grid, x, y))
		}
	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})
	return scores[0]
}

func checkVisibility(grid [][]int, x int, y int) bool {
	a,_ := checkVisibilityLeft(grid, x, y)
	b,_ := checkVisibilityRight(grid, x, y)
	c,_ := checkVisibilityUp(grid, x, y)
	d,_ := checkVisibilityDown(grid, x, y)
	return (a + b + c + d) > 0
}

func checkVisibility2(grid [][]int, x int, y int) int {
	_,l := checkVisibilityLeft(grid, x, y)
	_,r := checkVisibilityRight(grid, x, y)
	_,u := checkVisibilityUp(grid, x, y)
	_,d := checkVisibilityDown(grid, x, y)
	return l * r * u * d
}

func checkVisibilityLeft(grid [][]int, x int, y int) (int,int) {
	value := grid[y][x]
	move := 0
	for i := x - 1; i >= 0; i-- {
		move++
		if grid[y][i] >= value {
			return 0,move
		}
	}
	return 1,move
}

func checkVisibilityRight(grid [][]int, x int, y int) (int,int) {
	value := grid[y][x]
	move := 0
	for i := x + 1; i <= len(grid[y])-1; i++ {
		move++
		if grid[y][i] >= value {
			return 0,move
		}
	}
	return 1,move
}

func checkVisibilityUp(grid [][]int, x int, y int) (int,int) {
	value := grid[y][x]
	move := 0
	for i := y - 1; i >= 0; i-- {
		move++
		if grid[i][x] >= value {
			return 0,move
		}
	}
	return 1,move
}

func checkVisibilityDown(grid [][]int, x int, y int) (int, int) {
	value := grid[y][x]
	move := 0
	for i := y + 1; i <= len(grid)-1; i++ {
		move++
		if grid[i][x] >= value {
			return 0,move
		}
	}
	return 1,move
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
