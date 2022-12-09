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

type Rope struct {
	hx int
	hy int
	tx int
	ty int
	tails [][]int
}

func (r* Rope) runInstruction(instruction Instruction,positions map[string]bool) {
	for i := 0; i < instruction.turns; i++ {
		r.moveHead(instruction, positions)
		//key := fmt.Sprintf("%d,%d", r.tx, r.ty)
		// string(r.tx) + "," +string(r.ty)
		positions[r.getTailKey()] = true
	}
}

func (r* Rope) getTailKey() string {
	x := r.tails[len(r.tails)-1][0]
	y := r.tails[len(r.tails)-1][1]

	return fmt.Sprintf("%d,%d", x, y)
}

func (r* Rope) moveHead(instruction Instruction, positions map[string]bool) {
	r.hx += instruction.x
	r.hy += instruction.y

	/* WORKING
	a := r.hx - r.tx
	b := r.hy - r.ty
	dist := math.Sqrt(float64(a*a+b*b))
	if dist >= 2 {
		if r.hx > r.tx {
			r.tx++
		} else if r.hx < r.tx {
			r.tx--
		}

		if r.hy > r.ty {
			r.ty++
		} else if r.hy < r.ty {
			r.ty--
		}
	}
	*/
	fx := r.hx
	fy := r.hy
	for i := 0; i < len(r.tails); i++ {
		a := fx - r.tails[i][0]
		b := fy - r.tails[i][1]
		dist := math.Sqrt(float64(a*a+b*b))

		if dist >= 2 {
			if fx > r.tails[i][0] {
				r.tails[i][0]++
			} else if fx < r.tails[i][0] {
				r.tails[i][0]--
			}
	
			if fy > r.tails[i][1] {
				r.tails[i][1]++
			} else if fy < r.tails[i][1]{
				r.tails[i][1]--
			}
		}

		fx = r.tails[i][0]
		fy = r.tails[i][1]
	}

	//debug(r, positions)
	//fmt.Println(dist)
}


// Task 2: 2545 too low
func Day9() {
	lines := utils.GetInput("./day9/input.txt")
	instructions := parseInstructions(lines) 
	//fmt.Println("Task 1: ", task1(instructions))
	fmt.Println("Task 2: ", task2(instructions))

}

func task1(instructions []Instruction) int {

	rope := new(Rope)
	rope.tails = append(rope.tails, []int{0,0})
	positions := make(map[string]bool)

	for _,ins := range instructions {
		rope.runInstruction(ins, positions)
	}
	return len(positions)
}

func task2(instructions []Instruction) int {

	rope := new(Rope)
	for i := 0; i < 9; i++ {
		rope.tails = append(rope.tails, []int{0,0})
	}
	positions := make(map[string]bool)

	for _,ins := range instructions {
		rope.runInstruction(ins, positions)
	}
	return len(positions)
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

func debug(rope *Rope, positions map[string]bool) {
	cmd := exec.Command("clear") 
	cmd.Stdout = os.Stdout
    cmd.Run()
	fmt.Println(rope)
	ox := 11
	oy := 16
	for y := 0; y < 22; y++ {
		for x := 0; x < 26; x++ {
			key := fmt.Sprintf("%d,%d", x-ox,y-oy)
			if positions[key] {
				fmt.Print("#")
			} else if rope.hx == x -ox && rope.hy == y-oy {
				fmt.Print("H")
			} else if rope.tx == x -ox && rope.ty == y-oy {
				fmt.Print("T")
			} else {
				fmt.Print(".")
			}
			
		}
		fmt.Println("")
	}
	time.Sleep(50 * time.Millisecond)
}