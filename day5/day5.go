package day5

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day5() {
	fmt.Println("Day 5")
	lines := utils.GetInput("./day5/input.txt")
	fmt.Println("Task1: ", task1(lines))
	fmt.Println("Task2: ", task2(lines))
}

func task1(lines []string) string {
	stacks, lastindex := initStacks(lines)
	instructions := lines[lastindex+2:]

	for _, line := range instructions {
		amount, from, to := parseInstructions(line)
		for i := 0; i < amount; i++ {
			data, found := stacks[from].Pop()
			if found {
				stacks[to].Push(data)
			}
		}
	}

	return getPuzzleAnswer(&stacks)
}

func task2(lines []string) string {
	stacks, lastindex := initStacks(lines)
	instructions := lines[lastindex+2:]

	for _, line := range instructions {
		amount, from, to := parseInstructions(line)
		packages := ""
		for i := 0; i < amount; i++ {
			data, found := stacks[from].Pop()
			if found {
				packages += data
			}
		}
		for i := len(packages) - 1; i >= 0; i-- {
			stacks[to].Push(string(packages[i]))
		}
	}
	return getPuzzleAnswer(&stacks)
}

func getPuzzleAnswer(stacks *[9]utils.Stack[string]) string {
	out := ""
	for _, stack := range *stacks {
		data, exists := stack.Peak()
		if exists {
			out += data
		}
	}
	return out
}

func parseInstructions(line string) (int, int, int) {
	parts := strings.Split(line, " ")
	amount, _ := strconv.Atoi(parts[1])
	from, _ := strconv.Atoi(parts[3])
	to, _ := strconv.Atoi(parts[5])
	return amount, from - 1, to - 1
}

func initStacks(lines []string) ([9]utils.Stack[string], int) {
	var cols [9]string
	var stacks [9]utils.Stack[string]
	lastline := 0

	for i, line := range lines {
		if line[1] == '1' {
			lastline = i
			break
		}
	}

	for i := lastline; i >= 0; i-- {
		line := lines[i]
		x := 1
		for i := 0; i < len(cols); i++ {
			if line[x] != ' ' {
				stacks[i].Push(string(line[x]))
			}
			x += 4
		}
	}
	return stacks, lastline
}
