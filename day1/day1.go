package day1

import (
	"advent-of-code-2022/utils"
	"fmt"
	"sort"
	"strconv"
)

type elf struct {
	calories int
}

func Day1() {
	fmt.Println("Day 1")
	lines := utils.GetInput("./day1/input.txt")
	elves := GenerateElves(lines)
	calories1 := GetCalorieCount(elves,1)
	calories2 := GetCalorieCount(elves,3)

	fmt.Println("Tast 1", calories1)
	fmt.Println("Tast 1", calories2)

}

func GenerateElves(lines []string) []elf {
	var elfes []elf = []elf{}
	var currentElf *elf = &elf{}
	
	for _, line := range lines {
		if line == "" {
			elfes = append(elfes, *currentElf)
			currentElf = &elf{}
			continue
		} 
		cal, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Could not convert")
		}
		currentElf.calories += cal
	}

	return elfes
}

func GetCalorieCount(elfes []elf, count int) int {
	sort.Slice(elfes, func(i,j int) bool {
		return elfes[i].calories > elfes[j].calories
	})
	sum := 0
	for i := 0; i < count; i++ {
		sum += elfes[i].calories
	}
	return sum
} 