package day15

import (
	"advent-of-code-2022/utils"
	u "advent-of-code-2022/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func Day15() {
	lines := u.GetInput("./day15/input-test.txt")

	// task1 prod: 2000000 test:10
	fmt.Println("Task1:", task1(lines, 10))
	fmt.Println("Task2:", task2(lines, 20))

}

func task1(lines []string, height int) int {
	var sensors []Sensor

	for _, line := range lines {
		sensors = append(sensors, parseSensors(line))
	}
	minMax := getMinMaxX(sensors)
	minMax.X *= 2
	minMax.Y *= 2

	sum := 0
	for x := minMax.X; x <= minMax.Y; x++ {
		pos := u.Vec2{X: x, Y: height}
		found := false
		for _, sensor := range sensors {
			dist := manhattenDistance(sensor.pos, pos)
			if dist <= sensor.r && !pos.Equals(sensor.beacon) {
				found = true
			}
		}
		if found != false {
			sum++
		}
	}
	debug(sensors)
	return sum
}

func task2(lines []string, limit int) int {
	var sensors []Sensor

	for _, line := range lines {
		sensors = append(sensors, parseSensors(line))
	}
	minMax := getMinMaxX(sensors)
	//minMax.X *= 2
	//minMax.Y *= 2

	sum := 0
	for x := 0; x <= limit; x++ {
		for y := 0; y < limit; y++ {
			pos := u.Vec2{X: x, Y: y}
			found := false
			for _, sensor := range sensors {
				dist := manhattenDistance(sensor.pos, pos)
				if dist <= sensor.r && !pos.Equals(sensor.beacon) {
					found = true
				}
			}
			if found != false {
				sum++
				if pos.X > minMax.X && pos.Y < minMax.Y {
					fmt.Println("HEY", pos)

				}
			}
		}

	}
	debug(sensors)
	return sum
}

type Sensor struct {
	pos    u.Vec2
	r      int
	beacon u.Vec2
}

func parseSensors(line string) Sensor {
	var re = regexp.MustCompile(`([-]?\d+)`)
	var sensor Sensor
	matches := re.FindAllString(line, -1)
	sensor.pos.X, _ = strconv.Atoi(matches[0])
	sensor.pos.Y, _ = strconv.Atoi(matches[1])
	sensor.beacon.X, _ = strconv.Atoi(matches[2])
	sensor.beacon.Y, _ = strconv.Atoi(matches[3])
	sensor.r = manhattenDistance(sensor.pos, sensor.beacon)
	return sensor
}

func manhattenDistance(a u.Vec2, b u.Vec2) int {
	return int(math.Abs(float64(b.X-a.X)) + math.Abs(float64(a.Y-b.Y)))
}

func getMinMaxX(sensors []Sensor) u.Vec2 {
	var minMax u.Vec2
	minX := 999999
	maxX := -999999
	for _, sensor := range sensors {
		if sensor.pos.X-sensor.r < minX {
			minX = sensor.pos.X - sensor.r
		}
		if sensor.pos.X+sensor.r > maxX {
			maxX = sensor.pos.X + sensor.r
		}
	}
	minMax.X = minX
	minMax.Y = maxX
	return minMax
}

func debug(sensors []Sensor) {
	utils.ClearConsole()
	for y := -10; y < 30; y++ {
		row := fmt.Sprintf("|%6d| ", y)
		for x := -10; x < 30; x++ {
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
			}
			row += point
		}
		fmt.Println(row)
	}

}

func getSensor(sensors []Sensor, pos u.Vec2) Sensor {
	var s Sensor

	for _, sens := range sensors {
		if sens.pos.Equals(pos) {
			return sens
		}
	}

	return s
}
