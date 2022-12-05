package day5

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
	"strconv"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop()(string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Peak() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		return (*s)[len(*s) - 1], true
	}
}

func Day5() {
	fmt.Println("Day 5")
	lines := utils.GetInput("./day5/input.txt")
	fmt.Println("Task1: ", task1(lines))
	fmt.Println("Task2: ", task2(lines))
}

func task1(lines []string) string {
	stacks := initStacks()
	for _, line := range lines {
		amount, from, to := parseInstructions(line)
		for i := 0; i < amount; i++ {
			data, found := stacks[from].Pop()
			if found {
				stacks[to].Push(data)
			}
		}
	}
	out := ""
	for _, stack := range stacks {
		data, exists := stack.Peak()
		if exists {
			out += data
		}
	}
	return out
}

func task2(lines []string) string {
	stacks := initStacks()
	for _, line := range lines {
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
	out := ""
	for _, stack := range stacks {
		data, exists := stack.Peak()
		if exists {
			out += data
		}
	}
	return out
}

func parseInstructions(line string) (int, int, int) {
	parts := strings.Split(line, " ")
	amount,_ := strconv.Atoi(parts[1])
	from,_ := strconv.Atoi(parts[3])
	to,_ := strconv.Atoi(parts[5])
	return amount, from-1, to-1
}

func initStacks() [9]Stack {
	var stacks [9]Stack
	initStack(&stacks[0], "FDBZTJRN")
	initStack(&stacks[1], "RSNJH")
	initStack(&stacks[2], "CRNJGZFQ")
	initStack(&stacks[3], "FVNGRTQ")
	initStack(&stacks[4], "LTQF")
	initStack(&stacks[5], "QCWZBRGN")
	initStack(&stacks[6], "FCLSNHM")
	initStack(&stacks[7], "DNQMTJ")
	initStack(&stacks[8], "PGS")
	return stacks
}

func initStack(stack *Stack, input string) {
	for _, i := range input {
		stack.Push(string(i))
	}
}