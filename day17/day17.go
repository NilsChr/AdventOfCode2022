package day17

import (
	u "advent-of-code-2022/utils"
	"fmt"
)

func Day17() {
	lines := u.GetInput("./day17/input-test.txt")
	fmt.Println("Task1:", sim(lines,2022 ))
	//fmt.Println("Task2:", sim(lines,1000000000000 ))

}

func sim(lines []string, runs int) int {
	//fmt.Println(lines)
	shapes := createShapes()

	levelHeight := 5000
	level := make([][]int, levelHeight)
	for i := 0; i < len(level); i++ {
		level[i] = []int{0, 0, 0, 0, 0, 0, 0}
	}

	rockIndex := 0
	towerHeight := 0
	currentInstruction := -1
	fmt.Println("Started")
	for i := 0; i < runs; i++ {
		if i % 1000 == 0 {
			fmt.Printf("%d/%d\n",i,runs)
		}
		//nextY := towerHeight + rock.height + 2
		//fmt.Println("next: ", shapes[rockIndex], rockIndex)
		//fmt.Scanln()
		rock := shapes[rockIndex]
		nextY := levelHeight - (towerHeight + rock.height + 3)
		//fmt.Println("Y: ",  (towerHeight + rock.height + 3))
		//fmt.Scanln()
		rock.pos.Y = nextY
		rock.pos.X = 2

		moveComplete := false

		for !moveComplete {
			currentInstruction++
			if currentInstruction > len(lines[0]) -1 {
				currentInstruction = 0
				return towerHeight * 8
			}
			dirX := 0
			
			//debug(level, rock, towerHeight)
			//fmt.Println("Next move: ", string(lines[0][currentInstruction]))

			//fmt.Scanln()

			if lines[0][currentInstruction] == '<' {
				dirX = -1
			} else if lines[0][currentInstruction] == '>' {
				dirX = 1
			}
			canMove := rock.canMove(dirX, 0, level)
			if canMove {
				rock.pos.X += dirX
			}

			//debug(level, rock, towerHeight)
			//fmt.Scanln()

			canMoveY := rock.canMove(0, 1, level)
			if canMoveY {
				rock.pos.Y += 1
			} else {
				for _, pos := range rock.checks {
					y := rock.pos.Y + pos.Y
					x := rock.pos.X + pos.X
					level[y][x] = 1

				}
				towerHeight += rock.height
				moveComplete = true
				rockIndex++
				if rockIndex > len(shapes)-1 {
					rockIndex = 0
				}
				towerHeight = getTowerHeight(level)

			}
			//debug(level, rock, towerHeight)
			//fmt.Scanln()
		}

	}

	return towerHeight
}


func getTowerHeight(level [][]int) int {
	height := 0
	for y := len(level)-1; y >= 0; y-- {

		for x := len(level[y])-1; x >= 0; x-- {
			if level[y][x] == 1 {
				height++
				break
			}
		}
	}
	return height
}

type shape struct {
	name   string
	pos    u.Vec2
	checks []u.Vec2
	width  int
	height int
}

func (s *shape) canMove(dirX int, dirY int, level [][]int) bool {
	//fmt.Println("x,y:",dirX,dirY)
	//fmt.Println("sx + (w):",s.pos.X, s.pos.X+dirX+s.width)

	if s.pos.Y+dirY+s.height > len(level) {
		return false
	}
	if (s.pos.X+dirX)+s.width > 7 || s.pos.X+dirX < 0 {
		return false
	}
	canMove := true
	for _, pos := range s.checks {
		//fmt.Println((s.pos.X+pos.X)+dirX,(s.pos.Y + pos.Y)+dirY)

		if level[(s.pos.Y + pos.Y)+dirY][(s.pos.X+pos.X)+dirX] != 0 {
			canMove = false
		}
	}
	return canMove
}

func createShapes() []shape {
	lineH := &shape{"lineH", *u.NewVec2(0, 0), []u.Vec2{*u.NewVec2(0, 0), *u.NewVec2(1, 0), *u.NewVec2(2, 0), *u.NewVec2(3, 0)}, 4, 1}
	star := &shape{"star", *u.NewVec2(0, 0), []u.Vec2{*u.NewVec2(1, 0),*u.NewVec2(1, 1), *u.NewVec2(1, 2), *u.NewVec2(0, 1), *u.NewVec2(2, 1)}, 3, 3}
	revL := &shape{"revL", *u.NewVec2(0, 0), []u.Vec2{*u.NewVec2(2, 0), *u.NewVec2(2, 1), *u.NewVec2(2, 2), *u.NewVec2(1, 2), *u.NewVec2(0, 2)}, 3, 3}
	lineV := &shape{"liveV", *u.NewVec2(0, 0), []u.Vec2{*u.NewVec2(0, 0), *u.NewVec2(0, 1), *u.NewVec2(0, 2), *u.NewVec2(0, 3)}, 1, 4}
	box := &shape{"box", *u.NewVec2(0, 0), []u.Vec2{*u.NewVec2(0, 0), *u.NewVec2(1, 0), *u.NewVec2(0, 1), *u.NewVec2(1, 1)}, 2, 2}
	return []shape{*lineH, *star, *revL, *lineV, *box}
}

func debug(level [][]int, current shape, towerheight int) {
	renderHeight := towerheight
	if towerheight <= 10 {
		renderHeight = 10
	}
	u.ClearConsole()
	fmt.Println("Tower height: ", towerheight)
	fmt.Println("Y: ", current.pos.Y)
	fmt.Println("Cur: ", current)
	fmt.Println("Height: ",current.height )

	for y := len(level) - renderHeight; y < len(level); y++ {
		line := "|"
		for x, c := range level[y] {

			rendered := false
			for _, pos := range current.checks {
				if current.pos.X+pos.X == x && current.pos.Y+pos.Y == y {
					line += "@"
					rendered = true
				}
			}

			if !rendered && c == 0 {
				line += "."
			} else if !rendered && c == 1 {
				line += "#"
			}
		}
		line += "|"
		fmt.Println(line)
	}
	fmt.Println("+-------+")

}
