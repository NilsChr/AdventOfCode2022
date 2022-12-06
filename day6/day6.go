package day6

import (
	"advent-of-code-2022/utils"
	"fmt"
)

func Day6() {
	lines := utils.GetInput("./day6/input.txt")
	fmt.Println("Task1: ", findMarker(lines[0], 4))
	fmt.Println("Task2: ", findMarker(lines[0], 14))

}

func findMarker(line string, size int) int {
	for i := size; i < len(line); i++ {
		m := make(map[string]bool)
		for j := 0; j < size; j++ {
			m[string(line[i-j])] = true
		}
		if len(m) == size {
			return i+1;
		}
	}
	return -1
}