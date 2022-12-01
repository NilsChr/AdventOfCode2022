package day1

import (
	"advent-of-code-2022/utils"
	"testing"
)

func Test1(t *testing.T) {
	lines := utils.GetInput("./input-test.txt")
	if len(lines) != 14 {
		t.Errorf("Number of lines was incorrect, got: %d, want: %d.", len(lines), 14)
	}

	elfes := GenerateElves(lines)
	if len(elfes) != 4 {
		t.Errorf("Number of elfes was incorrect, got: %d, want: %d.", len(elfes), 4)
	}

	highestCalorie := GetCalorieCount(elfes, 1)
	if highestCalorie != 24000 {
		t.Errorf("Highest calorie was incorrect, got: %d, want: %d.", highestCalorie, 24000)
	}
}