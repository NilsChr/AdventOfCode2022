package day14

import (
	u "advent-of-code-2022/utils"
	"fmt"
)

func render(walls map[u.Vec2]bool, sands []u.Vec2, sand u.Vec2, dimX u.Vec2, dimY u.Vec2) {
	u.ClearConsole()
	
	fmt.Println("DEBUG")
	fmt.Println("Sand", sand)
	fmt.Println("Sands:", len(sands))
	fmt.Println("Walls:", len(walls))
	//fmt.Println(walls)

	for y := 0; y <= dimY.Y; y++ {
		line := fmt.Sprintf("%d ", y) //"1," + string(rune(y)) + " n"
		for x := 450; x <= 550; x++ {
			//fmt.Println(y)
			pos := u.NewVec2(x, y)
			if sand.Equals(*pos) {
				line += "o"
			} else if u.ContainsGeneric(sands, *pos) {
				line += "o"
			} else if walls[*pos] {
				line += "#"
			} else {
				line += "."
			}
			
		}
		fmt.Println(line)
	}
}