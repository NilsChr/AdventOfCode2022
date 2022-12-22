package main

import (
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day10"
	"advent-of-code-2022/day11"
	"advent-of-code-2022/day12"
	"advent-of-code-2022/day13"
	"advent-of-code-2022/day14"
	"advent-of-code-2022/day15"
	"advent-of-code-2022/day16"
	"advent-of-code-2022/day17"
	"advent-of-code-2022/day18"
	"advent-of-code-2022/day19"
	"advent-of-code-2022/day2"
	"advent-of-code-2022/day20"
	"advent-of-code-2022/day21"
	"advent-of-code-2022/day22"
	"advent-of-code-2022/day3"
	"advent-of-code-2022/day4"
	"advent-of-code-2022/day5"
	"advent-of-code-2022/day6"
	"advent-of-code-2022/day7"
	"advent-of-code-2022/day8"
	"advent-of-code-2022/day9"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("no arguments added")
		return
	}
	day := os.Args[1:][0]
	start := time.Now()

	switch day {
	case "1":
		fmt.Println("Running day1")
		day1.Day1()
	case "2":
		fmt.Println("Running day2")
		day2.Day2()
	case "3":
		fmt.Println("Running day3")
		day3.Day3()
	case "4":
		fmt.Println("Running day4")
		day4.Day4()
	case "5":
		fmt.Println("Running day5")
		day5.Day5()
	case "6":
		fmt.Println("Running day6")
		day6.Day6()
	case "7":
		fmt.Println("Running day7")
		day7.Day7()
	case "8":
		fmt.Println("Running day8")
		day8.Day8()
	case "9":
		fmt.Println("Running day8")
		day9.Day9()
	case "10":
		fmt.Println("Running day10")
		day10.Day10()
	case "11":
		fmt.Println("Running day11")
		day11.Day11()
	case "12":
		fmt.Println("Running day12")
		day12.Day12()
	case "13":
		fmt.Println("Running day13")
		day13.Day13_2()
	case "14":
		fmt.Println("Running day14")
		day14.Day14()
	case "15":
		fmt.Println("Running day15")
		day15.Day15()
	case "16":
		fmt.Println("Running day16")
		day16.Day16()
	case "17":
		fmt.Println("Running day17")
		day17.Day17()
	case "18":
		fmt.Println("Running day18")
		day18.Day18()
	case "19":
		fmt.Println("Running day19")
		day19.Day19()
	case "20":
		fmt.Println("Running day20")
		day20.Day20()
	case "21":
		fmt.Println("Running day21")
		day21.Day21()
	case "22":
		fmt.Println("Running day22")
		day22.Day22()
	}

	elapsed := time.Since(start)
	fmt.Printf("Time %f seconds\n", elapsed.Seconds())

}
