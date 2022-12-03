package day3

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
)

func Day3() {
	fmt.Println("Day 3")
	lines := utils.GetInput("./day3/input.txt")

	sum1 := task1(lines)
	fmt.Println("Task1: ", sum1)

	sum2 := task2(lines)
	fmt.Println("Task2: ", sum2)
}

func task1(lines []string) int {
	sum := 0
	for _, line := range lines {
		part1, part2 := getParts(line)
		common := getCommonItem(part1, part2)
		priority := getPriority(common)
		sum += priority
	}
	return sum
}

func task2(lines []string) int {
	sum := 0
	for i := 0; i < len(lines); i += 3 {
		common := getCommonItem3(lines[i], lines[i+1], lines[i+2])
		sum += getPriority(common)
	}
	return sum
}

func getParts(line string) (string, string) {
	part1 := line[:len(line)/2]
	part2 := line[len(line)/2:]
	return part1, part2
}

func getPriority(search string) int {
	if search == "" {
		return 0
	}
	priorityValues := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(priorityValues, search) + 1
}

func getCommonItem(part1 string, part2 string) string {
	for _, i := range part1 {
		for _, j := range part2 {
			if i == j {
				return string(j)
			}
		}
	}
	return ""
}
func getCommonItem3(part1 string, part2 string, part3 string) string {
	for _, i := range part1 {
		for _, j := range part2 {
			if i != j {
				continue
			}
			for _, k := range part3 {
				if j == k {
					return string(k)
				}
			}
		}
	}
	return ""
}
