package day14

import (
	u "advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day14() {
	lines := u.GetInput("./day14/input.txt")
	fmt.Println("Task1: ", runSimulation(lines, false))
	fmt.Println("Task2: ", runSimulation(lines, true))
}

func runSimulation(lines []string, infiniteFloor bool) int {
	walls, maxY := parseInput(lines)

	if infiniteFloor {
		maxY += 2
		for x := -999; x < 999; x++ {
			wall := u.NewVec2(x, maxY)
			walls[*wall] = true
		}
	}

	spawnPoint := u.NewVec2(500, 0)
	var sands []u.Vec2
	var sand *u.Vec2
	sand = spawnPoint.Copy()
	for {
		if sand == nil {
			sand = spawnPoint.Copy()
		}

		newPos := sand.Copy()
		newPos.Y++

		if walls[*newPos] {
			downLeft := u.NewVec2(newPos.X-1, newPos.Y)
			downRight := u.NewVec2(newPos.X+1, newPos.Y)
			if !walls[*downLeft] {
				newPos.X--
			} else if !walls[*downRight] {
				newPos.X++
			} else {
				walls[*sand] = true
				sands = append(sands, *sand)
				if sand.Equals(*spawnPoint) {
					return len(sands)
				}
				sand = nil
				continue
			}
		}

		sand.X = newPos.X
		sand.Y = newPos.Y
		if sand.Equals(*spawnPoint) {
			break
		}
		if !infiniteFloor && sand.Y > maxY {
			break
		}
	}

	return len(sands)
}

func parseInput(lines []string) (map[u.Vec2]bool, int) {
	maxY := 0
	walls := make(map[u.Vec2]bool)
	for _, line := range lines {
		line := strings.ReplaceAll(line, " ", "")
		pointInputs := strings.Split(line, "->")
		var points []u.Vec2
		for _, pi := range pointInputs {
			parts := strings.Split(pi, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			points = append(points, *u.NewVec2(x, y))
		}

		for i := 0; i < len(points)-1; i++ {
			start := points[i]
			target := points[i+1]
			walls[start] = true
			walls[target] = true

			crawler := start.Copy()
			dir := getCrawlDirection(start, target)

			for !crawler.Equals(target) {
				crawler.AddTo(dir)
				walls[*crawler] = true
			}

			if crawler.Y > maxY {
				maxY = crawler.Y
			}
		}
	}

	return walls, maxY
}

func getCrawlDirection(start u.Vec2, target u.Vec2) u.Vec2 {
	dir := u.NewVec2(0, 0)

	if target.X < start.X {
		dir.X = -1
	} else if target.X > start.X {
		dir.X = 1
	}

	if target.Y < start.Y {
		dir.Y = -1
	} else if target.Y > start.Y {
		dir.Y = 1
	}

	return *dir
}
