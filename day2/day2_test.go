package day2

import (
	"advent-of-code-2022/utils"
	"fmt"
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

func Test3(t *testing.T) {
	score := calculateScore2("A Y")
	expected := 4
	if score != expected {
		t.Errorf("Score is incorrect, got: %d, want: %d.", score, expected)
	}
}

func Test4(t *testing.T) {
	score := calculateScore2("B X")
	fmt.Println("Test Score,", score)
	expected := 1
	if score != expected {
		t.Errorf("Score is incorrect, got: %d, want: %d.", score, expected)
	}
}

func Test5(t *testing.T) {
	score := calculateScore2("C Z")
	fmt.Println("Test Score,", score)
	expected := 7
	if score != expected {
		t.Errorf("Score is incorrect, got: %d, want: %d.", score, expected)
	}
}
