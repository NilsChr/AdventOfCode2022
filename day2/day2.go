package day2

import (
	"advent-of-code-2022/utils"
	"fmt"
)

const (
	ROCK       int = 1
	PAPER          = 2
	SCISSOR        = 3
	LOOSE          = 1
	DRAW           = 2
	WIN            = 3
	DRAW_SCORE     = 3
	WIN_SCORE      = 6
)

func Day2() {
	lines := utils.GetInput("./day2/input.txt")
	task1 := runGame(lines, 1)
	fmt.Println("Task 1 score: ", task1)

	task2 := runGame(lines, 2)
	fmt.Println("Task 1 score: ", task2)
}

func runGame(lines []string, strat int) int {
	var score int = 0
	for _, line := range lines {
		if strat == 1 {
			score += calculateScore1(line)
		}
		if strat == 2 {
			score += calculateScore2(line)
		}
	}
	return score
}

func calculateScore1(line string) int {
	var handA = getValue(line[0])
	var handB = getValue(line[2])

	if handB == handA {
		return DRAW_SCORE + handB
	}
	if handB == ROCK && handA == SCISSOR {
		return WIN_SCORE + handB
	}
	if handB == PAPER && handA == ROCK {
		return WIN_SCORE + handB
	}
	if handB == SCISSOR && handA == PAPER {
		return WIN_SCORE + handB
	}

	return handB
}

func calculateScore2(line string) int {
	var hand = getValue(line[0])
	var condition = getValue(line[2])

	if condition == LOOSE {
		return getScore(hand, condition)
	}
	if condition == DRAW {
		return DRAW_SCORE + getScore(hand, condition)
	}
	if condition == WIN {
		return WIN_SCORE + getScore(hand, condition)
	}

	return -1
}

func getValue(char byte) int {
	if char == 'A' || char == 'X' {
		return 1
	}
	if char == 'B' || char == 'Y' {
		return 2
	}
	if char == 'C' || char == 'Z' {
		return 3
	}
	return 0
}

func getScore(score int, condition int) int {
	var out int = score
	if condition == DRAW {
		return out
	}
	if condition == LOOSE {
		out -= 1
	}
	if condition == WIN {
		out += 1
	}

	if out > 3 {
		out = out % 3
	}
	if out < 1 {
		out = 3
	}
	return out
}
