package day2

import (
	"advent-of-code-2022/utils"
	"testing"
)

func Test1(t *testing.T) {
	lines := utils.GetInput("./input-test.txt")
	score := runGame(lines, 1)
	if score != 15 {
		t.Errorf("Score is incorrect, got: %d, want: %d.", score, 15)
	}
}

func Test2(t *testing.T) {
	lines := utils.GetInput("./input-test.txt")
	score := runGame(lines, 2)
	if score != 12 {
		t.Errorf("Score is incorrect, got: %d, want: %d.", score, 12)
	}
}