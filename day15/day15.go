package day15

import (
	u "advent-of-code-2022/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
)


func Day15() {
	path, t1, t2 := config(false)
	lines := u.GetInput(path)
	fmt.Println("Task1:", task1(lines, t1))
	fmt.Println("Task2:", task2(lines, t2))
}

func config(test bool) (string, int, int) {
	if test {
		return "./day15/input-test.txt", 10, 20
	}
	return "./day15/input.txt", 2000000, 4000000
}	

func task1(lines []string, height int) int {
	var sensors []Sensor

	for _, line := range lines {
		sensors = append(sensors, parseSensors(line))
	}
	minMax := getMinMaxX(sensors)
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
		if found {
			sum++
		}
	}
	return sum
}

func task2(lines []string, limit int) int {
	var sensors []Sensor

	for _, line := range lines {
		sensors = append(sensors, parseSensors(line))
	}
	var goal u.Vec2
	for y := 0; y < limit; y++ {
		for x := 0; x < limit; x++ {
			pos := u.NewVec2(x, y)
			found := false
			for _, s := range sensors {
				dist := manhattenDistance(*pos, s.pos)
				if dist <= s.r {
					found = true
					diffY :=  int(math.Abs(float64(s.pos.Y - pos.Y)))
					x = s.pos.X + s.r-diffY
				}
			}
			if !found {
				goal = *pos
				return goal.X*4000000 + goal.Y
			}
		}
	}
	return goal.X*4000000 + goal.Y
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


