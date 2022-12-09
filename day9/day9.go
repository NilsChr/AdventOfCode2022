package day9

import (
	"advent-of-code-2022/utils"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"math"
)

type Instruction struct{
	x int
	y int
	turns int
}

func Day9() {
	lines := utils.GetInput("./day9/input.txt")
	instructions := parseInstructions(lines) 
	fmt.Println("Task 1: ", task(instructions,2))
	fmt.Println("Task 2: ", task(instructions,10))
}

func task(instructions []Instruction, ropeSize int) int {
	rope := createRope(ropeSize)
	positions := make(map[string]bool)

	for _,ins := range instructions {
		for i := 0; i < ins.turns; i++ {
			moveRope(rope, ins, positions)
			tail := rope[len(rope)-1]
			positions[fmt.Sprintf("%d,%d", tail[0], tail[1])] = true
		}
	}

	return len(positions)
}

func createRope(length int) [][]int {
	var rope [][]int = [][]int{}
	for i := 0; i < length; i++ {
		rope = append(rope, []int{0,0})
	}
	return rope
}

func moveRope(rope [][]int,instruction Instruction, positions map[string]bool) {
	rope[0][0] += instruction.x
	rope[0][1] += instruction.y

	fx := rope[0][0]
	fy := rope[0][1]
	for i := 0; i < len(rope); i++ {
		a := fx - rope[i][0]
		b := fy - rope[i][1]
		dist := math.Sqrt(float64(a*a+b*b))

		if dist >= 2 {
			if fx > rope[i][0] {
				rope[i][0]++
			} else if fx < rope[i][0] {
				rope[i][0]--
			}
	
			if fy > rope[i][1] {
				rope[i][1]++
			} else if fy < rope[i][1]{
				rope[i][1]--
			}
		}

		fx = rope[i][0]
		fy = rope[i][1]
	}

	//debug(rope, positions)
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

func debug(rope [][]int, positions map[string]bool) {
	cmd := exec.Command("clear") 
	cmd.Stdout = os.Stdout
    cmd.Run()
	fmt.Println(rope)
	ox := 11
	oy := 16
	for y := 0; y < 22; y++ {
		for x := 0; x < 26; x++ {
			key := fmt.Sprintf("%d,%d", x-ox,y-oy)
			rx := rope[0][0]
			ry := rope[0][1]

			if positions[key] {
				fmt.Print("#")
			} else if rx == x -ox && ry == y-oy {
				fmt.Print("H")
			} else if rx == x -ox && ry == y-oy {
				fmt.Print("T")
			} else {
				fmt.Print(".")
			}
			
		}
		fmt.Println("")
	}
	time.Sleep(50 * time.Millisecond)
}