package day9

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
	"math"
)

type Instruction struct{
	x int
	y int
	turns int
}

type Vec2 struct {
	x int 
	y int
}

func Day9() {
	lines := utils.GetInput("./day9/input.txt")
	instructions := parseInstructions(lines) 
	fmt.Println("Task 1: ", task(instructions,2))
	fmt.Println("Task 2: ", task(instructions,10))
}

func task(instructions []Instruction, ropeSize int) int {
	rope := createRope(ropeSize)
	positions := make(map[Vec2]bool)

	for _,ins := range instructions {
		for i := 0; i < ins.turns; i++ {
			moveRope(rope, ins)
			tail := rope[len(rope)-1]
			positions[tail] = true
		}
	}

	return len(positions)
}

func createRope(length int) []Vec2 {
	var rope []Vec2 = []Vec2{}
	for i := 0; i < length; i++ {
		rope = append(rope, *new(Vec2))
	}
	return rope
}

func moveRope(rope []Vec2,instruction Instruction) {
	rope[0].x += instruction.x
	rope[0].y += instruction.y

	fx := rope[0].x
	fy := rope[0].y
	for i := 0; i < len(rope); i++ {
		a := fx - rope[i].x
		b := fy - rope[i].y
		dist := math.Sqrt(float64(a*a+b*b))

		if dist >= 2 {
			if fx > rope[i].x {
				rope[i].x++
			} else if fx < rope[i].x {
				rope[i].x--
			}
	
			if fy > rope[i].y {
				rope[i].y++
			} else if fy < rope[i].y{
				rope[i].y--
			}
		}

		fx = rope[i].x
		fy = rope[i].y
	}
}

func parseInstructions(lines []string) []Instruction {
	var instructions []Instruction 
	for _, line := range lines {
		parts := strings.Split(line, " ")
		dir := parts[0]
		turns,_ := strconv.Atoi(parts[1])
		ins := new(Instruction)
		ins.turns = turns
		switch dir {
		case "R":
			ins.x = 1
			ins.y = 0
		case "L":
			ins.x = -1
			ins.y = 0
		case "U":
			ins.x = 0
			ins.y = -1
		case "D":
			ins.x = 0
			ins.y = 1
		}
		instructions = append(instructions, *ins)	
	}

	return instructions
}