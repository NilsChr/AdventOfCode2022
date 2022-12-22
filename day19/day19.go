package day19

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day19() {
	lines := utils.GetInput("./day19/input-test.txt")
	fmt.Println("Task1: ", task1(lines))
}

func task1(lines []string) string {


	f := createFactory(lines[0])
	fmt.Println(f)
	f.toString()

	for i := 0; i < 24; i++ {

		fmt.Printf("== Minute %d == %d ore, %d clay, %d obs, %d geode\n", (i+1), f.ore, f.clay, f.obsidian, f.geode)

		// Build
		buyType := 0
		if f.canBuyGeodeRobot() {
			f.buyGeodeRobot()
			buyType = 4
		} else if f.canBuyObsidianRobot() && i < 18  {
			f.buyObsidianRobot()
			buyType = 3
		} else if f.canBuyClayRobot() && i < 12 {
			f.buyClayRobot()
			buyType = 2
		} else if f.canBuyOreRobot() && i < 6 {
			f.buyOreRobot()
			buyType = 1
		}
		
		
		

		f.yield()
		// Add Robots
		if buyType == 1 {
			fmt.Println("Buy ore")
			f.addOreRobot()
		}
		if buyType == 2 {
			fmt.Println("Buy clay")
			f.addClayRobot()
		}
		if buyType == 3 {
			fmt.Println("Buy obsidian")
			f.addObsidianRobot()
		}
		if buyType == 4 {
			fmt.Println("Buy geode")
			f.addGeodeRobot()
		}

		fmt.Println()
	}

	fmt.Println(f)
	fmt.Println("Ore: ", f.ore)
	fmt.Println("Clay: ", f.clay)
	fmt.Println("Obsidian: ", f.obsidian)
	fmt.Println("Geode: ", f.geode)

	fmt.Println("OreRobots: ", f.oreRobots)
	fmt.Println("ClayRobots: ", f.clayRobots)
	fmt.Println("ObsidianRobots: ", f.obsidianRobots)
	fmt.Println("GeodeRobots: ", f.geodeRobots)

	return ""
}
//                                  6                           12                               18         21                             27        30
//Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.

func createFactory(line string) factory {
	parts := strings.Split(line, " ")
	id0 := strings.Replace(parts[1], ":", "",1)
	id,_ := strconv.Atoi(id0)

	var f factory
	f.id = id
	f.costRobotOre_ore,_  = strconv.Atoi(parts[6])
	f.costRobotClay_ore,_  = strconv.Atoi(parts[12])
	f.costRobotObsidian_ore,_  = strconv.Atoi(parts[18])
	f.costRobotObsidian_clay,_  = strconv.Atoi(parts[21])
	f.costRobotGeode_ore, _ = strconv.Atoi(parts[27])
	f.costRobotGeode_obsidian, _ = strconv.Atoi(parts[30])

	f.oreRobots = 1

	return f
}

type factory struct {
	id int

	ore int
	clay int
	obsidian int
	geode int

	oreRobots int
	clayRobots int
	obsidianRobots int
	geodeRobots int

	costRobotOre_ore int
	costRobotClay_ore int
	costRobotObsidian_ore int
	costRobotObsidian_clay int
	costRobotGeode_ore int
	costRobotGeode_obsidian int
}

func (f* factory) toString() {
	fmt.Println("FACTORY: ", f.id)
	fmt.Printf("Each ore robot costs %d ore.\n", f.costRobotOre_ore)
	fmt.Printf("Each clay robot costs %d ore.\n", f.costRobotClay_ore)
	fmt.Printf("Each obsidian robot costs %d ore and %d clay.\n", f.costRobotObsidian_ore, f.costRobotObsidian_clay)
	fmt.Printf("Each geode robot costs %d ore and %d obsidian.\n", f.costRobotGeode_ore, f.costRobotGeode_obsidian)

}

// yield
func (f* factory) yield() {
	f.ore += f.oreRobots
	f.clay += f.clayRobots
	f.obsidian += f.obsidianRobots
	f.geode += f.geodeRobots

	if f.oreRobots > 0 {
		fmt.Printf("%d ore-collecting robot collects %d ore; you now have %d ore.\n", f.oreRobots, f.oreRobots, f.ore)
	}
	if f.clayRobots > 0 {
		fmt.Printf("%d clay-collecting robot collects %d clay; you now have %d clay.\n", f.clayRobots, f.clayRobots, f.clay)
	}
	if f.obsidianRobots > 0 {
		fmt.Printf("%d obsidian-collecting robot collects %d obsidian; you now have %d obsidian.\n", f.obsidianRobots, f.obsidianRobots, f.obsidian)
	}
	if f.geodeRobots > 0 {
		fmt.Printf("%d geode-cracking robot cracks %d geode; you now have %d geode.\n", f.geodeRobots, f.geodeRobots, f.geode)
	}
}
// can afford OreRobot
func (f* factory) canBuyOreRobot() bool {
	return f.ore >= f.costRobotOre_ore
}
func (f* factory) buyOreRobot() {
	fmt.Printf("Spend %d ore to start building a ore-collecting robot.\n", f.costRobotOre_ore)
	f.ore -= f.costRobotOre_ore
}
func (f* factory) addOreRobot() {
	f.oreRobots++
}

// can afford ClayRobot
func (f* factory) canBuyClayRobot() bool {
	return f.ore >= f.costRobotClay_ore && f.clay + (f.clayRobots*2) <= f.costRobotObsidian_clay
}
func (f* factory) buyClayRobot() {
	fmt.Printf("Spend %d ore to start building a clay-collecting robot.\n", f.costRobotClay_ore)
	f.ore -= f.costRobotClay_ore
}
func (f* factory) addClayRobot() {
	f.clayRobots++
}

// can afford ObsidianRobot
func (f* factory) canBuyObsidianRobot() bool {
	return f.ore >= f.costRobotObsidian_ore && f.clay >= f.costRobotObsidian_clay && f.obsidian + (f.obsidianRobots*2) <= f.costRobotGeode_obsidian
}
func (f* factory) buyObsidianRobot() {
	fmt.Printf("Spend %d ore and %d clay to start building an obsidian-collecting robot.\n", f.costRobotObsidian_ore, f.costRobotObsidian_clay)
	f.ore -= f.costRobotObsidian_ore
	f.clay -= f.costRobotObsidian_clay
}
func (f* factory) addObsidianRobot() {
	f.obsidianRobots++
} 

// can afford GeodeRobot
func (f* factory) canBuyGeodeRobot() bool {
	return f.ore >= f.costRobotGeode_ore && f.obsidian >= f.costRobotGeode_obsidian
}
func (f* factory) buyGeodeRobot() {
	fmt.Printf("Spend %d ore and %d obsidian to start building an geode-cracking robot.\n", f.costRobotGeode_ore, f.costRobotGeode_obsidian)

	f.ore -= f.costRobotGeode_ore
	f.obsidian -= f.costRobotGeode_obsidian
}
func (f* factory) addGeodeRobot() {
	f.geodeRobots++
} 
