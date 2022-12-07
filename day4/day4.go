package day4

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day4() {
	fmt.Println("Day 4")
	lines := utils.GetInput("./day4/input.txt")
	fmt.Println("Task1: ", task1(lines))
	fmt.Println("Task2: ", task2(lines))
}

func task1(lines []string) int {
	sum := 0
	for _, line := range lines {
		x1, x2, x3, x4 := getSections(line)
		if checkContains(x1, x2, x3, x4) || checkContains(x3, x4, x1, x2) {
			sum++
		}
	}
	return sum
}

func task2(lines []string) int {
	sum := 0
	for _, line := range lines {
		x1, x2, x3, x4 := getSections(line)
		if checkOverlap(x1, x2, x3, x4) {
			sum++
		}
	}
	return sum
}

func getSections(line string) (int, int, int, int) {
	lines := strings.Split(line, ",")
	a := strings.Split(lines[0], "-")
	b := strings.Split(lines[1], "-")
	x1, _ := strconv.Atoi(a[0])
	x2, _ := strconv.Atoi(a[1])
	x3, _ := strconv.Atoi(b[0])
	x4, _ := strconv.Atoi(b[1])
	return x1, x2, x3, x4
}

func checkContains(x1 int, x2 int, x3 int, x4 int) bool {
	return x3 >= x1 && x4 <= x2
}

func checkOverlap(x1 int, x2 int, x3 int, x4 int) bool {
	return x2 >= x3 && x1 <= x4
}
