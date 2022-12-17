package day16

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day16() {
	lines := utils.GetInput("./day16/input-test.txt")
	fmt.Println("Task1: ", task1(lines))
}

type valve struct {
	name     string
	flowrate int
	children []*valve
	open     bool
}

func (v *valve) getHighestFlowRateChild( visited map[string]bool) *valve {
	var highest valve
	for _, valve := range v.children {
		if valve.flowrate > highest.flowrate && !valve.open {
			highest = *valve
		}
	}
	return &highest

}

func task1(lines []string) int {
	valves := parseValves(lines)
	current := findValveByName(valves, "AA")

	//next := current.getHighestFlowRateChild()
	//fmt.Println(next)
	pressureRelease := 0
	visited := make(map[string]bool)
	visited[current.name] = true
	for i := 0; i < 30; i++ {
		fmt.Printf("min %d: current %v\n", i, current)
		for j := 0; j < len(valves); j++ {
			if valves[j].open {
				pressureRelease += valves[j].flowrate
			}
		}
		if current.flowrate == 0 {
			current = current.getHighestFlowRateChild(visited)
			visited[current.name] = true
			continue
		}
		if !current.open {
			fmt.Println("Opening: ", current)
			current.open = true
			current = current.getHighestFlowRateChild(visited)
			visited[current.name] = true
			continue
		}

	}

	return pressureRelease
}

func parseValves(lines []string) []*valve {
	var valves []*valve
	connections := make(map[string]string)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		name := strings.TrimSpace(parts[1])
		flowrate := parts[4]
		flowrate = strings.Replace(flowrate, "rate=", "", 1)
		flowrate = strings.Replace(flowrate, ";", "", 1)
		flowrateVal, _ := strconv.Atoi(flowrate)
		valveSplit := "valve"
		if strings.Index(line, "valves") > 1 {
			valveSplit = "valves"
		}
		connections[name] = strings.Split(line, valveSplit)[1]
		valves = append(valves, &valve{name, flowrateVal, []*valve{}, false})
	}

	for i, valve := range valves {
		links := connections[valve.name]
		linkNames := strings.Split(links, ",")
		for _, link := range linkNames {
			link := strings.TrimSpace(link)
			child := findValveByName(valves, link)
			valves[i].children = append(valves[i].children, child)
		}
	}

	return valves
}

func findValveByName(valves []*valve, name string) *valve {
	var v valve
	for _, valve := range valves {
		if valve.name == name {
			return valve
		}
	}

	return &v
}
