package day24

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
)

// Feil 488
func Day24() {
	lines := utils.GetInput("./day24/input-test2.txt")
	fmt.Println("Task1: ", task1(lines))
}

func task1(lines []string) int {
	blizzards, dim := parseInput(lines)
	//fmt.Println(dim)
	//fmt.Println(blizzards)
	start := utils.NewVec2(1, 0)
	end := utils.NewVec2(len(lines[0])-2, len(lines)-1)
	//fmt.Println(start, end)

	//dirs := []utils.Vec2{*utils.NewVec2(1,0), *utils.NewVec2(0,1),*utils.NewVec2(0,-1), *utils.NewVec2(-1,0)}
	dirs := []utils.Vec2{*utils.NewVec2(0, -1), *utils.NewVec2(0, 1), *utils.NewVec2(-1, 0), *utils.NewVec2(1, 0)}
	count := 0

	var expeditions []expedition
	var e expedition //start.Copy()
	e.pos = *start.Copy()
	e.count = 0
	e.check = true
	e.checkMin = 0
	expeditions = append(expeditions, e)

	var final expedition
	found := false

	min := 0
	for !found {
		count++

		//fmt.Scanln()

		for i := 0; i < len(blizzards); i++ {
			b := &blizzards[i]
			b.pos.AddTo(b.vel)

			if b.pos.X >= dim.X-1 && b.vel.X == 1 {
				b.pos.X = 1
			} else if b.pos.X <= 0 && b.vel.X == -1 {
				b.pos.X = dim.X - 2
			}

			if b.pos.Y >= dim.Y-1 && b.vel.Y == 1 {
				b.pos.Y = 1
			} else if b.pos.Y <= 0 && b.vel.Y == -1 {
				b.pos.Y = dim.Y - 2
			}

		}

		expCount := len(expeditions)
		for i := 0; i < expCount; i++ {

			exp := expeditions[i]
			fmt.Scanln()
			debug(blizzards, dim, *start, *end,expeditions[i].pos )
			fmt.Println("Min", min)
			if exp.checkMin != min {
				continue
			}
			//expeditions[i].check = false
			wait := true
			for _, dir := range dirs {
				check := exp.pos.Add(dir)

				if check.Equals(*end) {
					//expeditions[i].pos = *end
					found = true
					final.count = exp.count + 1
					break
				}

				if check.X <= 0 || check.Y <= 0 || check.X >= dim.X-1 || check.Y >= dim.Y-1 {
					continue
				}

				//found := false
				for _, b := range blizzards {
					if b.pos.Equals(check) {
						//found = true
						//break
						var newExpedition expedition
						newExpedition.count = exp.count + 1
						newExpedition.pos = *check.Copy()
						newExpedition.check = true
						newExpedition.checkMin = exp.checkMin + 1

						expeditions = append(expeditions, newExpedition)
						wait = false
					}
				}
				if !found {
					//exp.pos = *check.Copy()
					var newExpedition expedition
					newExpedition.count = exp.count + 1
					newExpedition.pos = *check.Copy()
					newExpedition.check = true
					newExpedition.checkMin = exp.checkMin + 1

					expeditions = append(expeditions, newExpedition)
					wait = false
					//break
				}

			}

			if wait {
				var newExpedition expedition
				newExpedition.count = exp.count + 1
				newExpedition.pos = *exp.pos.Copy()
				newExpedition.check = true
				newExpedition.checkMin = exp.checkMin + 1

				expeditions = append(expeditions, newExpedition)
			}
		}

		min++

	}

	fmt.Println(final)
	return count
}

