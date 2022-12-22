package day18

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day18() {
	lines := utils.GetInput("./day18/input-test.txt")
	fmt.Println("Task1: ", task1(lines))
	fmt.Println("Task2: ", task2(lines))

}

func task1(lines []string) int {
	cubes := parseCubes(lines)

	sum := 0

	for key := range cubes {
	   sum += uncoveredFaces(cubes, key)
    }

	return sum
}

func task2(lines []string) int {
	sum := 0

	return sum
}

type cube struct {
	x int
	y int
	z int
}

func (c* cube) copy() cube {
	return cube{x: c.x, y:c.y, z:c.z}
}

func (c* cube) add(in cube) cube {
	return cube{x:c.x+in.x, y:c.y+in.y, z:c.z+in.z}
}

func parseCubes(lines []string) map[cube]bool {
	out := make(map[cube]bool)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x,_ := strconv.Atoi(parts[0])
		y,_ := strconv.Atoi(parts[1])
		z,_ := strconv.Atoi(parts[2])
		out[cube{x:x,y:y,z:z}] = true
	}
	return out
}

func uncoveredFaces(cubes map[cube]bool, c cube) int {
	count := 0
	dirs := []cube{{x:0, y:0, z:1}, {x:0, y:0, z:-1},{x:1, y:0, z:0},{x:-1, y:0, z:0},{x:0, y:1, z:0},{x:0, y:-1, z:0}}
	for _, dir := range dirs {
		check := c.copy()
		check.add(dir)
		if !cubes[check.add(dir)] {
			count++
		}
	}
	return count
}
