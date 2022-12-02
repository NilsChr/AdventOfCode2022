package day2

import (
	"fmt"
	"advent-of-code-2022/utils"
)

func Day2() {
	lines := utils.GetInput("./day2/input.txt")
	task1 := runGame(lines, 1)
	fmt.Println("Task 1 score: ", task1)
}

func runGame(lines []string, strat int) int {
	var score int = 0;
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
	var otherScore = getValue(line[0])
	var myScore = getValue(line[2])

	// 1 rock, 2 paper, 3 scissor
	if myScore == otherScore {
		return 3 + myScore
	}
	if myScore == 1 && otherScore == 3 {
		return 6 + myScore
	}
	if myScore == 2 && otherScore == 1 {
		return 6 + myScore
	}
	if myScore == 3 && otherScore == 2 {
		return 6 + myScore
	}

	return myScore
}

func calculateScore2(line string) int {
	var otherScore = getValue(line[0])
	var myScore = getValue(line[2])

	// 1 rock, 2 paper, 3 scissor
	// 1 loose, 2 draw, 3 win
	if myScore == 1 {
		return getLoose(byte(otherScore))
	}
	if myScore == 2 {
		return 3 + getDraw(byte(otherScore))
	}
	if myScore == 3 {
		return + getWin(byte(otherScore))
	}

	return myScore
}

func getValue(char byte) int {
	if char == 'A' || char == 'X' {
		return 1;
	}
	if char == 'B' || char == 'Y' {
		return 2;
	}
	if char == 'C' || char == 'Z' {
		return 3;
	}
	return 0
}

func getWin(char byte) int {
	if char == 'A' {
		return 2;
	}
	if char == 'B' {
		return 3;
	}
	if char == 'C' {
		return 1;
	}
	return 0
}
func getLoose(char byte) int {
	if char == 'A' {
		return 3;
	}
	if char == 'B' {
		return 1;
	}
	if char == 'C' {
		return 2;
	}
	return 0
}
func getDraw(char byte) int {
	if char == 'A' {
		return 1;
	}
	if char == 'B' {
		return 2;
	}
	if char == 'C' {
		return 3;
	}
	return 0
}
