package day15

import (
	u "advent-of-code-2022/utils"
	"fmt"
	"time"
)

func debug(sensors []Sensor, scanX int, scanY int) {
	time.Sleep(200 * time.Millisecond)
	u.ClearConsole()
	//var goal u.Vec2
	for y := 10; y < 40; y++ {
		row := fmt.Sprintf("|%6d| ", y)
		for x := -10; x < 20; x++ {
			pos := u.NewVec2(x, y)
			point := ""
			for _, s := range sensors {
				if manhattenDistance(*pos, s.pos) <= s.r {
					point = "#"
				}
				if pos.Equals(s.beacon) {
					point = "B"
				}
				if pos.Equals(s.pos) {
					point = "S"
				}
			}
			//sens := getSensor(sensors, *u.NewVec2(x, y))
			if point == "" {
				point = "."
				//fmt.Println(pos)
				//goal = *pos
			}

			if y == scanY && x == scanX {
				point = "*"
			}
			row += point
		}
		fmt.Println(row)
	}
	//fmt.Println("Goal", goal)

}