func task1_2(lines []string) int {
	blizzards, dim := parseInput(lines)
	//fmt.Println(dim)
	//fmt.Println(blizzards)
	start := utils.NewVec2(1, 0)
	end := utils.NewVec2(len(lines[0])-2, len(lines)-1)
	//fmt.Println(start, end)
	var exp expedition //start.Copy()
	exp.pos = *start.Copy()

	dirs := []utils.Vec2{*utils.NewVec2(1, 0), *utils.NewVec2(0, 1), *utils.NewVec2(0, -1), *utils.NewVec2(-1, 0)}
	count := 0
	for !exp.pos.Equals(*end) {
		//debug(blizzards, dim, *start, *end, *expedition)
		count++

		//fmt.Scanln()
		for i := 0; i < len(blizzards); i++ {
			b := &blizzards[i]
			b.pos.AddTo(b.vel)

			if b.pos.X >= dim.X-1 && b.vel.X == 1 {
				b.pos.X = 1
			} else if b.pos.X <= 0 && b.vel.X == -1 {
				b.pos.X = dim.X - 2
			}

			if b.pos.Y >= dim.Y-1 && b.vel.Y == 1 {
				b.pos.Y = 1
			} else if b.pos.Y <= 0 && b.vel.Y == -1 {
				b.pos.Y = dim.Y - 2
			}

		}

		for _, dir := range dirs {
			check := exp.pos.Add(dir)

			if check.Equals(*end) {
				exp.pos = *end
				break
			}

			if check.X == 0 || check.Y == 0 || check.X == dim.X-1 || check.Y == dim.Y-1 {
				continue
			}

			found := false
			for _, b := range blizzards {
				if b.pos.Equals(check) {
					found = true
					break
				}
			}
			if !found {
				exp.pos = *check.Copy()
				break
			}

		}

	}

	return count
}

type blizzard struct {
	pos  utils.Vec2
	vel  utils.Vec2
	sign string
}

type expedition struct {
	pos      utils.Vec2
	count    int
	check    bool
	checkMin int
}

func parseInput(lines []string) ([]blizzard, utils.Vec2) {
	var blizzards []blizzard

	dim := utils.NewVec2(len(lines[0]), len(lines))

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] != '.' && lines[y][x] != '#' {
				vel := getVel(lines[y][x])
				pos := utils.NewVec2(x, y)
				var b blizzard
				b.pos = *pos
				b.vel = vel
				b.sign = string(lines[y][x])
				blizzards = append(blizzards, b)
			}
		}
	}
	return blizzards, *dim
}

func getVel(sign byte) utils.Vec2 {
	vel := *utils.NewVec2(-1, 0)
	if sign == '<' {
		vel.X = -1
		vel.Y = 0
	} else if sign == '>' {
		vel.X = 1
		vel.Y = 0
	} else if sign == 'v' {
		vel.X = 0
		vel.Y = 1
	} else if sign == '^' {
		vel.X = 0
		vel.Y = -1
	}

	return vel
}

func onBorder(dim utils.Vec2, x int, y int) bool {
	if x == 0 || x == dim.X-1 || y == 0 || y == dim.Y-1 {
		return true
	}
	return false
}

func posHasBlizzard(pos utils.Vec2, blizzards []blizzard) (int, blizzard) {
	count := 0
	var blizz blizzard
	for _, b := range blizzards {
		if b.pos.Equals(pos) {
			count++
			blizz = b
		}
	}
	return count, blizz
}

func debug(blizzards []blizzard, dim utils.Vec2, start utils.Vec2, end utils.Vec2, exp utils.Vec2) {
	utils.ClearConsole()

	for y := 0; y < dim.Y; y++ {
		row := ""
		for x := 0; x < dim.X; x++ {
			if x == exp.X && y == exp.Y {
				row += "E"
				continue
			}
			if (x == start.X && y == start.Y) || (x == end.X && y == end.Y) {
				row += "."
				continue
			}
			b, blizzard := posHasBlizzard(*utils.NewVec2(x, y), blizzards)
			if onBorder(dim, x, y) {
				row += "#"
			} else if b > 0 {
				//fmt.Println("b",b)

				if b == 1 {
					row += blizzard.sign
				} else {
					row += strconv.Itoa(b) //"3"//string(rune(b))
				}
			} else {
				row += "."
			}
		}
		fmt.Println(row)
	}

}
