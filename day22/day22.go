package day22

import (
	"advent-of-code-2022/utils"
	"fmt"
	"time"

	//"math"
	"regexp"
	"strconv"
	//"time"
)

func Day22() {
	lines := utils.GetInput("./day22/input-test4.txt")
//	fmt.Println("Task1: ", task1(lines))
	fmt.Println("Task2: ", task2(lines))

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

func task2(lines []string) int {
	grid, instructions := parseInput(lines)
	startPos := getStartPos(grid)
	dir := utils.Vec2{X:1, Y:0}
	cubeSize := 2
	fmt.Print(instructions)
	path := make(map[utils.Vec2]string)

	for len(instructions) > 0 {
		ins, rot, steps := getNextInstruction(instructions)
		instructions = ins

		if rot != "NA" {
			rotateDir(&dir, rot)
			continue
		}

		
		for steps > 0 {

			utils.ClearConsole()
			fmt.Println(dir)
			printMap(grid, startPos,path)
			time.Sleep(1*time.Second)

			path[*startPos.Copy()] = "+"
			steps--
			nextPos := startPos.Copy().Add(dir)
			outofBounds := nextPos.X < 0 ||nextPos.Y < 0 || nextPos.X > len(grid[0])-1 || nextPos.Y > len(grid)-1 

			if grid[nextPos.Y][nextPos.X] == " " {
				outofBounds = true
			}

			if outofBounds || grid[nextPos.Y][nextPos.X] == " "{
				
				shouldBreak := false // If hits a wall during wrap

				x, y := moveCoordinateOnCubeTexture(startPos.X, startPos.Y,cubeSize, nextPos.X, nextPos.Y)
				fmt.Println(x,y)


				/*
				currentSide := getCurrentSide(cubeSize, &startPos)
				//nextSide := getCurrentSide(cubeSize, &nextPos)

				if currentSide == 1 && nextPos.X < startPos.X {
					// Wrap to side '4
					ny := cubeSize * 2 + (cubeSize - (startPos.Y % cubeSize))
					wrapPos := utils.NewVec2(0,ny)
					// 2 * 2 + (2 - 0 % 2)
					if grid[wrapPos.Y][wrapPos.X] == "#" {
						shouldBreak = true
					} else if grid[wrapPos.Y][wrapPos.X] == "." {
						nextPos = *wrapPos.Copy()
					}
				}
				*/

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

func getCurrentSide(size int, pos *utils.Vec2) int {
	//int gridCubeWidth = 16, gridCubeHeight = 16;

	//cube.Position.X = Math.round(cube.Position.X / gridCubeWidth) * gridCubeWidth;
	//cube.Position.Y = Math.round(cube.Position.Y / gridCubeHeight) * gridCubeHeight;
	x := (pos.X / size * size) / size //math.Floor(float64(pos.X / size)) * float64(size)
	y := (pos.Y / size * size) / size //math.Floor(1.1) * float64(size)
	if x == 1 && y == 0 {
		return 1
	}
	if x == 2 && y == 0 {
		return 2
	}
	if x == 1 && y == 1 {
		return 3
	}
	if x == 0 && y == 2 {
		return 4
	}
	if x == 1 && y == 2 {
		return 5
	}
	if x == 0 && y == 3 {
		return 6
	}

	return -1
}


// Here comes ChatGPT
func mapCoordinateToTexture(x, y, z int) (int, int) {
	// Determine which face of the cube the coordinate is on
	if x == 50 {
		// Coordinate is on the right face of the cube
		u := 50 - z
		v := y
		return u, v
	} else if x == 0 {
		// Coordinate is on the left face of the cube
		u := z
		v := y
		return u, v
	} else if y == 50 {
		// Coordinate is on the top face of the cube
		u := x
		v := z
		return u, v
	} else if y == 0 {
		// Coordinate is on the bottom face of the cube
		u := x
		v := 50 - z
		return u, v
	} else if z == 50 {
		// Coordinate is on the front face of the cube
		u := x
		v := y
		return u, v
	} else if z == 0 {
		// Coordinate is on the back face of the cube
		u := 50 - x
		v := y
		return u, v
	} else {
		// Coordinate is inside the cube, not on any face
		return 0, 0
	}
}

func moveCoordinateOnCubeTexture(x, y, size, dx, dy int) (int, int) {
	// Convert the 2D coordinate to a 3D coordinate
	z := 0
	if x > size/2 {
		x = size - x
		z = size
	}
	if y > size/2 {
		y = size - y
		z = size
	}
	// Map the 3D coordinate to a position on the texture
	u, v := mapCoordinateToTexture(x, y, z)
	// Add the displacement to the texture coordinates
	u += dx
	v += dy
	// Wrap the texture coordinates around if they go out of bounds
	u = u % size
	v = v % size
	// Convert the texture coordinates back to a 2D coordinate
	if z == size {
		x = size - u
	} else {
		x = u
	}
	if z == size {
		y = size - v
	} else {
		y = v
	}
	return x, y
}