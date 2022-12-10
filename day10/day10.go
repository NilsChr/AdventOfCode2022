package day10

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day10() {
	lines := utils.GetInput("./day10/input.txt")
	fmt.Println("Task1: ", task1(lines))
}

func task1(lines []string) int {
	cpu := new(Cpu)
	cpu.regX = 1
	cpu.signalListeners = []int{20,60,100,140,180,220}

	for _, line := range lines {
		cpu.runInstruction(line)
	}

	cpu.render()
	return cpu.regScore
}

type Cpu struct {
	regX int
	regScore int
	cycle int
	signalListeners []int
	pixelIndex int
	screenBuffer string
	screen []string
}

func (c* Cpu) runInstruction(input string) {
	parts := strings.Split(input, " ")
	instruction := parts[0]
	if instruction == "noop" {
		c.increment()
	} else if instruction == "addx" {
		c.increment()
		value,_ := strconv.Atoi(parts[1])
		c.increment()
		c.regX += value
	}
}

func (c* Cpu) increment() {
	c.cycle++
	if utils.Contains(c.signalListeners, c.cycle) {
		c.regScore += c.cycle * c.regX
	}
	c.addToScreenBuffer()
}

func (c* Cpu) addToScreenBuffer() {
	spritePos := c.regX
	if c.pixelIndex < spritePos -1 || c.pixelIndex > spritePos +1 {
		c.screenBuffer += " "
	} else {
		c.screenBuffer += "█"
	}
	c.pixelIndex++
	if(c.pixelIndex >= 40) {
		c.screen = append(c.screen, c.screenBuffer)
		c.screenBuffer = ""
		c.pixelIndex = 0
	} 
}

func (c* Cpu) render() {
	for _, line := range c.screen {
		fmt.Println(line)
	}
}