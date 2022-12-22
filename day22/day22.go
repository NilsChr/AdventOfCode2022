package day22

import (
	"advent-of-code-2022/utils"
	"fmt"
	"regexp"
	"strconv"
	//"time"
)

func Day22() {
	lines := utils.GetInput("./day22/input.txt")
	fmt.Println("Task1: ", task1(lines))
}

func task1(lines []string) int {
	grid, instructions := parseInput(lines)
	startPos := getStartPos(grid)
	dir := utils.Vec2{X:1, Y:0}

	path := make(map[utils.Vec2]string)

	for len(instructions) > 0 {
		ins, rot, steps := getNextInstruction(instructions)
		instructions = ins

		if rot != "NA" {
			rotateDir(&dir, rot)
			continue
		}

		
		for steps > 0 {

			path[*startPos.Copy()] = "+"
			steps--
			//utils.ClearConsole()
			//fmt.Println(dir)
			//printMap(grid, startPos,path)
			//time.Sleep(1*time.Second)
			nextPos := startPos.Copy().Add(dir)
			outofBounds := nextPos.X < 0 ||nextPos.Y < 0 || nextPos.X > len(grid[0])-1 || nextPos.Y > len(grid)-1 

			//fmt.Println(nextPos, outofBounds, dir)

			if outofBounds || grid[nextPos.Y][nextPos.X] == " "{
				shouldBreak := false // If hits a wall during wrap
				// Wrap Horizontal 
				if dir.X != 0 {
					//fmt.Println("WRAPPING X")

					if dir.X == 1 {
						//fmt.Println("WRAPPING X len:", len(grid[nextPos.Y])-1)

						for tx := 0; tx < len(grid[nextPos.Y])-1; tx++ {
							if grid[nextPos.Y][tx] == "#" {
								shouldBreak = true
								break
							}
							if grid[nextPos.Y][tx] == "." {
								nextPos.X = tx
								break
							}
						}
					} else if dir.X == -1 {
						for tx := len(grid[nextPos.Y])-1; tx > 0; tx-- {
							if grid[nextPos.Y][tx] == "#" {
								shouldBreak = true
								break
							}
							if grid[nextPos.Y][tx] == "." {
								nextPos.X = tx
								break
							}
						}
					}
				}

				if dir.Y != 0 {
					//fmt.Println("WRAP Y")

					if dir.Y == 1 {
						//fmt.Println("WRAP Y 1 ")

						for ty := 0; ty < len(grid)-1; ty++ {
							if grid[ty][nextPos.X] == "#" {
								shouldBreak = true
								break
							}
							if grid[ty][nextPos.X] == "." {
								nextPos.Y = ty
								break
							}
						}
					} else if dir.Y == -1 {
						//fmt.Println("WRAP Y -1")

						for ty := len(grid)-1; ty > 0; ty-- {
							if grid[ty][nextPos.X] == "#" {
								shouldBreak = true
								break
							}
							if grid[ty][nextPos.X] == "." {
								nextPos.Y = ty
								break
							}
						}
					}
				}

				if shouldBreak {
					//fmt.Println("breaking")
					break
				}
			}


			//fmt.Println("Checking", nextPos)

			if grid[nextPos.Y][nextPos.X] == "." {
				startPos = *nextPos.Copy()
				continue
			}
			if grid[nextPos.Y][nextPos.X] == "#" {
				break
			}
		}

	}
	//printMap(grid, startPos,path)
	fmt.Println(startPos)
	fmt.Println(dir)
	startPos.AddTo(*utils.NewVec2(1,1))
	sum := (startPos.Y * 1000) + (startPos.X * 4) + getFacingValue(dir)

	return sum
}

func parseInput(lines []string) ([][]string, string){
	instructions := lines[len(lines)-1]
	lines = lines[:len(lines)-2]
	maxX := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] != ' ' && x >= maxX {
				maxX = x
			}
		}
	}
	grid := make([][]string, 0)
	for y := 0; y < len(lines); y++ {
		var row []string
		for x := 0; x <= maxX; x++ {
			if x > len(lines[y]) -1 {
				row = append(row, " ")
			} else {
				row = append(row, string(lines[y][x]))
			}
		}
		grid = append(grid, row)
	}
	return grid, instructions
}

func printMap(grid [][]string, currentPos utils.Vec2, path map[utils.Vec2]string) {
	for y:= 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {

			pos := utils.Vec2{X:x, Y:y}
			val, ok := path[pos]
			if ok && !(currentPos.X == x && currentPos.Y == y)  {
				fmt.Print(val)
			} else if currentPos.X == x && currentPos.Y == y {
				fmt.Print("o")
			} else {
				fmt.Print(grid[y][x])
			}
			
		}
		fmt.Println()
	}
}

func getStartPos(grid [][]string) utils.Vec2 {
	var pos utils.Vec2
	for x := 0; x < len(grid[0]); x++ {
		if grid[0][x] == "." {
			return utils.Vec2{X:x, Y:0}
		}
	}
	return pos
}

func rotateDir(dir *utils.Vec2, rot string) {
	x := dir.X
	y := dir.Y
	if rot == "R" {
		dir.X = -y
		dir.Y = x
	} else if rot == "L" {
		dir.X = y
		dir.Y = -x
	}
}

func getNextInstruction(instructions string) (string, string, int) {
	first := string(instructions[0])
	if first == "L" || first == "R" {
		instructions = instructions[1:]
		return instructions, first, -1
	}
	var re = regexp.MustCompile(`(\d+)`)
	matches := re.FindAllString(instructions, -1)
	instructions = instructions[len(matches[0]):]
	next,_ := strconv.Atoi(matches[0])
	return instructions,"NA" ,next
}

func getFacingValue(dir utils.Vec2) int {
	if dir.X == 1 && dir.Y == 0 {
		return 0
	}
	if dir.X == 0 && dir.Y == 1 {
		return 1
	}
	if dir.X == -1 && dir.Y == 0 {
		return 2
	}
	if dir.X == 0 && dir.Y == -1 {
		return 3
	}
	return 0
}