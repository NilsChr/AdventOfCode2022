package day3

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
)

func Day3() {
	fmt.Println("Day 3")
	lines := utils.GetInput("./day3/input.txt")
	fmt.Println("Task1: ", task1(lines))
	fmt.Println("Task2: ", task2(lines))
}

func task1(lines []string) int {
	sum := 0
	for _, line := range lines {
		half := len(line)/2
		sum += getPriority(getCommon([]string{line[:half], line[half:]}))
	}
	return sum
}

func task2(lines []string) int {
	sum := 0
	for i := 0; i < len(lines); i += 3 {
		sum += getPriority(getCommon([]string{lines[i], lines[i+1], lines[i+2]}))
	}
	return sum
}

func getPriority(search string) int {
	return strings.Index("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", search) + 1
}

func getCommon(parts []string) string {
	for i := range parts[0] {
		currentChar := string(parts[0][i])
		occurance := 1
		for _, value := range parts[1:] {
			if strings.Contains(value, currentChar) {
				occurance++
				if occurance == len(parts) {
					return currentChar
				}
			}
		}
	}
	return ""
}