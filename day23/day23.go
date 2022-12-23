package day23

import (
	"advent-of-code-2022/utils"
	"fmt"
	"time"
)

var DIRS = []utils.Vec2{*utils.NewVec2(0, -1), *utils.NewVec2(0, 1), *utils.NewVec2(-1, 0), *utils.NewVec2(1, 0)}
var CHECKS = make(map[utils.Vec2][]utils.Vec2)

func Day23() {
	CHECKS[*utils.NewVec2(0, -1)] = []utils.Vec2{*utils.NewVec2(-1, -1), *utils.NewVec2(0, -1), *utils.NewVec2(1, -1)}
	CHECKS[*utils.NewVec2(0, 1)] = []utils.Vec2{*utils.NewVec2(-1, 1), *utils.NewVec2(0, 1), *utils.NewVec2(1, 1)}
	CHECKS[*utils.NewVec2(-1, 0)] = []utils.Vec2{*utils.NewVec2(-1, -1), *utils.NewVec2(-1, 0), *utils.NewVec2(-1, 1)}
	CHECKS[*utils.NewVec2(1, 0)] = []utils.Vec2{*utils.NewVec2(1, -1), *utils.NewVec2(1, 0), *utils.NewVec2(1, 1)}

	lines := utils.GetInput("./day23/input.txt")
	fmt.Println("Task1: ", task1(lines))
	fmt.Println("Task2: ", task2(lines))

}

func task1(lines []string) int {
	elfes, dim := parseInput(lines)

	fmt.Println(dim)
	dir := 0

	for i := 0; i < 10; i++ {
		uniqueMoves := make(map[utils.Vec2]int)

		for i := 0; i < len(elfes); i++ {
			elfes[i].considerMove(elfes, dir)
			if elfes[i].shouldMove {
				uniqueMoves[elfes[i].proposedPos]++
			}
		}
		for i := 0; i < len(elfes); i++ {
			if uniqueMoves[elfes[i].proposedPos] == 1 && elfes[i].shouldMove {
				elfes[i].pos = *elfes[i].proposedPos.Copy()

			}
		}

		dir = (dir + 1) % 4
	}

	minY := 999
	minX := 999
	maxY := 0
	maxX := 0

	for _, elf := range elfes {
		if elf.pos.X < minX {
			minX = elf.pos.X 
		} 
		if elf.pos.X > maxX {
			maxX = elf.pos.X 
		}
		if elf.pos.Y < minY {
			minY = elf.pos.Y
		} 
		if elf.pos.Y > maxY {
			maxY = elf.pos.Y 
		}
	}

	sum := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if !positionHasElf(elfes, *utils.NewVec2(x,y)) {
				sum++
			}
		}
	}

	return sum
}

func task2(lines []string) int {
	elfes, dim := parseInput(lines)

	fmt.Println(dim)
	dir := 0
	i := 0
	for {
		i++

		uniqueMoves := make(map[utils.Vec2]int)

		for i := 0; i < len(elfes); i++ {
			elfes[i].considerMove(elfes, dir)
			if elfes[i].shouldMove {
				uniqueMoves[elfes[i].proposedPos]++
			}
		}
		nonmoving := 0
		for i := 0; i < len(elfes); i++ {
			if !elfes[i].shouldMove {
				nonmoving++
			}
			if uniqueMoves[elfes[i].proposedPos] == 1 && elfes[i].shouldMove {
				elfes[i].pos = *elfes[i].proposedPos.Copy()

			}
		}

		if nonmoving == len(elfes) {
			return i
		}

		dir = (dir + 1) % 4
	}
}

func parseInput(lines []string) ([]elf, *utils.Vec2) {
	dim := utils.NewVec2(len(lines[0]), len(lines))

	var elfes []elf
	for y := 0; y < len(lines); y++ {
		for x := 0; x <= len(lines[y])-1; x++ {
			if lines[y][x] == '.' {
				continue
			}
			elf := &elf{}
			elf.pos = *utils.NewVec2(x, y)
			elf.checkIndex = 0
			elf.proposedPos = *utils.NewVec2(0, 0)
			elfes = append(elfes, *elf)
		}
	}
	return elfes, dim
}

type elf struct {
	pos         utils.Vec2
	proposedPos utils.Vec2
	checkIndex  int
	shouldMove  bool
}

func (e *elf) considerMove(elfes []elf, firstDir int) {
	e.shouldMove = false

	hasNeighbour := false
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if y == 0 && x == 0 {
				continue
			}
			offset := utils.NewVec2(x, y)
			if positionHasElf(elfes, e.pos.Add(*offset)) {
				hasNeighbour = true
				break
			}
		}
		if hasNeighbour {
			break
		}
	}

	if !hasNeighbour {
		return
	}

	var proposed utils.Vec2
	found := false
	checks := 0

	for !found {
		canMove := true
		for _, pos := range CHECKS[DIRS[firstDir]] {
			if positionHasElf(elfes, e.pos.Add(pos)) {
				canMove = false
			}
		}
		if canMove {
			found = true
			proposed = *DIRS[firstDir].Copy()
			break
		}

		firstDir = (firstDir + 1) % 4

		checks++
		if checks > 3 {
			break
		}
	}
	if found {
		e.shouldMove = true
		e.proposedPos = proposed.Add(e.pos)
	}

}

func positionHasElf(elfes []elf, pos utils.Vec2) bool {
	for _, elf := range elfes {
		if elf.pos.Equals(pos) {
			return true
		}
	}

	return false
}

func debug(elfes []elf, dim utils.Vec2, round int) {
	time.Sleep(250 * time.Millisecond)
	utils.ClearConsole()
	fmt.Println("DEBUG: ", round)

	for y := -dim.Y; y < dim.Y*2; y++ {
		row := ""
		for x := -dim.X; x < dim.X*2; x++ {
			pos := utils.NewVec2(x, y)
			found := false
			for _, elf := range elfes {
				if elf.pos.Equals(*pos) {
					row += "#"
					found = true
				}
			}
			if !found {
				row += "."
			}
		}
		fmt.Println(row)
	}

}
